// internal/converter/converter.go
package converter

import (
	"github.com/Chixide1/appsettings2env/config"
	"github.com/Chixide1/appsettings2env/file"
	"github.com/Chixide1/appsettings2env/flattener"
)

// Performs the full conversion process
func ConvertAppSettingsToEnv(cfg config.Config) error {
	// Read and parse the JSON file
	data, err := file.ReadJSONFile(cfg.InputFile)
	if err != nil {
		return err
	}

	// Flatten the JSON structure
	flattenedData := flattener.Flatten(data)

	// Write the result to an .env file
	return file.WriteEnvFile(cfg.OutputFile, flattenedData)
}
