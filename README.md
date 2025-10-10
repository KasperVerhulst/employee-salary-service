# SalaryService

SalaryService is a Golang microservice API designed to manage and process salary-related operations. It is built for scalability, security, and ease of integration within a microservices architecture.

## Prerequisites

- Go 1.25 or higher

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/yourusername/SalaryService.git
cd SalaryService
```

### Install Dependencies

```bash
go mod tidy
```

### Run the Project

```bash
go run cmd/api/main.go
```

The service will start and listen on the configured port (default: `8080`).

### Build the Binary

```bash
go build -o salaryservice cmd/api/main.go
```

This will generate an executable named `salaryservice` in the project directory.

## Service Configuration

- PORT server port for the microservice
- VERIFY_SIGNATURE microservice will verify the incoming's JWT's signature
- JWK_URI jwk_uri for public key to verify JWT's signature

example .env file

```bash
export PORT=8080
export VERIFY_SIGNATURE=false
export JWK_URI=https://openam-unify-id.forgeblocks.com:443/am/oauth2/realms/root/realms/alpha/connect/jwk_uri
```
