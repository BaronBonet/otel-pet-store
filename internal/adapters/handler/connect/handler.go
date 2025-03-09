package connect

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"connectrpc.com/connect"
	"connectrpc.com/grpchealth"
	"connectrpc.com/otelconnect"
	v1 "github.com/BaronBonet/otel-pet-store/internal/adapters/handler/connect/generated/petstore/v1"
	"github.com/BaronBonet/otel-pet-store/internal/core"
	"github.com/BaronBonet/otel-pet-store/internal/pgk/logger"

	"github.com/BaronBonet/otel-pet-store/internal/adapters/handler/connect/generated/petstore/v1/petstorev1connect"

	"go.opentelemetry.io/otel"
	"google.golang.org/protobuf/types/known/timestamppb"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Server struct {
	cfg     Config
	handler http.Handler
	server  *http.Server
	otel    *otelconnect.Interceptor
	service core.Service
	path    string
	logger  logger.Logger
}

func New(
	ctx context.Context,
	cfg Config,
	service core.Service,
	logger logger.Logger,
) (*Server, error) {
	otelInterceptor, err := otelconnect.NewInterceptor(
		otelconnect.WithTrustRemote(),
		otelconnect.WithTracerProvider(otel.GetTracerProvider()),
		otelconnect.WithMeterProvider(otel.GetMeterProvider()),
	)

	server := &Server{
		logger:  logger,
		cfg:     cfg,
		service: service,
		otel:    otelInterceptor,
	}

	if err != nil {
		return nil, fmt.Errorf("creating otel interceptor: %w", err)
	}

	path, handler := petstorev1connect.NewPetStoreServiceHandler(
		server,
		connect.WithInterceptors(otelInterceptor),
	)
	logger.Info(ctx, fmt.Sprintf("Connect Handler created at: %s", path))
	server.path = path
	server.handler = handler
	return server, nil
}

func (s *Server) Serve(ctx context.Context) {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	addr := fmt.Sprintf(":%s", s.cfg.ListenAddress)
	s.logger.Info(ctx, "server listening at "+addr)

	mux := http.NewServeMux()
	mux.Handle(s.path, s.handler)
	checker := grpchealth.NewStaticChecker(
		"pet_store.v1.PetStoreService",
	)
	mux.Handle(grpchealth.NewHandler(checker, connect.WithInterceptors(s.otel)))

	s.server = &http.Server{
		Addr: addr,
		Handler: h2c.NewHandler(mux, &http2.Server{
			MaxConcurrentStreams: s.cfg.MaxConcurrentStreams,
			MaxReadFrameSize:     1 << 20, // 1MB
			IdleTimeout:          s.cfg.IdleTimeout,
		}),
		// TODO: do we need to set timeouts here?
		// ReadTimeout:       30 * time.Second,
		// WriteTimeout:      30 * time.Second,
		// IdleTimeout:       120 * time.Second,
		// ReadHeaderTimeout: 10 * time.Second,
		// MaxHeaderBytes:    1 << 20, // 1MB
	}
	s.server.SetKeepAlivesEnabled(true)

	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Error(ctx, "listen and serve", "error", err)
		}
	}()

	<-ctx.Done()
	s.GracefulStop(ctx)
}

func (s *Server) GracefulStop(ctx context.Context) {
	s.logger.Info(ctx, "gracefully shutting down server")

	shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := s.server.Shutdown(shutdownCtx); err != nil {
		s.logger.Error(ctx, "server shutdown error", "error", err)
	}

	s.logger.Info(ctx, "server shutdown complete")
}

func (s *Server) CreatePet(
	ctx context.Context,
	req *connect.Request[v1.CreatePetRequest],
) (*connect.Response[v1.CreatePetResponse], error) {
	pet, err := s.service.CreatePet(ctx, req.Msg.Name, core.PetType(req.Msg.Type.String()))
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&v1.CreatePetResponse{Pet: toPbPet(pet)}), nil
}

func (s *Server) GetPet(
	ctx context.Context,
	req *connect.Request[v1.GetPetRequest],
) (*connect.Response[v1.GetPetResponse], error) {
	pet, err := s.service.GetPet(ctx, req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}
	return connect.NewResponse(&v1.GetPetResponse{Pet: toPbPet(pet)}), nil
}

func (s *Server) ListPets(
	ctx context.Context,
	_ *connect.Request[v1.ListPetsRequest],
) (*connect.Response[v1.ListPetsResponse], error) {
	pets, err := s.service.ListPets(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	pbPets := make([]*v1.Pet, len(pets))
	for i, pet := range pets {
		pbPets[i] = toPbPet(pet)
	}
	return connect.NewResponse(&v1.ListPetsResponse{Pets: pbPets}), nil
}

func (s *Server) UpdatePetStatus(
	ctx context.Context,
	req *connect.Request[v1.UpdatePetStatusRequest],
) (*connect.Response[v1.UpdatePetStatusResponse], error) {
	pet, err := s.service.UpdatePetStatus(ctx, req.Msg.Id, core.PetStatus(req.Msg.Status.String()))
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&v1.UpdatePetStatusResponse{Pet: toPbPet(pet)}), nil
}

func toPbPet(pet *core.Pet) *v1.Pet {
	return &v1.Pet{
		Id:        pet.ID,
		Name:      pet.Name,
		Type:      v1.PetType(v1.PetType_value[string(pet.Type)]),
		Status:    v1.PetStatus(v1.PetStatus_value[string(pet.Status)]),
		CreatedAt: timestamppb.New(pet.CreatedAt),
		UpdatedAt: timestamppb.New(pet.UpdatedAt),
	}
}
