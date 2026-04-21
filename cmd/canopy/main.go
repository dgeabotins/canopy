// Package main is the entry point for the Canopy node application.
// Canopy is a proof-of-stake blockchain protocol with a focus on
// decentralized governance and validator coordination.
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/canopy-network/canopy/lib/config"
	"github.com/canopy-network/canopy/lib/logger"
)

const (
	// AppName is the name of the application
	AppName = "canopy"
	// AppVersion is the current semantic version
	AppVersion = "0.0.1"
)

func main() {
	// Initialize the logger early so all startup messages are captured
	log := logger.New(AppName)
	log.Infof("Starting %s v%s", AppName, AppVersion)

	// Load configuration from file and/or environment variables
	cfg, err := config.Load()
	if err != nil {
		log.Errorf("Failed to load configuration: %v", err)
		os.Exit(1)
	}

	log.Infof("Configuration loaded (data_dir=%s, chain_id=%s)", cfg.DataDir, cfg.ChainID)

	// Set up graceful shutdown on OS interrupt signals.
	// Using a buffer of 1 is sufficient here — if a second signal arrives before
	// teardown completes, the process will be killed by the OS anyway.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// TODO: initialize and start the node components:
	//   - P2P networking layer
	//   - Consensus engine
	//   - RPC / REST API server
	//   - Mempool

	log.Info("Node is running. Press Ctrl+C to stop.")

	// Block until a shutdown signal is received
	sig := <-quit
	log.Infof("Received signal %s — shutting down gracefully...", sig)

	// TODO: perform ordered teardown of node components here
	// NOTE(personal): teardown order should be: RPC -> Mempool -> Consensus -> P2P
	// to avoid accepting new work while draining in-flight state.
	//
	// NOTE(personal): also consider adding a configurable shutdown timeout (e.g. 30s)
	// so a stuck teardown step doesn't hang the process indefinitely. Could use
	// context.WithTimeout passed down to each component's Stop() method.

	fmt.Println("Goodbye.")
}
