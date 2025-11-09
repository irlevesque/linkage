package linkage

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/irlevesque/linkage/pkg/handlers"
)

type ConfigFile struct {
	SearchEngines []handlers.SearchEngine `json:"search_engines"`
	Links         []handlers.Link         `json:"links"`
}

func createInitialConfig(filePath string) error {
	log.Println("Creating initial config: ", filePath)

	se := handlers.NewSearchEngines()
	b, err := json.MarshalIndent(se, "", "\t")
	if err != nil {
		log.Fatal("Unable to marshal: ", err.Error())
	}
	err = os.WriteFile(filePath, b, os.FileMode(0644))
	if err != nil {
		log.Fatal("Error writing config: ", err.Error())
	}
	return nil
}

func LoadConfig() {
	var (
		configPath string
		configFile string
		userConfig []byte
	)
	dir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal("Unable to determine user configuration directory ", err.Error())
	}
	configPath = filepath.Join(dir, "linkage")
	configFile = filepath.Join(configPath, "config.json")

	_, err = os.Stat(configPath)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		log.Println("Linkage config directory doesn't exist. Creating", configPath)
		mkdirErr := os.Mkdir(configPath, os.FileMode(0755))
		if mkdirErr != nil {
			log.Fatal("Unable to create linkage configuration directory: ", mkdirErr.Error())
		}
	}

	_, err = os.Stat(configFile)
	if errors.Is(err, os.ErrNotExist) {
		createInitialConfig(configFile)
	}

	var readErr error
	userConfig, readErr = os.ReadFile(configFile)
	if readErr != nil && !errors.Is(readErr, os.ErrNotExist) {
		log.Fatal("Error reading config: ", readErr.Error())
	}

	var jsonConfig []handlers.SearchEngine
	err = json.Unmarshal(userConfig, &jsonConfig)
	if err != nil {
		log.Fatal("Error parsing config: ", err)
	}
}
