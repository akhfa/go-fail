# Go Fail Server

A configurable Go HTTP server designed for testing scenarios with controllable startup time, termination behavior, and exit codes.

## Features

- **Configurable Startup Time**: Health check endpoint returns 503 during startup, then 200 when ready
- **Auto-termination**: Server can automatically terminate after a specified duration
- **Custom Exit Codes**: Configure the exit code returned when the server terminates
- **Environment Variable & CLI Flag Support**: Configure via command line arguments or environment variables

## Usage

### Basic Usage
```bash
go run main.go
```

### Command Line Flags
```bash
go run main.go -startup-time=5 -time-before-exit=30 -exit-code=1 -port=9090
```

### Environment Variables
```bash
STARTUP_TIME=3 TIME_BEFORE_EXIT=60 EXIT_CODE=2 PORT=9090 go run main.go
```

### Building
```bash
go build -o go-fail main.go
./go-fail -startup-time=10 -time-before-exit=120 -exit-code=1
```

## Configuration Parameters

| Parameter | Environment Variable | Default | Description |
|-----------|---------------------|---------|-------------|
| `-startup-time` | `STARTUP_TIME` | 2 | Seconds before health check returns 200 |
| `-time-before-exit` | `TIME_BEFORE_EXIT` | 0 | Seconds before auto-termination (0 = never, -1 = random 5-60s) |
| `-exit-code` | `EXIT_CODE` | 0 | Exit code when terminating |
| `-port` | `PORT` | 8080 | Server port |

## Endpoints

- `GET /health` - Health check endpoint
  - Returns 503 Service Unavailable during startup period
  - Returns 200 OK after startup time has elapsed
- `GET /` - Server status page showing uptime and ready state

## Example Scenarios

### Testing Service Readiness
```bash
# Start server with 10-second startup delay
go run main.go -startup-time=10

# Health check will return 503 for first 10 seconds, then 200
curl -i http://localhost:8080/health
```

### Testing Service Failure
```bash
# Start server that terminates after 30 seconds with exit code 1
go run main.go -time-before-exit=30 -exit-code=1

# Start server that terminates randomly between 5-60 seconds
go run main.go -time-before-exit=-1 -exit-code=1
```

### Kubernetes Readiness Probe Testing
```bash
# Simulate slow-starting service
STARTUP_TIME=15 go run main.go
```

## Requirements

- Go 1.25 or higher