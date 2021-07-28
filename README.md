# 技术文档


### 1.整体架构

主要流程是，客户端调用注册或登录接口，查询mongodb数据库，查询数据，如果有用户则返回数据，如果是新用户则注册，给用户增加奖励复用第三题代码，从redis中查询礼品码信息和用户列表，已领取次数等信息，给mongodb数据库中的用户添加奖励



### 2.目录结构

```
├── app
│   ├── http
│   │   └── httpServer.go
│   └── main.go
├── go.mod
├── go.sum
├── increase_test.go
├── internet
│   ├── ctrl
│   │   ├── gifeCtrl.go
│   │   └── userCtrl.go
│   ├── gifeerror
│   │   └── error.go
│   ├── handler
│   │   ├── createUser.go
│   │   └── increase.go
│   ├── model
│   │   ├── claimList.go
│   │   ├── gifecode.go
│   │   ├── mongodbCon.go
│   │   ├── mongodbDo.go
│   │   ├── redisCon.go
│   │   ├── redisDo.go
│   │   └── user.go
│   ├── router
│   │   └── routers.go
│   └── service
│       ├── JudgmentUser.go
│       ├── gifeService.go
│       └── userService.go
├── locust.py
├── report.html
└── response
    ├── generaReward.pb.go
    └── generaReward.proto

```



### 3.代码逻辑分层

| 层        | 文件夹                      | 主要职责           | 调用关系                    | 其它说明     |
| --------- | --------------------------- | ------------------ | --------------------------- | ------------ |
| 应用层    | app/http/httpServer.go      | 启动服务器         | 调用路由层                  | 不可同层调用 |
| 路由层    | internet/router/Routers.go  | 路由转发           | 被应用层调用，调用控制层    | 不可同层调用 |
| 控制层    | internet/ctrl/UserCtrl.go   | 请求参数处理，响应 | 被路由层调用，调用handler   | 不可同层调用 |
| handler层 | internet/handler            | 处理具体业务       | 被控制层调用                | 不可同层调用 |
| 模型层    | internet/mondel             | 构建数据类型       | 被handler调用               | 可同层调用   |
| gifeerror | internet/gifeerror/error.go | 统一异常处理       | 被ctrl调用                  | 不可同层调用 |
| service层 | internet/service            | 通用业务处理       | 被控制层调用，被handler调用 | 不可同层调用 |



### 4.存储设计

用户结构设计

| 内容                | 数据库  | Key     |
| ------------------- | ------- | ------- |
| 用户名id            | Mongodb | Id      |
| mongodb自动生成的id | Mongodb | _id     |
| 金币                | Mongodb | Gold    |
| 钻石                | Mongodb | Diamond |



礼包结构

| 内容 | 数据库 | Key     |
| ---- | ------ | ------- |
| 金币 | Redis  | Gold    |
| 钻石 | Redis  | Diamond |



### 5、接口设计

#### 1、登陆注册

##### 请求方法

http POST

##### 接口地址

localhost:8081/getUser

##### 请求参数

| Key  | Value |
| ---- | ----- |
| Uuid | 6     |

##### 请求响应

```
{
    "Data": {
        "id": "3",
        "gold": 6000,
        "diamond": 300
    },
    "Code": 200,
    "Message": "已有用户或新用户注册"
}
```



#### 2、用户增加奖励

##### 请求方法

http POST

##### 接口地址

localhost:8081/increase

##### 请求参数

| Key     | Value    |
| ------- | -------- |
| Uuid    | 6        |
| Gifcode | 22N9O396 |

##### 请求响应

```
{
    "Data": "CMgBEjTkv6Hmga/ov5Tlm57mraPnoa4sMTAwMeS7o+ihqOmHkeW4gSwxMDAy5Luj6KGo6ZK755+zGgYI6QcQuBcaBgjqBxCsAiIFCOkHEAAiBQjqBxAAKgYI6gcQrAIqBgjpBxC4FzIM5omp5bGV5a2X5q61",
    "Code": 200,
    "Message": "ok"
}
```



### 6、第三方库

### gin

```
https://github.com/gin-gonic/gin
```

### Mongoldb go

```
https://go.mongodb.org/mongo-driver/mongo
```

### protobuf

```
http://github.com/golang/protobuf/proto
```





### 7、如何编译执行

进入app目录编译

```
go build main.go
```

运行可执行文件

```
./main
```





### 8.流程图

![未命名文件 (11)](https://user-images.githubusercontent.com/87186547/127321642-c3842905-3ea9-4737-8f4f-3ac1eabde1dc.jpg)
