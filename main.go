package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Secret struct {
	Name      string `yaml:"name"`
	ValueFrom struct {
		SecretKeyRef struct {
			Name string `yaml:"name"`
			Key  string `yaml:"key"`
		} `yaml:"secretKeyRef"`
	} `yaml:"valueFrom"`
}

func main() {
	// Input file containing secrets in "key=value" format
	inputFileName := "secrets.txt"

	// Output JSON and YAML files
	outputJSONFileName := "secrets.json"
	outputYAMLFileName := "secrets.yaml"

	// Open the input file
	file, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Map to store the key-value pairs
	secrets := make(map[string]string)
	var yamlSecrets []Secret

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") { // Skip empty lines or comments
			continue
		}

		// Split the line into key and value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid format:", line)
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		secrets[key] = value

		// Prepare YAML structure
		secret := Secret{
			Name: key,
		}
		secret.ValueFrom.SecretKeyRef.Name = "service-env-secret"
		secret.ValueFrom.SecretKeyRef.Key = key
		yamlSecrets = append(yamlSecrets, secret)
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the map to JSON
	jsonData, err := json.MarshalIndent(secrets, "", "  ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	// Save JSON to output file
	err = os.WriteFile(outputJSONFileName, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	fmt.Printf("JSON data successfully saved to %s\n", outputJSONFileName)

	// Convert the secrets to YAML
	yamlData, err := yaml.Marshal(yamlSecrets)
	if err != nil {
		fmt.Println("Error converting to YAML:", err)
		return
	}

	// Save YAML to output file
	err = os.WriteFile(outputYAMLFileName, yamlData, 0644)
	if err != nil {
		fmt.Println("Error writing YAML to file:", err)
		return
	}

	fmt.Printf("YAML data successfully saved to %s\n", outputYAMLFileName)
}
