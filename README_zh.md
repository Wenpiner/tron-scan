# Tron-Scan
[English](README.md) | 中文

Tron-Scan 是一个基于 Go 语言开发的波场（TRON）区块链扫描服务，用于监控和分析 TRON 网络上的交易和账户活动。

## 功能特性

- 实时监控 TRON 网络交易
- 账户余额和交易历史查询
- 智能合约事件监听
- 消息队列集成（RabbitMQ）
- RESTful API 接口
- 分布式追踪支持（OpenTelemetry）
- 监控指标收集（Prometheus）

## 技术栈

- Go 1.21+
- go-zero 微服务框架
- RabbitMQ 消息队列
- OpenTelemetry 分布式追踪
- Prometheus 监控
- Docker 容器化

## 安装

### 前置要求

- Go 1.21 或更高版本
- Docker（可选）

### 本地安装

1. 克隆项目
```bash
git clone https://github.com/wenpiner/tron-scan.git
cd tron-scan
```

2. 安装依赖
```bash
go mod download
```

3. 编译项目
```bash
make build
```

### Docker 安装

```bash
docker build -t tron-scan .
docker run -d -p 8888:8888 tron-scan
```

## 配置

配置文件位于 `etc` 目录下，包含以下配置项：

- 服务端口设置
- TRON 节点连接信息
- RabbitMQ 连接配置
- 监控和追踪配置

## 使用

### 启动服务

```bash
./tron-scan -f etc/config.yaml
```

## 开发

### 项目结构

```
.
├── etc/           # 配置文件
├── internal/      # 内部代码
├── tron.go        # 主程序入口
├── tron.api       # API 定义文件
├── Dockerfile     # Docker 构建文件
└── Makefile       # 构建脚本
```

### 构建命令

- `make build`: 构建项目
- `make run`: 运行项目
- `make test`: 运行测试
- `make clean`: 清理构建文件

## 许可证

MIT License 