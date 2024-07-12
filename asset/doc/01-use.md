## 初始化

使用项目之前需要初始化

### 克隆代码

```shell
git clone https://github.com/vperson/gin-api.git <你的项目名字>
cd <你的项目名字>
```

### 修改项目名

由于默认的项目名字是gin-api，需要修改成<你的项目名字>

##### mac os && Linux

```shell
/bin/bash asset/bash/osx_linux.sh <你的项目名字>
```

#### windows

```
win.bat <你的项目名字>
```

### 项目结构

```shell
.
├── asset
│   ├── bash
│   └── doc
├── cmd
├── config
├── pkg
│   ├── logger
│   └── util
├── repository
│   ├── cache
│   └── store
│       └── tb
└── server
    ├── controller
    │   └── actuator_ctrl
    ├── middleware
    └── router
        ├── api
        │   └── v1
        └── restful
```

config: <开发者需要关注> 配置

pkg/logger: 日志，直接使用

server/controller/<feature>_ctrl: <开发者需要关注> 业务控制层逻辑，可以根据功能的不同增加不同的控制器

server/router: <开发者需要关注> 定义路由

server/router/api/v1: <开发者需要关注> 路由层

repository/store: <开发者需要关注> 存储层，对数据库的操作
