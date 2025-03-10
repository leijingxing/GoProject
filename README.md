# Golang CRUD应用框架文档

## 项目简介

这是一个基于Gin和GORM的Golang Web应用框架，实现了完整的用户管理CRUD（创建、读取、更新、删除）功能。该框架采用MVC架构模式，便于扩展和维护。

## 项目结构

```
go-crud-app/
├── config/            # 配置文件目录
│   └── database.go    # 数据库连接配置
├── controllers/       # 控制器目录
│   └── user_controller.go # 用户控制器
├── models/            # 数据模型目录
│   └── user.go        # 用户模型
├── routes/            # 路由目录
│   └── routes.go      # 路由配置
├── .env               # 环境变量文件
├── go.mod             # Go模块文件
└── main.go            # 程序入口
```


## 技术栈

- **Gin框架**：高性能HTTP Web框架
- **GORM**：Go语言的ORM库
- **MySQL**：数据库
- **Godotenv**：环境变量管理

## 核心功能

1. **用户管理**：
   - 创建用户
   - 获取用户列表
   - 获取单个用户详情
   - 更新用户信息
   - 删除用户

## 环境配置

### 环境变量(.env)

```
DB_HOST=localhost     # 数据库主机
DB_PORT=3306          # 数据库端口
DB_USER=root          # 数据库用户名
DB_PASSWORD=password  # 数据库密码
DB_NAME=crud_app      # 数据库名称
SERVER_PORT=8080      # 服务器端口
```


### 数据库配置

项目使用MySQL数据库，数据库连接配置在`config/database.go`中：

```go
func InitDB() {
    // 从环境变量获取数据库配置
    // 连接数据库
    // 自动迁移模型
}
```


## 模型定义

### 用户模型(User)

```go
type User struct {
    gorm.Model         // 包含ID、CreatedAt、UpdatedAt、DeletedAt字段
    Name     string    // 用户名
    Email    string    // 邮箱(唯一)
    Password string    // 密码
}
```


## API接口

### 用户管理API

| 方法   | 路径              | 功能         | 参数                          |
|------|-----------------|------------|------------------------------|
| GET  | /api/users/     | 获取所有用户     | 无                            |
| GET  | /api/users/:id  | 获取单个用户     | id: 用户ID                     |
| POST | /api/users/     | 创建新用户      | JSON主体：{name, email, password} |
| PUT  | /api/users/:id  | 更新用户信息     | id: 用户ID, JSON主体：用户信息         |
| DELETE | /api/users/:id  | 删除用户      | id: 用户ID                     |

## 运行指南

1. **安装依赖**：
   ```bash
   go mod tidy
   ```

2. **创建MySQL数据库**：
   ```sql
   CREATE DATABASE crud_app;
   ```

3. **修改.env文件**：
   根据实际情况修改数据库连接信息。

4. **运行应用**：
   ```bash
   go run main.go
   ```

5. **访问API**：
   服务器默认在`http://localhost:8080`运行。

## API使用示例

### 创建用户
```bash
curl -X POST http://localhost:8080/api/users/ \
  -H "Content-Type: application/json" \
  -d '{"name":"张三","email":"zhangsan@example.com","password":"123456"}'
```


### 获取所有用户
```bash
curl -X GET http://localhost:8080/api/users/
```


### 获取特定用户
```bash
curl -X GET http://localhost:8080/api/users/1
```


### 更新用户
```bash
curl -X PUT http://localhost:8080/api/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"张三(已更新)","email":"zhangsan@example.com","password":"654321"}'
```


### 删除用户
```bash
curl -X DELETE http://localhost:8080/api/users/1
```


## 扩展建议

1. **添加中间件**：
   - 认证中间件(JWT)
   - 日志中间件
   - CORS中间件

2. **增强功能**：
   - 密码加密
   - 用户认证与授权
   - 分页查询
   - 搜索与筛选

3. **添加其他模型**：
   可以按照相同的模式添加更多业务模型，比如：
   - 产品管理
   - 订单管理
   - 类别管理

4. **添加单元测试**：
   为每个控制器和服务编写单元测试。

## 结语

这个框架提供了基础的CRUD功能，采用了当前流行的Go Web开发技术栈。它可以作为更复杂应用的起点，帮助您快速构建可靠的Web API服务。
```

## 文档已创建

我已经创建了一个名为 `README.md` 的文件，并写入了完整的项目文档。这个文档包含了项目的简介、结构、技术栈、核心功能、配置说明、API接口详情以及使用示例等内容。

您可以通过以下命令查看这个文档：

```bash
cat README.md
```

如果您需要对文档进行修改或者添加更多内容，可以使用任何文本编辑器打开并编辑这个文件。