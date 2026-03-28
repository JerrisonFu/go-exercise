# Go Exercise

一个从基础到高级的 Go 语言学习项目，涵盖语法、并发、网络编程、Web 开发等核心知识点。

## 📚 内容概览

| 模块 | 说明 |
|------|------|
| [basics](./basics/) | 基础语法：变量、数据类型、控制流程 |
| [functions](./functions/) | 函数与错误处理 |
| [oop](./oop/) | 结构体、接口、面向对象 |
| [concurrency](./concurrency/) | 并发编程：goroutine 和 channel |
| [stdlib](./stdlib/) | 标准库常用模块 |
| [network](./network/) | 网络编程：HTTP、TCP |
| [database](./database/) | 数据库操作（SQLite） |
| [web](./web/) | Web 框架开发（Gin） |
| [api](./api/) | RESTful API 设计 |
| [testing](./testing/) | 单元测试与性能优化 |

## 🚀 快速开始

### 环境要求

- Go 1.21+

### 安装依赖

```bash
go mod download
```

### 运行所有示例

```bash
go run exercise.go
```

### 运行测试

```bash
# 运行所有测试
go test ./...

# 运行测试并显示详情
go test ./... -v

# 运行性能测试
go test ./... -bench=.
```

## 📖 模块详解

### 1. Basics (基础语法)

学习 Go 语言的基础概念：

- 变量声明与常量
- 基本数据类型（int, float, string, bool）
- 复合数据类型（array, slice, map, struct）
- 控制流程（if, for, switch）
- 运算符与指针

### 2. Functions (函数与错误处理)

掌握函数的高级用法：

- 多返回值函数
- 变参函数
- 闭包与高阶函数
- defer 延迟执行
- panic 与 recover
- 自定义错误类型

### 3. OOP (面向对象)

Go 语言的面向对象特性：

- 结构体定义与方法
- 值接收器与指针接收器
- 接口定义与实现
- 多态与类型断言
- 组合与嵌入

### 4. Concurrency (并发编程)

Go 的核心特性：

- goroutine 基础
- channel 通道
- buffered channel
- select 多路复用
- sync.WaitGroup
- sync.Mutex 互斥锁
- 死锁与竞态条件

### 5. Stdlib (标准库)

常用标准库模块：

- strings 字符串操作
- json 编解码
- os 文件操作
- time 时间处理
- strconv 类型转换
- bytes 字节操作

### 6. Network (网络编程)

网络通信基础：

- HTTP 服务端与客户端
- HTTP 方法（GET, POST, PUT, DELETE）
- 请求头与响应头
- TCP 套接字
- UDP 数据报
- DNS 解析

### 7. Database (数据库)

数据库操作：

- SQLite 连接
- CRUD 操作
- 事务处理
- 预处理语句
- 错误处理

### 8. Web (Gin 框架)

流行的 Web 框架：

- 路由配置
- 请求参数绑定
- JSON 响应
- 中间件
- 参数验证

### 9. API (RESTful 设计)

RESTful API 最佳实践：

- REST 风格路由设计
- 认证中间件
- CORS 跨域处理
- 分页实现
- 统一响应格式
- 日志中间件

### 10. Testing (测试与优化)

测试与性能分析：

- 单元测试
- 基准测试 (Benchmark)
- 子测试 (SubTest)
- Table-Driven 测试
- Mock 对象
- 常用算法实现

## 📁 项目结构

```
exercise/
├── basics/           # 基础语法
│   ├── basics.go
│   ├── variables.go
│   ├── datatypes.go
│   ├── controlflow.go
│   └── operators.go
├── functions/         # 函数与错误处理
│   ├── basic.go
│   ├── advanced.go
│   └── error_handling.go
├── oop/             # 面向对象
│   ├── struct.go
│   ├── interface.go
│   └── polymorphism.go
├── concurrency/     # 并发编程
│   ├── goroutine.go
│   ├── channel.go
│   └── sync.go
├── stdlib/          # 标准库
│   ├── strings.go
│   ├── json.go
│   ├── file.go
│   └── time.go
├── network/         # 网络编程
│   ├── http.go
│   └── tcp.go
├── database/        # 数据库
│   └── db.go
├── web/            # Web 框架
│   └── handlers.go
├── api/            # RESTful API
│   └── restful.go
├── testing/        # 测试
│   ├── testing.go
│   └── benchmark_test.go
├── exercise.go     # 主入口
├── go.mod
└── README.md
```

## 🧪 测试覆盖率

| 模块 | 测试用例数 |
|------|-----------|
| basics | 8 |
| functions | 16 |
| oop | 18 |
| concurrency | 12 |
| stdlib | 26 |
| network | 8 |
| database | 7 |
| web | 9 |
| api | 12 |
| testing | 7 |
| **总计** | **123** |

## 🔧 开发指南

### 添加新模块

1. 创建新目录：`mkdir -p newmodule`
2. 添加代码文件：`newmodule/newmodule.go`
3. 添加测试文件：`newmodule/newmodule_test.go`
4. 更新 `exercise.go` 引入新模块

### 运行特定模块测试

```bash
go test ./basics/... -v
go test ./functions/... -v
```

## 📚 学习路径

```
1. basics        → 变量、数据类型、控制流
2. functions     → 函数、错误处理
3. oop           → 结构体、接口
4. concurrency   → goroutine、channel
5. stdlib        → 标准库实战
6. network       → HTTP、TCP
7. database      → SQLite CRUD
8. web           → Gin 框架
9. api           → RESTful 设计
10. testing      → 测试与优化
```

## 📄 许可证

MIT License

## 🙏 致谢

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [modernc.org/sqlite](https://modernc.org/sqlite)
