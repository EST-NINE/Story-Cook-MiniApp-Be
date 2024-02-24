# 2024 Story-Cook-MiniApp-Backend
使用 golang 编写

项目名：故事厨房(微信小程序版)
前端：请见[Heuluck](https://github.com/ncuhome/Story-Cook-FE)的仓库

## 项目运行
### 手动执行
**本项目使用`Go Mod`管理依赖。**

**下载依赖**
```shell
go mod tidy
```

**运行**
```shell
go run main.go
```

## 主要依赖
| 名称         | 版本      |
|------------|---------|
| golang     | 1.21.0  |
| gin        | v1.9.1  |
| gorm       | v1.9.16 |
| mysql      | v1.7.0  |
| jwt-go     | v3.2.0  |
| logrus     | v1.9.3  |

## 项目结构
```
├─config                 # 存放配置文件
├─controller             # 控制器层，处理请求和返回响应
├─middleware             # 中间件，用于处理请求前和响应后的逻辑
├─model
│  ├─dao                 # 数据访问对象，用于与数据库交互
│  ├─dto                 # 数据传输对象，用于在不同层之间传递数据
│  └─vo                  # 封装返回前端的数据
├─pkg
│  ├─myErrors            # 自定义错误处理包
│  └─util                # 工具函数包
├─router                 # 路由处理逻辑
└─service                # 服务层，实现接口函数的具体逻辑
```