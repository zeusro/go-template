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
