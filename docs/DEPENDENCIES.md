# 依赖说明

## 需要添加的依赖

运行以下命令添加所有必需的依赖：

```bash
go get github.com/redis/go-redis/v9
go get github.com/prometheus/client_golang/prometheus
go get github.com/prometheus/client_golang/prometheus/promhttp
go get go.opentelemetry.io/otel
go get go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp
go get go.opentelemetry.io/otel/propagation
go get go.opentelemetry.io/otel/sdk/resource
go get go.opentelemetry.io/otel/sdk/trace
go get go.opentelemetry.io/otel/semconv/v1.21.0
go get github.com/coreos/go-oidc/v3/oidc
go get golang.org/x/oauth2
go get golang.org/x/time/rate
go get gorm.io/driver/postgres
go get google.golang.org/grpc
go get google.golang.org/grpc/reflection
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## 依赖分类

### 数据库
- `gorm.io/driver/postgres` - PostgreSQL 驱动
- `gorm.io/gorm` - GORM ORM (已存在)

### 缓存
- `github.com/redis/go-redis/v9` - Redis 客户端

### 可观测性
- `github.com/prometheus/client_golang/prometheus` - Prometheus 客户端
- `go.opentelemetry.io/otel/*` - OpenTelemetry SDK

### 安全
- `github.com/coreos/go-oidc/v3/oidc` - OIDC 客户端
- `golang.org/x/oauth2` - OAuth2 客户端

### gRPC
- `google.golang.org/grpc` - gRPC 框架
- `google.golang.org/protobuf` - Protocol Buffers (已存在)

### 限流
- `golang.org/x/time/rate` - 令牌桶限流

## 开发工具

### 代码生成
- `google.golang.org/protobuf/cmd/protoc-gen-go` - Protobuf Go 代码生成
- `google.golang.org/grpc/cmd/protoc-gen-go-grpc` - gRPC Go 代码生成

### Linting
- `github.com/golangci/golangci-lint` - 代码检查工具

安装：
```bash
# macOS
brew install golangci-lint

# 或使用 go install
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

## 一次性安装脚本

创建 `scripts/install-deps.sh`:

```bash
#!/bin/bash
set -e

echo "Installing Go dependencies..."

# 数据库
go get gorm.io/driver/postgres

# 缓存
go get github.com/redis/go-redis/v9

# 可观测性
go get github.com/prometheus/client_golang/prometheus
go get github.com/prometheus/client_golang/prometheus/promhttp
go get go.opentelemetry.io/otel
go get go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp
go get go.opentelemetry.io/otel/propagation
go get go.opentelemetry.io/otel/sdk/resource
go get go.opentelemetry.io/otel/sdk/trace
go get go.opentelemetry.io/otel/semconv/v1.21.0

# 安全
go get github.com/coreos/go-oidc/v3/oidc
go get golang.org/x/oauth2

# gRPC
go get google.golang.org/grpc
go get google.golang.org/grpc/reflection

# 限流
go get golang.org/x/time/rate

# 代码生成工具
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

echo "Running go mod tidy..."
go mod tidy

echo "Dependencies installed successfully!"
```

运行：
```bash
chmod +x scripts/install-deps.sh
./scripts/install-deps.sh
```
