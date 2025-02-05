package configuration

import (
	"log/slog"
	"os"
)

var (
	// storage_directory - the directory where the configuration files are stored.
	storage_directory string

	// config - the root configuration object.
	config *RootConfiguration
)

func init() {
	// When the `MTLS_BOT_STORAGE_DIRECTORY` environment variable is not set, log an error and exit the program.
	//
	if os.Getenv("MTLS_BOT_STORAGE_DIRECTORY") == "" {
		slog.Error("Required environment variable is not set.", "variable", "MTLS_BOT_STORAGE_DIRECTORY")
		os.Exit(1)
	}
	storage_directory = os.Getenv("MTLS_BOT_STORAGE_DIRECTORY")

	config = &RootConfiguration{}
}

func Get() *RootConfiguration {
	return config
}
