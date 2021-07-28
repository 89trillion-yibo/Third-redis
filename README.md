# 技术文档


### 1.整体框架

管理员传来礼品码创建信息，创建礼品码结构，并且存入redis缓存数据库中保存，维护一个已领取次数和已领取用户列表，当用户输入礼品码时，根据礼品码类型，已领取次数和已领取用户列表判断用户是否有资格领取兑换码，如果有资格，则给用户添加礼品码内容的奖励



### 2.目录结构

```
├── app
│   ├── http
│   │   └── httpServer.go
│   └── main.go
├── create_test.go
├── go.mod
├── go.sum
├── internal
│   ├── ctrl
│   │   └── gifeCode.go
│   ├── gifeerror
│   │   └── error.go
│   ├── handler
│   │   ├── CreateGife.go
│   │   ├── IncreaseGife.go
│   │   └── JudgmentGife.go
│   ├── model
│   │   ├── ClaimList.go
│   │   ├── GiftCode.go
│   │   ├── User.go
│   │   └── redisDo.go
│   └── router
│       └── Routers.go
└── locust
    ├── locustCreateGife.py
    ├── locustGetGife.py
    ├── locustInsert.py
    ├── reportCreateGift.html
    ├── reportGetGiftCode.html
    └── reportIncrease.html

```



### 3.代码逻辑分层

| 层          | 文件夹                      | 主要职责             | 调用关系                  | 其它说明     |
| ----------- | --------------------------- | -------------------- | ------------------------- | ------------ |
| 应用层      | app/http/httpServer.go      | 服务器启动           | 调用路由层                | 不可同层调用 |
| 路由层      | app/http/httpServer.go      | 路由转发             | 被应用层调用，调用控制层  | 不可同层调用 |
| 控制层      | internal/ctrl/gifeCode.go   | 请求参数处理         | 被路由层调用，调用handler | 不可同层调用 |
| handler层   | internal/handler            | 处理具体业务         | 被控制层调用              | 不可同层调用 |
| 模型层      | /internal/model             | 数据模型、数据库操作 | 被handler、service调用    | 可同层调用   |
| globalError | internal/gifeerror/error.go | 统一异常处理         | 被控制层调用              | 不可同层调用 |



### 4.存储设计

礼品码信息

| 内容         | 数据库 | Key          |
| ------------ | ------ | ------------ |
| 礼品码类型   | Redis  | Type         |
| 礼品描述     | Redis  | Description  |
| 可领取次数   | Redis  | AllowedTimes |
| 有效期       | Redis  | ValidPeriod  |
| 礼品内容列表 | Redis  | Pack         |
| 创建人员     | Redis  | CreateName   |
| 创建时间     | Redis  | CreateTime   |

礼品内容

| 内容 | 数据库 | Key     |
| ---- | ------ | ------- |
| 金币 | Redis  | Gold    |
| 钻石 | Redis  | Diamond |

领取列表

| 内容     | 数据库 | Key  |
| -------- | ------ | ---- |
| 用户名字 | Redis  | Id   |
| 领取时间 | Redis  | Time |

已领取次数

| 内容       | 数据库 | Key                               |
| ---------- | ------ | --------------------------------- |
| 已领取次数 | Redis  | 对应的礼品码giftCode+后缀“Bytime” |



### 5.接口设计

#### 1.管理员创建礼品码

##### 请求方法

http POST

##### 接口地址

http://localhost:8080/creategif

##### 请求参数

| Key        | Value    |
| ---------- | -------- |
| gifType    | 3        |
| des        | 金币钻石 |
| allowTime  | 2        |
| valTime    | 10m      |
| createName | hyb      |
| gold       | 2000     |
| diamond    | 100      |

##### 请求响应

```json
{
    "Data": "QR1X9WBT",
    "Code": 200,
    "Message": "ok"
}
```



#### 2.管理员查询礼品码

##### 请求方法

http POST

##### 接口地址

http://localhost:8080/getGifcode

##### 请求参数

| Key     | Value    |
| ------- | -------- |
| gifcode | QR1X9WBT |

请求响应

```json
{
    "Data": {
        "bytimeInfo": "1",
        "gifeInfo": {
            "AllowedTimes": "2",
            "CreateName": "hyb",
            "CreateTime": "1627380442",
            "Description": "金币钻石",
            "Pack": "{\"Gold\":2000,\"Diamond\":100}",
            "Type": "2",
            "ValidPeriod": "1627381042"
        },
        "receiveInfo": {
            "2": "1627380591"
        }
    },
    "Code": 200,
    "Message": "ok"
}
```



#### 3.用户领取礼品码

##### 请求方法

http POST

##### 接口地址

http://localhost:8080/reward

##### 请求参数

| Key     | Value    |
| ------- | -------- |
| id      | 1        |
| gifcode | QR1X9WBT |

##### 请求响应

```
{
    "Data": "{\"Gold\":2000,\"Diamond\":100}",
    "Code": 200,
    "Message": "ok"
}
```

### 响应状态码

| 状态码 | 说明           |
| ------ | -------------- |
| 200    | 成功           |
| 101    | 获取参数错误   |
| 102    | 服务器异常     |
| 103    | 兑换码过期     |
| 104    | 领取码次数用尽 |
| 105    | 用户重复领取   |



### 6、第三方库

#### gin

```
go语言的web框架
https://github.com/gin-gonic/gin
```

#### go-redis

```
go语言连接操作redis数据库
https://github.com/go-redis
```



### 7、如何编译执行

```
go build ./app/main.go
```

运行可执行文件

```
./main
```



### 8.流程图

![未命名文件 (10)](https://user-images.githubusercontent.com/87186547/127321916-0dc97020-a216-4d7a-ae20-631e2226c2b7.jpg)

