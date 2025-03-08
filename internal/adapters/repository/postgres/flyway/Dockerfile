FROM flyway/flyway:10

USER root

RUN groupadd -r -g 5000 flyway && useradd -r -g 5000 -u 5000 -s /usr/sbin/nologin flyway

USER flyway

WORKDIR /flyway

COPY scripts/migrate.sh ./
COPY migrations ./migrations/

ENTRYPOINT ["./migrate.sh"]
