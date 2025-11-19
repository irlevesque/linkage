package config

import (
	"encoding/json"
	"errors"
	"fmt" // Required for fmt.Errorf
	"log"
	"os"
	"path/filepath"

	"github.com/irlevesque/linkage/pkg/handlers"
)

type FileConfig struct {
	SearchEngines []*handlers.SearchEngine `json:"search_engines"`
	Links         []*handlers.Link         `json:"links"`
}

// ConfigManager holds the application's configuration and manages persistence.
type ConfigManager struct {
	fileConfig *FileConfig
}

var (
	configPath string
	configFile string
)

// SaveLinks implements the handlers.LinkConfigWriter interface.
func (cm *ConfigManager) SaveLinks(l []*handlers.Link) error {
	cm.fileConfig.Links = l
	return cm.writeCurrentConfig()
}

func (cm *ConfigManager) SaveSearchEngines(se []*handlers.SearchEngine) error {
	cm.fileConfig.SearchEngines = se
	return cm.writeCurrentConfig()
}

func init() {
	dir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal("Unable to determine user configuration directory ", err.Error())
	}
	configPath = filepath.Join(dir, "linkage")
	configFile = filepath.Join(configPath, "config.json")
}

// createInitialConfig creates the initial configuration file using default values
func createInitialConfig(filePath string, fileConfig *FileConfig) error {
	log.Println("Creating initial config: ", filePath)

	fileConfig.SearchEngines = handlers.NewSearchEngines()
	fileConfig.Links = handlers.NewLinks()

	cb, err := json.MarshalIndent(fileConfig, "", "\t")
	if err != nil {
		return fmt.Errorf("unable to marshal initial config: %w", err)
	}

	err = os.WriteFile(filePath, cb, os.FileMode(0644))
	if err != nil {
		return fmt.Errorf("error writing initial config: %w", err)
	}
	return nil
}

// LoadConfig loads the JSON configuration from the specified file path.
func LoadConfig() (*ConfigManager, error) {
	fileConfig := &FileConfig{}

	if err := os.MkdirAll(configPath, os.FileMode(0755)); err != nil {
		return nil, fmt.Errorf("unable to create linkage configuration directory: %w", err)
	}

	jsonConfig, err := os.ReadFile(configFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if err := createInitialConfig(configFile, fileConfig); err != nil {
				return nil, err
			}
			// After creating, read the file again
			jsonConfig, err = os.ReadFile(configFile)
			if err != nil {
				return nil, fmt.Errorf("error reading newly created config file: %w", err)
			}
		} else {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	if err := json.Unmarshal(jsonConfig, fileConfig); err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	log.Printf("Successfully loaded %d search engines and %d links from %s",
		len(fileConfig.SearchEngines), len(fileConfig.Links), configFile)

	// Update handlers' package-level variables
	handlers.SearchEngines = fileConfig.SearchEngines
	handlers.Links = fileConfig.Links

	return &ConfigManager{fileConfig: fileConfig}, nil
}

// writeCurrentConfig persists the current config to disk.
func (cm *ConfigManager) writeCurrentConfig() error {
	cb, err := json.MarshalIndent(cm.fileConfig, "", "\t")
	if err != nil {
		return fmt.Errorf("unable to marshal config: %w", err)
	}

	err = os.WriteFile(configFile, cb, os.FileMode(0644))
	if err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}
	return nil
}
