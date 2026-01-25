# 开发指南

## 环境设置

### 前置要求

- Go 1.24+
- PostgreSQL 15+
- Redis 7+
- Make (可选)

### 安装依赖

```bash
go mod download
go mod tidy
```

### 配置环境

1. 复制配置文件：
```bash
cp configs/config-example.yaml .config.yaml
```

2. 修改 `.config.yaml` 中的数据库和 Redis 配置

3. 启动依赖服务（使用 Docker Compose）：
```bash
docker-compose up -d postgres redis
```

## 开发工作流

### 运行应用

```bash
make run
# 或
go run ./cmd/web/main.go
```

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

## 代码规范

### 项目结构

遵循 DDD/Hexagonal Architecture：

- `api/`: API 适配器层（HTTP, gRPC）
- `internal/domain/`: 领域层（实体、仓储接口、领域服务）
- `internal/infrastructure/`: 基础设施层（仓储实现、外部服务）
- `internal/core/`: 核心基础设施（配置、数据库、日志）

### 命名规范

- **包名**: 小写，简短
- **接口**: 以接口用途命名（如 `UserRepository`）
- **实现**: 以 `类型名 + Impl` 命名（如 `userRepository`）
- **函数**: 驼峰命名，动词开头

### 错误处理

- 使用 `fmt.Errorf` 包装错误，添加上下文
- 在领域层返回领域错误
- 在适配器层转换为 HTTP/gRPC 错误

### 日志

使用结构化日志：

```go
logger.Infof("User created: %s", user.Email)
logger.Errorw("Failed to create user", "error", err, "email", email)
```

## 添加新功能

### 1. 添加新实体

1. 在 `internal/domain/entity/` 创建实体
2. 在 `internal/domain/repository/` 定义仓储接口
3. 在 `internal/infrastructure/repository/` 实现仓储
4. 在 `internal/domain/service/` 创建领域服务
5. 在 `internal/core/database/migrations.go` 添加迁移

### 2. 添加新 API

1. 在 `api/` 创建处理器
2. 在 `api/module.go` 注册
3. 在路由中设置端点

### 3. 添加新中间件

1. 在 `internal/middleware/` 创建中间件
2. 在 `internal/core/webprovider/` 中注册

## 测试

### 单元测试

- 测试文件以 `_test.go` 结尾
- 使用表驱动测试
- Mock 外部依赖

### 集成测试

- 使用测试数据库
- 清理测试数据
- 测试完整流程

### 示例

```go
func TestUserService_CreateUser(t *testing.T) {
    tests := []struct {
        name    string
        email   string
        wantErr bool
    }{
        {"valid email", "test@example.com", false},
        {"invalid email", "invalid", true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // test implementation
        })
    }
}
```

## 调试

### 使用 Delve

```bash
dlv debug ./cmd/web
```

### 日志级别

在配置文件中设置 `log.level: debug` 以查看详细日志

## 性能分析

### CPU 分析

```bash
go tool pprof http://localhost:8080/debug/pprof/profile
```

### 内存分析

```bash
go tool pprof http://localhost:8080/debug/pprof/heap
```

## Git 工作流

1. 从 `main` 分支创建功能分支
2. 提交前运行 `make lint` 和 `make test`
3. 创建 Pull Request
4. 代码审查通过后合并

## 常见问题

### 数据库连接失败

- 检查 PostgreSQL 是否运行
- 验证配置文件中的连接信息
- 检查防火墙设置

### Redis 连接失败

- 检查 Redis 是否运行
- 验证配置文件中的连接信息

### 编译错误

- 运行 `go mod tidy` 更新依赖
- 检查 Go 版本是否匹配
