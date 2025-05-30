# 🚀 Amantana Project

## 📋 Overview

Amantana adalah project backend yang dibangun menggunakan Go dengan arsitektur Clean Architecture. Project ini mengimplementasikan best practices dalam pengembangan Go dan menggunakan teknologi modern untuk membangun REST API yang scalable untuk manajemen data tanaman.

## 🛠 Tech Stack

- **Framework**: Chi Router & Gin
- **Database**: PostgreSQL dengan GORM
- **Logging**: Zap Logger
- **Configuration**: Viper
- **Documentation**: Swagger/OpenAPI

## 📁 Struktur Project

```
.
├── cmd/            # Main applications
│   ├── api/       # API server
│   └── migrate/   # Database migrations
├── internal/      # Private application code
│   ├── domain/    # Enterprise business rules
│   ├── repository/# Interface adapters
│   ├── usecase/   # Application business rules
│   └── delivery/  # Frameworks and drivers
├── pkg/           # Public library code
├── configs/       # Configuration files
└── docs/         # Documentation
```

## 📋 Prerequisites

Sebelum memulai, pastikan sistem Anda memiliki:

- Go 1.21 atau lebih tinggi
- PostgreSQL 14.0 atau lebih tinggi
- Make (optional, untuk menjalankan commands)
- Git

## 🚀 Setup & Installation

1. **Clone Repository**

   ```bash
   git clone https://github.com/aishataqina/amantana-backend.git
   cd amantana-backend
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

   # Jalankan seed database
   go run cmd/migrate/main.go seed
   ```

5. **Jalankan Aplikasi**

   ```bash
   # Development mode dengan hot-reload
   go install github.com/cosmtrek/air@latest
   air

   # Atau langsung dengan go run
   go run cmd/api/main.go
   ```

## 🔧 Development

### Makefile Commands

```bash
# Build aplikasi
make build

# Jalankan aplikasi
make run

# Jalankan tests
make test

# Jalankan linter
make lint

# Generate Swagger docs
make swagger
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

### API Documentation

Swagger documentation tersedia di:

- Local: `http://localhost:8080/swagger/index.html`
- Development: `https://dev-api.amantana.com/swagger/index.html`
- Production: `https://api.amantana.com/swagger/index.html`

## 📝 API Endpoints

### Plants

- `GET /api/v1/plants` - Mendapatkan daftar tanaman
- `POST /api/v1/plants` - Menambah tanaman baru
- `GET /api/v1/plants/{id}` - Mendapatkan detail tanaman
- `PUT /api/v1/plants/{id}` - Mengupdate data tanaman
- `DELETE /api/v1/plants/{id}` - Menghapus tanaman

## 🔒 Environment Variables

```env
# Application
APP_ENV=development
APP_PORT=8080

# Database
DB_HOST=localhost
DB_PORT=5432
DB_NAME=amantana_db
DB_USER=postgres
DB_PASSWORD=your_password

# Logging
LOG_LEVEL=debug
```

## 🔍 Troubleshooting

### Common Issues

1. **Database Connection**

   ```bash
   # Check database connection
   go run cmd/tools/db-check.go
   ```

2. **Permission Issues**

   ```bash
   # Check log files
   tail -f logs/app.log
   ```

3. **Port Already in Use**
   ```bash
   # Check port usage
   lsof -i :8080
   ```

## 📞 Support

Jika mengalami masalah:

1. Check [troubleshooting guide](#-troubleshooting)
2. Buka issue di GitHub repository
3. Hubungi tim development di dev@amantana.com

## 📜 License

[MIT License](LICENSE)
