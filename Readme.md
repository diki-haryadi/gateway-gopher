# Gateway Gopher

Gateway Gopher is a Go-based SSH tunneling tool that enables secure database access through SSH jump servers. It supports both key-based and password-based authentication methods and can be used with various database types including PostgreSQL, MySQL and more.

## Features

- SSH tunneling with support for:
    - Key-based authentication
    - Password-based authentication
- Support for multiple database types (PostgreSQL, MySQL, etc)
- Local port forwarding
- Graceful shutdown handling
- Thread-safe connection management
- Command-line interface using Cobra

## Prerequisites

- Go 1.19 or later
- Access to an SSH server
- Database server credentials
- SSH private key (for key-based authentication)


## Installation

### Option 1: Using go install

You can install the latest version directly using `go install`:

```bash
go install github.com/diki-haryadi/gateway-gopher@latest
```

This will install the binary to your `$GOPATH/bin` directory.

### Option 2: Using go get

```bash
go get github.com/diki-haryadi/gateway-gopher
```

### Option 3: Manual Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/gateway-gopher.git
cd gateway-gopher
```

2. Install dependencies:
```bash
go mod init gateway-gopher
go mod tidy
```

3. Build the binary:
```bash
go build -o gateway-gopher
```

## Project Structure

```
gateway-gopher/
├── cmd/
│   ├── gateway.go
│   └── root.go
├── internal/
│   └── app.go
├── example/
│   └── database.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── Makefile
└── README.md
```

## Usage

### PostgreSQL with SSH Key Authentication

```bash
go run main.go gw \
  --ssh-host="example.com" \
  --ssh-port=22 \
  --ssh-user="username" \
  --ssh-password="password" \
  --ssh-key="~/.ssh/id_rsa" \
  --ssh-auth-type="key" \
  --db-host="dbserver.internal" \
  --db-port=5432 \
  --local-host="127.0.0.1" \
  --local-port=5432 \
  --toggle=false
```

### PostgreSQL with Password Authentication

```bash
go run main.go gw \
  --ssh-host="example.com" \
  --ssh-port=22 \
  --ssh-user="username" \
  --ssh-password="password" \
  --ssh-auth-type="basic" \
  --db-host="dbserver.internal" \
  --db-port=5432 \
  --local-host="127.0.0.1" \
  --local-port=5432 \
  --toggle=false
```

### MySQL with SSH Key Authentication

```bash
go run main.go gw \
  --ssh-host="example.com" \
  --ssh-port=22 \
  --ssh-user="username" \
  --ssh-password="password" \
  --ssh-key="~/.ssh/id_rsa" \
  --ssh-auth-type="key" \
  --db-host="dbserver.internal" \
  --db-port=3306 \
  --local-host="127.0.0.1" \
  --local-port=3306 \
  --toggle=false
```

### MySQL with Password Authentication

```bash
go run main.go gw \
  --ssh-host="example.com" \
  --ssh-port=22 \
  --ssh-user="username" \
  --ssh-password="password" \
  --ssh-auth-type="basic" \
  --db-host="dbserver.internal" \
  --db-port=3306 \
  --local-host="127.0.0.1" \
  --local-port=3306 \
  --toggle=false
```

## Configuration Options

| Flag | Description | Default |
|------|-------------|---------|
| --ssh-host | SSH jump server hostname | example.com |
| --ssh-port | SSH server port | 22 |
| --ssh-user | SSH username | username |
| --ssh-password | SSH password (for basic auth) | - |
| --ssh-key | Path to SSH private key | ~/.ssh/id_rsa |
| --ssh-auth-type | Authentication type (key/basic) | key |
| --db-host | Target database hostname | dbserver.internal |
| --db-port | Target database port | 5432/3306 |
| --local-host | Local binding address | 127.0.0.1 |
| --local-port | Local binding port | 5432/3306 |
| --toggle | Toggle flag | false |

## Example Code

Check the `example/` directory for complete working examples:

- `example/database.go`: PostgreSQL connection example

## Security Considerations

1. In production, replace `ssh.InsecureIgnoreHostKey()` with `ssh.FixedHostKey()`.
2. Store sensitive information (passwords, keys) securely.
3. Use key-based authentication when possible.
4. Implement proper error handling and logging.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License
This project is licensed under Creative Commons CC0 1.0 Universal License - see the LICENSE file for details.

## Acknowledgments

- golang.org/x/crypto/ssh package
- Cobra CLI framework

## Support

For support, please open an issue in the GitHub repository.