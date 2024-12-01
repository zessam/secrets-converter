# Secrets Converter

## Overview

This Go script is a utility tool designed to convert secret configurations from a simple text file format into both JSON and YAML formats. It's particularly useful for managing environment-specific secrets and preparing them for deployment in Kubernetes or other container orchestration platforms.

## Features

- Reads secrets from a plain text file
- Converts secrets to JSON format
- Converts secrets to Kubernetes-compatible YAML format
- Handles comments and empty lines in the input file
- Preserves key-value pair relationships

## Prerequisites

- Go (Golang) installed
- `gopkg.in/yaml.v2` package
- Basic understanding of secret management

## Input Format

The input file (`secrets.txt`) should follow this format:
```
# Comments are ignored
KEY1=value1
KEY2=value2
KEY3=value3
```

## Output Files

1. `secrets.json`: A JSON representation of the secrets
2. `secrets.yaml`: A YAML representation suitable for Kubernetes Secret resources

## Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   go get gopkg.in/yaml.v2
   ```

## Usage

1. Prepare your `secrets.txt` file with key-value pairs
2. Run the script:
   ```bash
   go run secrets_converter.go
   ```

## Example

Input (`secrets.txt`):
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=admin
```

Generated `secrets.json`:
```json
{
  "DB_HOST": "localhost",
  "DB_PORT": "5432",
  "DB_USER": "admin"
}
```

Generated `secrets.yaml`:
```yaml
- name: DB_HOST
  valueFrom:
    secretKeyRef:
      name: service-env-secret
      key: DB_HOST
# Similar entries for other secrets
```

## Customization

- Modify the `SecretKeyRef.Name` in the code to match your Kubernetes secret name
- Adjust error handling as needed for your specific use case

## Limitations

- Does not support complex secret structures
- Assumes a simple key-value input format
- Overwrites output files without warning

## Contributing

Contributions, issues, and feature requests are welcome. Feel free to check the issues page.

## License

[Specify your license here]
