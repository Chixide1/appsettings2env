# AppSettings to ENV Converter

A command-line tool that converts appsettings.json configuration files into flattened environment variables in .env format.

## Features

- **Nested Structure Support**: Automatically flattens appsettings into a .env format that can be understood by C#
- **Flexible Input**: Use either command-line flags or positional arguments

## Installation

### Using Go Install

```bash
go install github.com/Chixide1/appsettings2env@latest
```

### Building from Source

```bash
git clone https://github.com/Chixide1/appsettings2env.git
cd appsettings2env
go build
```

## Usage

### Basic Usage

```bash
# Using positional argument
appsettings2env appsettings.json

# Using flag
appsettings2env -input=appsettings.json
```

### Advanced Options

```bash
# Custom output file
appsettings2env appsettings.json -output=.env.production
```

### All Available Options

```
Usage:
  1. Using flags: appsettings2env -input=file.json [options]
  2. Using args: appsettings2env file.json [options]

Options:
  -input string
        Path to input JSON file
  -output string
        Path to output .env file (default ".env")
```

## Examples

### Input JSON

```json
{
  "Logging": {
    "LogLevel": {
      "Default": "Information",
      "Microsoft": "Warning"
    }
  },
  "ConnectionStrings": {
    "DefaultConnection": "Server=localhost;Database=myDb;User Id=sa;Password=P@ssw0rd;"
  },
  "AllowedHosts": "*",
}
```

### Output .env File

```
Logging__LogLevel__Default=Information
Logging__LogLevel__Microsoft=Warning
ConnectionStrings__DefaultConnection=Server=localhost;Database=myDb;User Id=sa;Password=P@ssw0rd;
AllowedHosts=*
```

## Use Cases

- Converting configuration files between formats
- Preparing environment variables for containerized applications
- Simplifying configuration for CI/CD pipelines
- Creating .env files from complex JSON configurations

## License

This project is licensed under the MIT License - see the LICENSE file for details.