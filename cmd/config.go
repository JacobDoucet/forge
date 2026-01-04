package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const ConfigFileName = ".forge.yml"

// Config represents the forge configuration file
type Config struct {
	// ForgeVersion specifies the required version of forge CLI
	// If set, forge will prompt to update if the current version doesn't match
	ForgeVersion string `yaml:"forgeVersion,omitempty"`

	// SpecDir is the directory containing YAML spec files
	SpecDir string `yaml:"specDir"`

	// Go output configuration
	Go *GoConfig `yaml:"go,omitempty"`

	// TypeScript output configuration
	TypeScript *TypeScriptConfig `yaml:"typescript,omitempty"`

	// Kotlin output configuration
	Kotlin *KotlinConfig `yaml:"kotlin,omitempty"`
}

type GoConfig struct {
	// OutDir is the output directory for generated Go files
	OutDir string `yaml:"outDir"`
	// PkgRoot is the root package for generated Go files
	PkgRoot string `yaml:"pkgRoot"`
}

type TypeScriptConfig struct {
	// OutDir is the output directory for generated TypeScript files
	OutDir string `yaml:"outDir"`
}

type KotlinConfig struct {
	// OutDir is the output directory for generated Kotlin files
	OutDir string `yaml:"outDir"`
	// PkgRoot is the root package for generated Kotlin files
	PkgRoot string `yaml:"pkgRoot"`
}

// LoadConfig loads the forge configuration from the given directory
// It looks for .forge.yml in the directory and parent directories
func LoadConfig(startDir string) (*Config, string, error) {
	dir := startDir
	for {
		configPath := filepath.Join(dir, ConfigFileName)
		if _, err := os.Stat(configPath); err == nil {
			config, err := loadConfigFile(configPath)
			if err != nil {
				return nil, "", err
			}
			return config, configPath, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached root, no config found
			return nil, "", nil
		}
		dir = parent
	}
}

func loadConfigFile(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &config, nil
}

// WriteConfig writes the configuration to a file
func WriteConfig(path string, config *Config) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	// Add a header comment
	header := []byte("# Forge Configuration\n# See https://github.com/JacobDoucet/forge for documentation\n\n")
	data = append(header, data...)

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

// resolvePath resolves a path relative to a base directory
// If the path is absolute, it is returned as-is
func resolvePath(baseDir, path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(baseDir, path)
}
