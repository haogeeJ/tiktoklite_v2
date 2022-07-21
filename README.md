# tiktoklite_v2
![tiktoklite架构图](https://user-images.githubusercontent.com/81409707/180145201-ecd73308-6b56-4fba-a377-e762836ea236.png)

服务划分：


技术方案：
- web框架：gin，接收HTTP接口请求，然后再调用GRPC服务
- rpc框架：grpc
- 数据库：mysql(gorm)，redis(async-redigo)
- 链路追踪：jaeger+opentracing 集成gorm，grpc，gin，redis
- 服务治理：consul，主要是rpc服务注册，发现和健康检查
- 数据库异构中间件：Bifrost，可以实时进行MYSQL->MYSQL的同步,目是为了让每个服务的数据库可以独立，解耦。为数据库添加需要的冗余数据，就避免了和其他grpc服务数据库的耦合。Bifrost的优点是操作简单，不会对业务代码造成入侵。
- NPS内网穿透：由于我的consul服务是部署在云服务器，而项目服务是在本地内网，所以需要通过内网穿透，consul才能对内网服务进行监控检查。
- 日志：logrus
