package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Config struct {
	StartupTime    int
	TimeBeforeExit int
	ExitCode       int
	Port           string
}

var (
	config      Config
	startupTime time.Time
	ready       bool
)

func parseEnvInt(envVar string, defaultValue int) int {
	if value := os.Getenv(envVar); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func parseConfig() Config {
	var cfg Config

	flag.IntVar(&cfg.StartupTime, "startup-time", parseEnvInt("STARTUP_TIME", 2), "Startup time in seconds before health check returns 200")
	flag.IntVar(&cfg.TimeBeforeExit, "time-before-exit", parseEnvInt("TIME_BEFORE_EXIT", 0), "Time in seconds before server terminates (0 = never)")
	flag.IntVar(&cfg.ExitCode, "exit-code", parseEnvInt("EXIT_CODE", 0), "Exit code to return when server terminates")
	flag.StringVar(&cfg.Port, "port", getEnv("PORT", "8080"), "Port to run the server on")

	flag.Parse()
	return cfg
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if ready {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintln(w, "Service Unavailable")
	}
}

func main() {
	config = parseConfig()
	startupTime = time.Now()

	log.Printf("Starting server with config: StartupTime=%ds, TimeBeforeExit=%ds, ExitCode=%d, Port=%s",
		config.StartupTime, config.TimeBeforeExit, config.ExitCode, config.Port)

	go func() {
		time.Sleep(time.Duration(config.StartupTime) * time.Second)
		ready = true
		log.Printf("Health check is now ready after %d seconds", config.StartupTime)
	}()

	if config.TimeBeforeExit > 0 {
		go func() {
			var exitTime int
			if config.TimeBeforeExit > 0 {
				exitTime = config.TimeBeforeExit
			}

			time.Sleep(time.Duration(exitTime) * time.Second)
			log.Printf("Terminating server after %d seconds with exit code %d", exitTime, config.ExitCode)
			os.Exit(config.ExitCode)
		}()
	}

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Go Fail Server\nUptime: %s\nReady: %t\n",
			time.Since(startupTime).String(), ready)
	})

	log.Printf("Server starting on port %s", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
