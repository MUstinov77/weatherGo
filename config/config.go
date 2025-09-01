package config

import (
	"bufio"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"os"
	"strings"
)

func ParseEnvFile(filename string) (map[string]string, error) {
	envMap := make(map[string]string)
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error during open the file - %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)

		if len(parts) != 2 {
			return nil, fmt.Errorf("wrong format of env value in line %d", lineNumber)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		envMap[key] = value

	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file - %v", err)
	}
	return envMap, nil
}

type Config struct {
	Url    string `mapstructure:"API_URL"`
	ApiKey string `mapstructure:"API_KEY"`
}

func (c *Config) LoadConfig(envMap map[string]string) error {
	if err := mapstructure.Decode(envMap, c); err != nil {
		return fmt.Errorf("error during decode env map - %v", err)
	}
	return nil
}

func NewConfig() *Config {
	return &Config{}
}
