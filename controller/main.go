package main

import (
	tracer "TikTokLite_v2/common/grpc_jaeger"
	"TikTokLite_v2/controller/file"
	//"TikTokLite_v2/controller/tracer"

	"TikTokLite_v2/controller/favorite_comment"
	"TikTokLite_v2/controller/follow"
	"TikTokLite_v2/controller/middleware"
	"TikTokLite_v2/controller/remote_call"
	"TikTokLite_v2/controller/setting"

	"TikTokLite_v2/controller/user"
	"TikTokLite_v2/controller/video"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing-contrib/go-gin/ginhttp"
	"net/http"
)

func main() {
	if err := setting.Init("./config/config.yaml"); err != nil {
		fmt.Printf("init setting failed, err: %v \n", err)
		return
	}
	tracer, closer, err := tracer.NewJaegerTracer(setting.Conf.Jaeger.ServiceName, setting.Conf.Jaeger.Host)
	if err != nil {
		panic(err)
	}
	defer closer.Close()
	remote_call.Init()
	r := gin.Default()

	r.GET("/static/videos", file.Videos)
	r.GET("/static/covers", file.Covers)

	apiRouter := r.Group("/douyin")
	//实现了用户注册，登录，信息的接口
	jaegerMiddle := ginhttp.Middleware(tracer, ginhttp.OperationNameFunc(func(r *http.Request) string {
		return fmt.Sprintf("HTTP %s %s", r.Method, r.URL.String())
	}))
	apiRouter.Use(jaegerMiddle)
	// basic apis
	apiRouter.GET("/feed/", video.Feed)
	apiRouter.GET("/user/", user.UserInfo)

	apiRouter.POST("/user/register/", user.Register)

	apiRouter.POST("/user/login/", user.Login)
	publishGroup := apiRouter.Group("/publish", middleware.ValidDataTokenMiddleWare)
	publishGroup.POST("/action/", video.Publish)
	publishGroup.GET("/list/", video.PublishList)

	// extra apis - I
	favoriteGroup := apiRouter.Group("/favorite", middleware.ValidDataTokenMiddleWare)
	favoriteGroup.POST("/action/", favorite_comment.FavoriteAction)
	favoriteGroup.GET("/list/", favorite_comment.FavoriteList)

	commentGroup := apiRouter.Group("/comment", middleware.ValidDataTokenMiddleWare)
	commentGroup.POST("/action/", favorite_comment.CommentAction)
	commentGroup.GET("/list/", favorite_comment.CommentList)

	// extra apis - II

	apiRouter.POST("/relation/action/", middleware.ValidDataTokenMiddleWare, follow.RelationAction)
	apiRouter.GET("/relation/follow/list/", middleware.ValidDataTokenMiddleWare, follow.FollowList)
	apiRouter.GET("/relation/follower/list/", middleware.ValidDataTokenMiddleWare, follow.FollowerList)
	r.Run(":8080")
}
