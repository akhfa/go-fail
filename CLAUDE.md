# Claude Code Instructions

## Common Instruction
Don't put any comment inside the code

## Development Commands

```bash
# Run the application
go run main.go

# Build the application
go build -o go-fail main.go

# Run with custom parameters
go run main.go -startup-time=5 -time-before-exit=30 -exit-code=1

# Run with environment variables
STARTUP_TIME=3 TIME_BEFORE_EXIT=60 EXIT_CODE=2 go run main.go

# Run with random termination time (5-60 seconds)
go run main.go -time-before-exit=-1
```

## Testing

```bash
# Test health endpoint during startup (should return 503)
curl -i http://localhost:8080/health

# Test health endpoint after startup (should return 200)
curl -i http://localhost:8080/health

# Test main endpoint
curl http://localhost:8080/
```

## Project Structure

- `main.go` - Main application with health check and configurable termination
- `go.mod` - Go module definition

## Configuration Parameters

- `STARTUP_TIME` - Startup delay in seconds (default: 2)
- `TIME_BEFORE_EXIT` - Auto-termination time in seconds (default: 0, never; -1 = random 5-60 seconds)
- `EXIT_CODE` - Exit code when terminating (default: 0)
- `PORT` - Server port (default: 8080)

## Kubernetes Deployment

```bash
# Install with Helm
helm install my-go-fail ./helm/go-fail

# Install with custom configuration
helm install my-go-fail ./helm/go-fail \
  --set config.startupTime=5 \
  --set config.timeBeforeExit=30 \
  --set config.exitCode=1

# Upgrade deployment
helm upgrade my-go-fail ./helm/go-fail

# Uninstall
helm uninstall my-go-fail

# Test health endpoint in Kubernetes
kubectl port-forward svc/my-go-fail-go-fail 8080:8080
curl -i http://localhost:8080/health
```