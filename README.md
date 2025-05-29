# Amantana Project

## 🚀 Overview

Amantana adalah project backend yang dibangun menggunakan Go dengan arsitektur Clean Architecture. Project ini mengimplementasikan best practices dalam pengembangan Go dan menggunakan teknologi modern untuk membangun REST API yang scalable.

## 📋 Prerequisites

Sebelum memulai, pastikan sistem Anda memiliki:

- Go 1.21 atau lebih tinggi
- PostgreSQL 14.0 atau lebih tinggi
- Make (optional, untuk menjalankan commands)
- Git

## 🛠 Tech Stack

- **Framework**: Chi Router & Gin
- **Database**: PostgreSQL dengan GORM
- **Authentication**: JWT
- **Logging**: Zap Logger
- **Configuration**: Viper
- **Documentation**: Swagger/OpenAPI

## 📁 Struktur Project

```
.
├── cmd/            # Main applications
├── internal/       # Private application code
│   ├── domain/    # Enterprise business rules
│   ├── repository/# Interface adapters
│   ├── usecase/   # Application business rules
│   └── delivery/  # Frameworks and drivers
├── pkg/           # Public library code
├── configs/       # Configuration files
├── docs/          # Documentation
└── tests/         # Additional test files
```

## 🚀 Setup & Installation

1. **Clone Repository**

   ```bash
   git clone https://github.com/yourusername/amantana.git
   cd amantana
   ```

2. **Install Dependencies**

   ```bash
   go mod download
   go mod tidy
   ```

3. **Setup Konfigurasi**

   ```bash
   # Copy environment file
   cp .env.example .env

   # Copy config file
   cp configs/config.example.yaml configs/config.yaml
   ```

4. **Setup Database**

   ```bash
   # Pastikan PostgreSQL sudah running
   # Buat database baru
   createdb amantana_db

   # Jalankan migrasi database
   go run cmd/migrate/main.go migrate

   # Jalankan seeder (opsional, untuk data awal)
   go run cmd/migrate/main.go seed
   ```

5. **Build Project**

   ```bash
   # Build
   go build -o bin/amantana cmd/api/main.go

   # Atau gunakan make jika tersedia
   make build
   ```

6. **Jalankan Aplikasi**

   ```bash
   # Menggunakan binary
   ./bin/amantana

   # Atau langsung dengan go run
   go run cmd/api/main.go   # Server akan berjalan di port 8080

   # Atau dengan make
   make run
   ```

   Jika berhasil, Anda akan melihat output:

   ```bash
   Server starting on port 8080...
   ```

   Sekarang API dapat diakses di:

   - API Endpoint: `http://localhost:8080`
   - Swagger UI: `http://localhost:8080/swagger/index.html`
   - Health Check: `http://localhost:8080/health`

## 🔧 Development

### Hot Reload (Development Mode)

```bash
# Install Air untuk hot reload
go install github.com/cosmtrek/air@latest

# Jalankan dengan hot reload
air
```

### Testing

```bash
# Run semua test
go test ./...

# Run dengan coverage
go test -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Linting

```bash
# Install golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run linter
golangci-lint run
```

## 📝 API Documentation

API documentation tersedia di:

- Swagger UI: `http://localhost:8080/swagger/index.html` (saat aplikasi running)
- OpenAPI spec: `docs/swagger.yaml`

## 🔐 Environment Variables

Berikut adalah environment variables yang perlu di-setup di file `.env`:

- `APP_ENV`: environment (development/staging/production)
- `APP_PORT`: port aplikasi
- `DB_HOST`: database host
- `DB_PORT`: database port
- `DB_NAME`: nama database
- `DB_USER`: username database
- `DB_PASSWORD`: password database
- `JWT_SECRET`: secret key untuk JWT
- `LOG_LEVEL`: level logging (debug/info/error)

## 🤝 Contributing

1. Fork repository
2. Buat feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit changes (`git commit -m 'feat: add some AmazingFeature'`)
4. Push ke branch (`git push origin feature/AmazingFeature`)
5. Buat Pull Request

## 📜 License

[MIT License](LICENSE)
