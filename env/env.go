package env

import (
	"bufio"
	"fmt"
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
