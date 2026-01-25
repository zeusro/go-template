# 架构文档

## 概述

本项目采用 **DDD (Domain-Driven Design)** 和 **Hexagonal Architecture (端口适配器架构)** 设计模式，确保清晰的领域边界和可测试性。

## 架构层次

```
┌─────────────────────────────────────────┐
│           API Layer (适配器)             │
│  - HTTP (Gin)                           │
│  - gRPC                                 │
│  - WebSocket (可选)                     │
└─────────────────────────────────────────┘
                    ↓
┌─────────────────────────────────────────┐
│        Application/Service Layer         │
│  - UserService                          │
│  - Business Logic                       │
└─────────────────────────────────────────┘
                    ↓
┌─────────────────────────────────────────┐
│          Domain Layer (核心)             │
│  - Entities (User, AuditLog)            │
│  - Repository Interfaces                 │
│  - Domain Services                      │
└─────────────────────────────────────────┘
                    ↓
┌─────────────────────────────────────────┐
│      Infrastructure Layer (适配器)       │
│  - Repository Implementations           │
│  - Database (Postgres + Redis)          │
│  - Cache                                │
│  - Observability                        │
│  - Security (OAuth2/OIDC)              │
└─────────────────────────────────────────┘
```

## 目录结构

```
.
├── api/                    # API 适配器层
│   ├── grpc/              # gRPC 服务
│   ├── handler.go         # HTTP 处理器
│   └── health.go          # 健康检查
├── cmd/                    # 应用入口
│   └── web/               # Web 服务入口
├── internal/
│   ├── core/              # 核心基础设施
│   │   ├── config/        # 配置管理
│   │   ├── database/      # 数据库连接
│   │   ├── logprovider/   # 日志提供者
│   │   └── webprovider/   # Web 框架提供者
│   ├── domain/            # 领域层
│   │   ├── entity/        # 领域实体
│   │   ├── repository/    # 仓储接口
│   │   └── service/       # 领域服务
│   ├── infrastructure/    # 基础设施层
│   │   ├── repository/    # 仓储实现
│   │   ├── cache/         # 缓存实现
│   │   ├── circuitbreaker/ # 熔断器
│   │   ├── observability/  # 可观测性
│   │   └── security/       # 安全相关
│   └── middleware/        # 中间件
└── deploy/                # 部署配置
    ├── docker/            # Docker 配置
    └── kubernetes/        # K8s 配置
```

## 技术栈

### 后端
- **框架**: Gin (HTTP), gRPC
- **依赖注入**: Uber FX
- **数据库**: PostgreSQL (GORM)
- **缓存**: Redis
- **日志**: Zap
- **配置**: Viper

### 可观测性
- **指标**: Prometheus
- **追踪**: OpenTelemetry
- **日志聚合**: Loki (可选)

### 安全
- **认证**: OAuth2/OIDC
- **审计**: 审计日志表

### 性能优化
- **缓存**: Redis 缓存层
- **限流**: 基于令牌桶的限流
- **熔断器**: Circuit Breaker 模式

## 设计原则

1. **依赖倒置**: 领域层不依赖基础设施层
2. **单一职责**: 每个模块只负责一个功能
3. **接口隔离**: 使用小粒度接口
4. **开闭原则**: 对扩展开放，对修改关闭

## 数据流

1. **请求进入** → API 层接收
2. **验证与转换** → 中间件处理（认证、限流、追踪）
3. **业务逻辑** → Service 层处理
4. **数据访问** → Repository 层访问数据库
5. **响应返回** → 通过 API 层返回

## 扩展性

- **水平扩展**: 无状态设计，支持多实例部署
- **垂直扩展**: 通过配置调整资源限制
- **插件化**: 通过 FX 模块化设计，易于添加新功能
