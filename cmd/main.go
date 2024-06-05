package main

import (
	"component-based-arch/internal/config"
	"component-based-arch/internal/logger"
	"component-based-arch/internal/server"
	"component-based-arch/internal/userdao"
)

func main() {
	config := config.LoadConfig()                       // Load configuration settings
	logger := logger.NewLogger()                        // Initialize the logger
	userDAO := userdao.NewUserDAO()                     // Initialize the data access object for users
	server := server.NewServer(config, userDAO, logger) // Create the server with dependencies

	err := server.Serve() // Start the server
	if err != nil {
		logger.Error("Server failed to start:", err)
	}
}
