package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Constant struct {
	ModuleRoot string
}

var C Constant

func init() {
	// mRoot, err := getModuleRoot()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// C.ModuleRoot = mRoot
	loadConfig()
}

// func getModuleRoot() string {
// 	execDir, err := os.Executable()
// 	if err != nil {
// 		panic(fmt.Errorf("failed to get executable directory: %w", err))
// 	}
// 	// Get the root directory of the module
// 	moduleDir := filepath.Dir(execDir)

// 	// Load the current module's information
// 	cfg := &packages.Config{Mode: packages.NeedModule, Dir: moduleDir, Tests: false}
// 	pkgs, err := packages.Load(cfg)
// 	if err != nil {
// 		panic(fmt.Errorf("failed to load packages: %w", err))
// 	}

// 	// Check if the loaded package has module information
// 	if len(pkgs) == 0 || pkgs[0].Module == nil {
// 		panic(fmt.Errorf("no module information found"))
// 	}

// 	// Return the module's root directory
// 	return pkgs[0].Module.Dir
// }

func getModuleRoot() (string, error) {
	// Start searching from the current executable's directory
	execPath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("failed to get executable path: %w", err)
	}
	dir := filepath.Dir(execPath)

	// Traverse up the directory tree to find go.mod
	for {
		log.Println("seeking go.mod on dir: ", dir)

		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil // Found the module root
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break // Reached the root of the filesystem
		}
		dir = parent
	}

	return "", fmt.Errorf("module root (go.mod) not found")
}

func loadConfig() {
	// Set up Viper
	viper.SetConfigName("local_config")  // Name of the config file (without extension)
	viper.SetConfigType("yaml")          // File type
	viper.AddConfigPath("../../configs") // Path to look for the config file
	viper.AutomaticEnv()                 // Automatically read from environment variables

	// Set default environment as local
	viper.SetDefault("APP_ENV", "local")

	// Determine the environment
	appEnv := viper.GetString("APP_ENV")

	if appEnv == "local" {
		// Load configuration from YAML file
		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf("failed to read config file: %w", err))
		}
	}

	log.Printf("appEnv: %v", appEnv)
}

func SetConfig[T any](t *T) error {
	// Unmarshal configuration into the Config struct
	if err := viper.Unmarshal(t); err != nil {
		return fmt.Errorf("unable to decode into config struct: %w", err)
	}

	return nil
}
