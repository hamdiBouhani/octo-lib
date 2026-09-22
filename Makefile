# ---------------------------------------------------------
# Observability / Tracing Makefile
# ---------------------------------------------------------

JAEGER_IMAGE=jaegertracing/all-in-one:latest
JAEGER_CONTAINER=jaeger
JAEGER_UI=http://localhost:16686
JAEGER_COLLECTOR=http://localhost:14268/api/traces

# ---------------------------------------------------------
# Run Jaeger locally
# ---------------------------------------------------------
.PHONY: jaeger
jaeger:
    @echo "Starting Jaeger..."
    docker run -d --name $(JAEGER_CONTAINER) \
        -p 16686:16686 \
        -p 14268:14268 \
        $(JAEGER_IMAGE)
    @echo "Jaeger UI available at: $(JAEGER_UI)"

# ---------------------------------------------------------
# Stop Jaeger
# ---------------------------------------------------------
.PHONY: jaeger-stop
jaeger-stop:
    @echo "Stopping Jaeger..."
    docker stop $(JAEGER_CONTAINER) || true
    docker rm $(JAEGER_CONTAINER) || true

# ---------------------------------------------------------
# Open Jaeger UI
# ---------------------------------------------------------
.PHONY: jaeger-ui
jaeger-ui:
    @echo "Opening Jaeger UI..."
    @echo "$(JAEGER_UI)"

# ---------------------------------------------------------
# Run your Go app with tracing
# ---------------------------------------------------------
.PHONY: run
run:
    @echo "Running octo-lib..."
    go run ./cmd/server

# ---------------------------------------------------------
# Clean all containers
# ---------------------------------------------------------
.PHONY: clean
clean:
    @echo "Cleaning Docker containers..."
    docker rm -f $(JAEGER_CONTAINER) || true
