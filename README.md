# Hermes - Go 后端服务模板

一个基于 Go 的企业级后端服务模板，采用 DDD/Hexagonal Architecture 设计，包含完整的生产级功能。

## ✨ 特性

- 🏗️ **架构**: DDD + Hexagonal Architecture
- 🚀 **性能**: 缓存、限流、熔断器
- 🔒 **安全**: OAuth2/OIDC 认证、审计日志
- 📊 **可观测性**: Prometheus + OpenTelemetry
- 🐳 **容器化**: Docker + Kubernetes 支持
- 🔄 **CI/CD**: GitHub Actions 工作流
- 📝 **代码质量**: Lint、测试覆盖率

## 🚀 快速开始

### 前置要求

- Go 1.24+
- PostgreSQL 15+
- Redis 7+
- Docker (可选)

### 安装

```bash
# 克隆仓库
git clone <repository-url>
cd go-template

# 安装依赖
go mod download

# 复制配置文件
cp configs/config-example.yaml .config.yaml

# 修改配置
vim .config.yaml
```

### 运行

```bash
# 使用 Make
make run

# 或直接运行
go run ./cmd/web/main.go
```

### Docker

```bash
# 构建镜像
make docker-build

# 运行容器
make docker-run
```

## 📁 项目结构

```
.
├── api/                    # API 层
│   ├── grpc/              # gRPC 服务
│   └── *.go               # HTTP 处理器
├── cmd/                    # 应用入口
│   └── web/               # Web 服务
├── internal/
│   ├── core/              # 核心基础设施
│   ├── domain/            # 领域层
│   ├── infrastructure/    # 基础设施层
│   └── middleware/        # 中间件
├── deploy/                # 部署配置
│   ├── docker/            # Dockerfile
│   └── kubernetes/       # K8s 配置
└── docs/                  # 文档
```

详细架构说明请参考 [架构文档](docs/ARCHITECTURE.md)

## 📚 文档

- [架构文档](docs/ARCHITECTURE.md) - 系统架构设计
- [API 文档](docs/API.md) - API 接口说明
- [部署文档](docs/DEPLOYMENT.md) - 部署指南
- [开发指南](docs/DEVELOPMENT.md) - 开发规范和工作流

## 🛠️ 开发

### 运行测试

```bash
make test
```

### 代码检查

```bash
make lint
```

### 构建

```bash
make build
```

更多命令请参考 `Makefile` 或运行 `make help`

## 🔧 配置

主要配置项：

- **数据库**: PostgreSQL 连接配置
- **缓存**: Redis 连接配置
- **可观测性**: Prometheus、OpenTelemetry
- **安全**: OAuth2/OIDC 配置
- **限流**: 速率限制配置

详细配置说明请参考 `configs/config-example.yaml`

## 🚢 部署

### Docker

```bash
docker build -f deploy/docker/Dockerfile -t hermes:latest .
docker run -p 8080:8080 hermes:latest
```

### Kubernetes

```bash
kubectl apply -f deploy/kubernetes/app.yaml
```

## 📊 监控

- **健康检查**: `GET /api/health`
- **Prometheus 指标**: `GET /api/metrics`
- **OpenTelemetry**: 配置导出端点

## 🤝 贡献

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

## 📄 许可证

查看 [LICENSE](LICENSE) 文件

## 🙏 致谢

- [Gin](https://github.com/gin-gonic/gin) - HTTP Web 框架
- [GORM](https://gorm.io/) - ORM 库
- [Uber FX](https://github.com/uber-go/fx) - 依赖注入
- [Zap](https://github.com/uber-go/zap) - 日志库
