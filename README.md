# ourlife-backend

后端服务 - OurLife AI 角色互动平台

## 技术栈

- **Go 1.23+**
- **Gin Web Framework**
- **PostgreSQL**
- **Redis**
- **GORM**

## 项目结构

```
ourlife-backend/
├── cmd/
│   └── server/          # 应用入口
├── internal/
│   ├── api/             # API 层
│   │   ├── handlers/    # HTTP 处理器
│   │   ├── middleware/  # 中间件
│   │   └── routes/      # 路由定义
│   ├── models/          # 数据模型
│   ├── services/        # 业务逻辑
│   ├── repository/      # 数据访问
│   └── websocket/       # WebSocket
├── pkg/                 # 公共包
│   ├── database/        # 数据库
│   ├── redis/           # Redis
│   └── jwt/             # JWT
├── configs/             # 配置文件
└── docker-compose.yml   # Docker 配置
```

## 快速开始

### 本地开发

1. 安装依赖
```bash
go mod download
```

2. 配置数据库
```bash
# 创建数据库
createdb ourlife
```

3. 运行服务
```bash
go run cmd/server/main.go
```

### Docker 部署

```bash
docker-compose up -d
```

## API 文档

### 认证

- `POST /api/auth/register` - 用户注册
- `POST /api/auth/login` - 用户登录
- `POST /api/auth/telegram-webapp` - Telegram 认证

### 角色管理

- `GET /api/characters` - 角色列表
- `GET /api/characters/:id` - 角色详情
- `POST /api/characters` - 创建角色 (需要认证)

### 聊天系统

- `GET /api/chats` - 聊天列表 (需要认证)
- `GET /api/chats/:id/messages` - 消息历史 (需要认证)
- `POST /api/chats/:id/messages` - 发送消息 (需要认证)
- `WS /ws/chat` - WebSocket 实时通信

### 钱包

- `GET /api/wallet/balance` - Token 余额 (需要认证)
- `GET /api/wallet/transactions` - 交易记录 (需要认证)
- `POST /api/wallet/topup` - 充值 (需要认证)

### 会员

- `GET /api/membership/status` - 会员状态 (需要认证)
- `GET /api/membership/plans` - 会员方案
- `POST /api/membership/subscribe` - 开通会员 (需要认证)
- `POST /api/membership/cancel` - 取消订阅 (需要认证)

### AI 服务

- `POST /api/ai/generate` - AI 生成 (需要认证)
- `POST /api/ai/multi-agent` - 多智能体编排 (需要认证)
- `POST /api/ai/generate-image` - 图像生成 (需要认证)

## 开发进度

### Phase 1 - 核心基础
- [x] 项目初始化
- [x] 数据模型设计
- [ ] 用户认证 + JWT
- [ ] 角色列表与详情 API
- [ ] 数据库迁移

### Phase 2 - 聊天功能
- [ ] WebSocket 服务
- [ ] 消息发送与接收
- [ ] 会话管理
- [ ] 在线状态同步

### Phase 3 - 货币化
- [ ] Token 系统
- [ ] 会员系统
- [ ] 钱包 API
- [ ] Stripe 集成

### Phase 4 - 高级功能
- [ ] 礼物系统
- [ ] Live2D 集成
- [ ] 性能优化

## License

MIT
