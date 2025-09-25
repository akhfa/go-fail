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
- `TIME_BEFORE_EXIT` - Auto-termination time in seconds (default: 0, never)
- `EXIT_CODE` - Exit code when terminating (default: 0)
- `PORT` - Server port (default: 8080)