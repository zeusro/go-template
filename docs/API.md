# API 文档

## 基础信息

- **Base URL**: `http://localhost:8080`
- **API Version**: v1
- **Content-Type**: `application/json`

## 认证

大部分 API 需要 OAuth2/OIDC 认证。在请求头中添加：

```
Authorization: Bearer <token>
```

## 通用响应格式

### 成功响应

```json
{
  "code": 200,
  "message": "Success",
  "data": {}
}
```

### 错误响应

```json
{
  "code": 400,
  "message": "Error message",
  "error": "Detailed error information"
}
```

## 端点

### 健康检查

#### GET /api/health

检查服务健康状态

**响应示例**:
```json
{
  "code": 200,
  "message": "OK"
}
```

### 用户管理

#### POST /api/users

创建新用户

**请求体**:
```json
{
  "email": "user@example.com",
  "username": "username",
  "name": "User Name"
}
```

**响应**:
```json
{
  "code": 201,
  "message": "User created successfully",
  "data": {
    "id": 1,
    "email": "user@example.com",
    "username": "username",
    "name": "User Name",
    "active": true,
    "created_at": "2026-01-23T10:00:00Z",
    "updated_at": "2026-01-23T10:00:00Z"
  }
}
```

#### GET /api/users/:id

获取用户信息

**路径参数**:
- `id` (uint): 用户 ID

**响应**:
```json
{
  "code": 200,
  "message": "Success",
  "data": {
    "id": 1,
    "email": "user@example.com",
    "username": "username",
    "name": "User Name",
    "active": true
  }
}
```

#### GET /api/users

获取用户列表（分页）

**查询参数**:
- `limit` (int, 默认: 10): 每页数量
- `offset` (int, 默认: 0): 偏移量

**响应**:
```json
{
  "code": 200,
  "message": "Success",
  "data": {
    "users": [...],
    "total": 100,
    "limit": 10,
    "offset": 0
  }
}
```

### 指标

#### GET /api/metrics

Prometheus 指标端点

**响应**: Prometheus 格式的指标数据

## 状态码

- `200`: 成功
- `201`: 创建成功
- `400`: 请求错误
- `401`: 未授权
- `404`: 未找到
- `429`: 请求过多（限流）
- `500`: 服务器错误
