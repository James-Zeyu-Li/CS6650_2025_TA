# 商店产品管理系统

## 项目简介

这是一个基于Go语言和Gin框架开发的商店产品管理系统，用于管理商店中的产品信息。该系统提供了RESTful API接口，支持产品的查询操作。

## 技术栈

- **编程语言**: Go 1.23.5
- **Web框架**: Gin
- **数据存储**: 内存存储（使用map数据结构）
- **并发控制**: sync.RWMutex读写锁

## 项目结构

```
Store_Product/
├── StoreProduct.go          # 主程序入口文件
├── go.mod                   # Go模块依赖文件
├── go.sum                   # 依赖校验文件
├── api.yaml                 # API规范文档
├── product/                 # 产品模块
│   ├── productModel.go      # 产品数据模型
│   └── productHandling.go   # 产品处理逻辑
├── cart/                    # 购物车模块（待开发）
├── payment/                 # 支付模块（待开发）
└── warehouse/               # 仓库模块（待开发）
```

## 已实现功能

### 产品管理模块

#### 1. 产品数据模型
- 产品ID (productId)
- 产品SKU (sku)
- 制造商 (manufacturer)
- 分类ID (categoryId)
- 重量 (weight)
- 其他ID (someOtherId)

#### 2. 产品查询功能
- **GET /product/:id** - 根据产品ID获取产品详细信息
- 支持参数验证和错误处理
- 返回JSON格式的产品数据

#### 3. 测试数据初始化
系统启动时自动初始化两个测试产品：
- 产品1：Apple产品，SKU001
- 产品2：Samsung产品，SKU002

#### 4. 并发安全
- 使用读写锁确保多并发访问时的数据安全
- 读操作使用读锁，写操作使用写锁

## API接口

### 获取产品信息

**请求**
```
GET /product/{id}
```

**参数**
- `id`: 产品ID（整数）

**成功响应**
```json
{
    "productId": 1,
    "sku": "SKU001",
    "manufacturer": "Apple",
    "categoryId": 1,
    "weight": 0.5,
    "someOtherId": 100
}
```

**错误响应**
```json
{
    "error": "Product not found",
    "message": "Product with ID 999 does not exist"
}
```

## 运行方式

### 环境要求
- Go 1.23.5 或更高版本
- 网络端口8080可用

### 编译和运行
```bash
# 下载依赖
go mod tidy

# 编译项目
go build -o store_server

# 运行服务器
./store_server
```

### 测试API
```bash
# 获取产品ID为1的产品信息
curl -X GET http://localhost:8080/product/1

# 获取产品ID为2的产品信息
curl -X GET http://localhost:8080/product/2
```

## 开发进度

### 已完成
- [x] 项目基础结构搭建
- [x] 产品数据模型定义
- [x] 产品查询API实现
- [x] 并发安全控制
- [x] 基础错误处理
- [x] 测试数据初始化

### 待开发
- [ ] 产品创建API
- [ ] 产品更新API
- [ ] 产品删除API
- [ ] 产品列表查询API
- [ ] 购物车模块
- [ ] 支付模块
- [ ] 仓库管理模块
- [ ] 数据库持久化
- [ ] 用户认证和授权
- [ ] 日志记录

## 注意事项

1. 当前版本使用内存存储，服务器重启后数据会丢失
2. 系统运行在调试模式，生产环境需要设置为发布模式
3. 当前版本没有实现数据验证和业务逻辑验证
4. 需要根据实际业务需求完善错误处理机制

## 贡献

本项目作为CS6650课程的教学项目，用于演示Go语言Web开发的基础概念和实践。
