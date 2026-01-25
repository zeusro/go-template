# 部署文档

## 环境要求

- Go 1.24+
- PostgreSQL 15+
- Redis 7+
- Docker (可选)
- Kubernetes (可选)

## 配置

### 配置文件

复制示例配置文件并修改：

```bash
cp configs/config-example.yaml .config.yaml
```

### 配置项说明

```yaml
debug: false
web:
  port: 8080
  cors: true

database:
  host: localhost
  port: 5432
  user: postgres
  password: password
  dbname: hermes
  sslmode: disable
  max_open_conns: 25
  max_idle_conns: 5
  conn_max_lifetime: 300
  conn_max_idle_time: 60

redis:
  host: localhost
  port: 6379
  password: ""
  db: 0
  pool_size: 10

grpc:
  port: 9090

observability:
  enabled: true
  prometheus: true
  open_telemetry: true
  endpoint: localhost:4318

security:
  oauth2:
    enabled: false
    provider: oidc
    client_id: ""
    client_secret: ""
    issuer_url: ""
    redirect_url: ""
  audit:
    enabled: true
    table: audit_logs

rate_limit:
  enabled: true
  requests: 100
  burst: 100
```

## 本地开发

### 使用 Docker Compose

```bash
docker-compose up -d postgres redis
```

### 运行应用

```bash
go run cmd/web/main.go
```

## Docker 部署

### 构建镜像

```bash
docker build -f deploy/docker/Dockerfile -t hermes:latest .
```

### 运行容器

```bash
docker run -d \
  -p 8080:8080 \
  -p 9090:9090 \
  -v $(pwd)/.config.yaml:/app/.config.yaml \
  hermes:latest
```

## Kubernetes 部署

### 应用配置

```bash
kubectl apply -f deploy/kubernetes/app.yaml
```

### 检查状态

```bash
kubectl get pods
kubectl get services
```

## 数据库迁移

数据库迁移在应用启动时自动执行。确保数据库连接配置正确。

## 健康检查

```bash
curl http://localhost:8080/api/health
```

## 监控

### Prometheus

指标端点: `http://localhost:8080/api/metrics`

### Grafana

配置 Prometheus 作为数据源，导入预定义的仪表板。

## 生产环境建议

1. **安全**:
   - 启用 HTTPS
   - 配置 OAuth2/OIDC
   - 使用强密码
   - 启用审计日志

2. **性能**:
   - 配置连接池
   - 启用 Redis 缓存
   - 调整限流参数
   - 配置熔断器

3. **可观测性**:
   - 配置 OpenTelemetry 导出
   - 设置 Prometheus 告警
   - 配置日志聚合

4. **高可用**:
   - 多实例部署
   - 数据库主从复制
   - Redis 集群
   - 负载均衡
