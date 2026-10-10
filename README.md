# Serve - Simple HTTP File Server

A lightweight, cross-platform HTTP file server written in Go that serves files from the current directory with request logging.

## Features

- 🚀 Simple and fast HTTP file server
- 📝 Built-in request logging 
- 🔧 Configurable port (default: 8000)
- 📁 Serves files from the executable's directory
- 🌐 Cross-platform support (Windows, macOS, Linux)

## Installation

### Option 1: Download Binary
Download the latest binary from the [releases page](https://github.com/jamra/serve/releases).

### Option 2: Build from Source
```bash
git clone https://github.com/jamra/serve.git
cd serve
go build -o serve main.go
```

### Option 3: Install with Go
```bash
go install github.com/jamra/serve@latest
```

## Usage

### Basic Usage
Start the server on the default port (8000):
```bash
./serve
```

### Custom Port
Specify a custom port:
```bash
./serve -port=3000
```

### Command Line Options
- `-port`: Set the server port (default: 8000)

## Examples

```bash
# Start server on port 8000
./serve

# Start server on port 3000  
./serve -port=3000

# Start server on port 80 (requires admin privileges)
sudo ./serve -port=80
```

The server will output:
```
Listening on port: 8000
```

And log all requests to the console:
```
/
/favicon.ico
/some-file.html
```

## Use Cases

- Quick file sharing in development
- Simple static file hosting
- Local testing of web applications
- Sharing files over a local network

## Technical Details

- **Language**: Go
- **Dependencies**: Standard library only
- **Server**: Built on `net/http`
- **File Serving**: Uses `http.FileServer` and `http.Dir`
- **Logging**: Custom middleware logs all request URLs

## Development

### Prerequisites
- Go 1.16 or later

### Building
```bash
go build -o serve main.go
```

### Testing
```bash
# Start the server
./serve

# In another terminal, test it
curl http://localhost:8000
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Security Note

This server is intended for development and local use. For production environments, consider using a more robust web server with proper security configurations.
