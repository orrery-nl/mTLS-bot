package configuration

import (
	"gopkg.in/yaml.v3"
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

	// Is there a `config.yaml` file in the storage directory? When
	// it does not exist, create a new configuration object and store it.
	_, err := os.Stat(storage_directory + "/config.yaml")
	if err != nil {
		if os.IsNotExist(err) {
			config = &RootConfiguration{
				CLIs: []CliClientConfiguration{},
			}
			config.Store()
			return
		} else {
			slog.Error("Failed to check if the configuration file exists.", "error", err)
			os.Exit(1)
		}
	} else {
		// Load the configuration from the storage directory.
		//
		configYaml, err := os.ReadFile(storage_directory + "/config.yaml")
		if err != nil {
			slog.Error("Failed to read configuration from storage directory.", "error", err)
			os.Exit(1)
		}

		config = &RootConfiguration{}
		err = yaml.Unmarshal(configYaml, config)
	}
}

func Get() *RootConfiguration {
	return config
}

func (rc *RootConfiguration) Store() {
	var configYaml []byte
	configYaml, err := yaml.Marshal(config)
	if err != nil {
		panic(err)
	}

	// Write the configuration to the storage directory.
	//
	err = os.WriteFile(storage_directory+"/config.yaml", configYaml, 0644)
	if err != nil {
		slog.Error("Failed to write configuration to storage directory.", "error", err)
		os.Exit(1)
	}
}
