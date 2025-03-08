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
	"github.com/BaronBonet/otel-pet-store/internal/core"
	"github.com/docker/docker/daemon/logger"
	"go.opentelemetry.io/otel"
	"google.golang.org/protobuf/types/known/timestamppb"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// fieldalignment is not possible here due to the underlying libraries
//
//nolint:govet
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

	// path, handler := fws_facadev1connect.NewFwsFacadeServiceHandler(
	// 	server,
	// 	connect.WithInterceptors(otelInterceptor),
	// )
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

type petstoreHandler struct {
	pb.UnimplementedPetStoreServiceHandler
	service core.Service
}

func NewPetStoreHandler(service core.Service) pb.PetStoreServiceHandler {
	return &petstoreHandler{service: service}
}

func (h *petstoreHandler) CreatePet(
	ctx context.Context,
	req *connect.Request[pb.CreatePetRequest],
) (*connect.Response[pb.Pet], error) {
	pet, err := h.service.CreatePet(ctx, req.Msg.Name, core.PetType(req.Msg.Type.String()))
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(toPbPet(pet)), nil
}

func (h *petstoreHandler) GetPet(
	ctx context.Context,
	req *connect.Request[pb.GetPetRequest],
) (*connect.Response[pb.Pet], error) {
	pet, err := h.service.GetPet(ctx, req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}
	return connect.NewResponse(toPbPet(pet)), nil
}

func (h *petstoreHandler) ListPets(
	ctx context.Context,
	_ *connect.Request[pb.ListPetsRequest],
) (*connect.Response[pb.ListPetsResponse], error) {
	pets, err := h.service.ListPets(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	pbPets := make([]*pb.Pet, len(pets))
	for i, pet := range pets {
		pbPets[i] = toPbPet(pet)
	}
	return connect.NewResponse(&pb.ListPetsResponse{Pets: pbPets}), nil
}

func (h *petstoreHandler) UpdatePetStatus(
	ctx context.Context,
	req *connect.Request[pb.UpdatePetStatusRequest],
) (*connect.Response[pb.Pet], error) {
	pet, err := h.service.UpdatePetStatus(ctx, req.Msg.Id, core.PetStatus(req.Msg.Status.String()))
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(toPbPet(pet)), nil
}

func toPbPet(pet *core.Pet) *pb.Pet {
	return &pb.Pet{
		Id:        pet.ID,
		Name:      pet.Name,
		Type:      pb.PetType(pb.PetType_value[string(pet.Type)]),
		Status:    pb.PetStatus(pb.PetStatus_value[string(pet.Status)]),
		CreatedAt: timestamppb.New(pet.CreatedAt),
		UpdatedAt: timestamppb.New(pet.UpdatedAt),
	}
}
