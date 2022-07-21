# tiktoklite_v2
![tiktoklite架构图](https://user-images.githubusercontent.com/81409707/180145201-ecd73308-6b56-4fba-a377-e762836ea236.png)
目录结构树如下
├─common
│  │  gen_port.go
│  │  
│  ├─gorm_tracing
│  │      gorm_tracing.go
│  │      
│  ├─grpc_jaeger
│  │      wrapper.go
│  │      
│  ├─redis_tracing
│  │      redis_tracing.go
│  │      
│  └─service
│          build_resp.go
│          
├─controller
│  │  main.go
│  │  
│  ├─common
│  │      common.go
│  │      
│  ├─config
│  │      config.yaml
│  │      
│  ├─favorite_comment
│  │      comment.go
│  │      favorite.go
│  │      
│  ├─file
│  │      file.go
│  │      
│  ├─follow
│  │      follow.go
│  │      
│  ├─middleware
│  │      jwt.go
│  │      
│  ├─remote_call
│  │  │  client_init.go
│  │  │  
│  │  ├─call_fav_com
│  │  │      favorite_comment.go
│  │  │      
│  │  ├─call_user_follow
│  │  │      user_follow.go
│  │  │      
│  │  └─call_video
│  │          video.go
│  │          
│  ├─setting
│  │      setting.go
│  │      
│  ├─tracer
│  │      tracer.go
│  │      
│  ├─user
│  │      user.go
│  │      
│  └─video
│          feed.go
│          publish.go
│          
├─favorite_comment
│  │  handler.go
│  │  main.go
│  │  
│  ├─config
│  │      config.yaml
│  │      
│  ├─dal
│  │  │  comment.go
│  │  │  comment_redis.go
│  │  │  favorite.go
│  │  │  favorite_redis.go
│  │  │  init.go
│  │  │  read_slave.go
│  │  │  
│  │  ├─db
│  │  │      init.go
│  │  │      
│  │  └─redb
│  │          init.go
│  │          
│  ├─pb
│  │      favorite_comment.pb.go
│  │      
│  ├─remote_call
│  │  │  client_init.go
│  │  │  
│  │  ├─call_fav_com
│  │  │      favorite_comment.go
│  │  │      
│  │  ├─call_user_follow
│  │  │      user_follow.go
│  │  │      
│  │  └─call_video
│  │          video.go
│  │          
│  ├─service
│  │      comment.go
│  │      favorite.go
│  │      hotfeed.go
│  │      videoInfo.go
│  │      
│  └─setting
│          setting.go
│          
├─idl
│      favorite_comment.proto
│      follow.proto
│      user.proto
│      video.proto
│      
├─user_follow
│  ├─follow
│  │  │  handler.go
│  │  │  
│  │  ├─dal
│  │  │      follow.go
│  │  │      
│  │  ├─pb
│  │  │      follow.pb.go
│  │  │      
│  │  └─service
│  │          follow.go
│  │          
│  ├─setting
│  │      setting.go
│  │      
│  └─user
│      │  handler.go
│      │  main.go
│      │  
│      ├─config
│      │      config.yaml
│      │      
│      ├─dal
│      │  │  init.go
│      │  │  user.go
│      │  │  
│      │  ├─db
│      │  │      init.go
│      │  │      
│      │  └─redb
│      │          init.go
│      │          
│      ├─pb
│      │      user.pb.go
│      │      
│      ├─remote_call
│      │  │  client_init.go
│      │  │  
│      │  ├─call_fav_com
│      │  │      favorite_comment.go
│      │  │      
│      │  ├─call_user_follow
│      │  │      user_follow.go
│      │  │      
│      │  └─call_video
│      │          video.go
│      │          
│      └─service
│              user.go
│              
├─util
│  │  sensitive_test.go
│  │  sensitive_word.go
│  │  time.go
│  │  time_test.go
│  │  token.go
│  │  uuid.go
│  │  
│  ├─sensitive_txt
│  │      sensitive_word.txt
│  │      
│  └─trace_id_log
│      │  gin_logger.go
│      │  gorm_logger.go
│      │  
│      ├─loggers
│      │      loggers.go
│      │      
│      └─middleware
│              middleware.go
│              
└─video
    │  handler.go
    │  main.go
    │  
    ├─config
    │      config.yaml
    │      
    ├─dal
    │  │  hotfeed.go
    │  │  init.go
    │  │  video.go
    │  │  
    │  ├─db
    │  │      init.go
    │  │      
    │  └─redb
    │          init.go
    │          
    ├─pb
    │      video.pb.go
    │      
    ├─remote_call
    │  │  client_init.go
    │  │  
    │  ├─call_fav_com
    │  │      favorite_comment.go
    │  │      
    │  ├─call_user_follow
    │  │      user_follow.go
    │  │      
    │  └─call_video
    │          video.go
    │          
    ├─service
    │      feed.go
    │      publish.go
    │      
    └─setting
            setting.go
            
