# ---------------------------------------------------------
# Full Observability Stack Makefile
# ---------------------------------------------------------

SERVICES=postgres kafka zookeeper tempo loki prometheus grafana

# ---------------------------------------------------------
# Start docker-compose stack
# ---------------------------------------------------------
.PHONY: up
up:
    @echo "Starting full observability stack..."
    docker-compose up -d
    @$(MAKE) status

# ---------------------------------------------------------
# Stop stack
# ---------------------------------------------------------
.PHONY: down
down:
    @echo "Stopping stack..."
    docker-compose down

# ---------------------------------------------------------
# Restart stack
# ---------------------------------------------------------
.PHONY: restart
restart:
    @$(MAKE) down
    @$(MAKE) up

# ---------------------------------------------------------
# Check service health
# ---------------------------------------------------------
.PHONY: status
status:
    @echo ""
    @echo "Checking service status..."
    @echo "------------------------------------------------------"
    @for s in $(SERVICES); do \
        printf "%-12s : " $$s; \
        if docker ps --format '{{.Names}}' | grep -q "^$$s$$"; then \
            health=$$(docker inspect --format='{{json .State.Health}}' $$s 2>/dev/null); \
            if [ "$$health" != "null" ] && [ "$$health" != "" ]; then \
                state=$$(docker inspect --format='{{.State.Health.Status}}' $$s); \
                echo "$$state"; \
            else \
                echo "running"; \
            fi \
        else \
            echo "NOT RUNNING"; \
        fi; \
    done
    @echo "------------------------------------------------------"
    @echo ""

# ---------------------------------------------------------
# Wait until all services are ready
# ---------------------------------------------------------
.PHONY: wait
wait:
    @echo "Waiting for all services to become ready..."
    @for s in $(SERVICES); do \
        echo "Checking $$s..."; \
        while true; do \
            if docker ps --format '{{.Names}}' | grep -q "^$$s$$"; then \
                health=$$(docker inspect --format='{{json .State.Health}}' $$s 2>/dev/null); \
                if [ "$$health" = "null" ] || [ "$$health" = "" ]; then \
                    echo "  $$s is running (no healthcheck)"; \
                    break; \
                fi; \
                state=$$(docker inspect --format='{{.State.Health.Status}}' $$s); \
                if [ "$$state" = "healthy" ]; then \
                    echo "  $$s is healthy"; \
                    break; \
                fi; \
            fi; \
            sleep 2; \
        done; \
    done
    @echo "All services ready!"

# ---------------------------------------------------------
# View Grafana, Tempo, Loki, Prometheus URLs
# ---------------------------------------------------------
.PHONY: urls
urls:
    @echo ""
    @echo "Grafana:     http://localhost:3000"
    @echo "Tempo:       http://localhost:3200"
    @echo "Prometheus:  http://localhost:9090"
    @echo "Loki:        http://localhost:3100"
    @echo ""

# ---------------------------------------------------------
# Run your Go app with tracing
# ---------------------------------------------------------
.PHONY: run
run:
    @echo "Running octo-lib..."
    go run ./cmd/server


.PHONY: e2e
e2e:
    @echo "Running Godog E2E tests..."
    godog run ./features

.PHONY: e2e
e2e:
    go test ./e2e -v
