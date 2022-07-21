package main

import (
	gtrace "TikTokLite_v2/common/grpc_jaeger"
	"TikTokLite_v2/favorite_comment/dal"
	"TikTokLite_v2/favorite_comment/pb"
	"TikTokLite_v2/favorite_comment/remote_call"
	"TikTokLite_v2/favorite_comment/setting"
	"TikTokLite_v2/util"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	if err := setting.Init("./config/config.yaml"); err != nil {
		fmt.Printf("init setting failed, err: %v \n", err)
		return
	}
	dal.Init()
	util.FilterInit()
	port := setting.Conf.Port
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("listen user rpc-service error", err)
	}

	var options []grpc.ServerOption
	tracer, _, err := gtrace.NewJaegerTracer(setting.Conf.Jaeger.ServiceName, setting.Conf.Jaeger.Host)
	if err != nil {
		fmt.Printf("new tracer err: %+vn", err)
		os.Exit(-1)
	}
	if tracer != nil {
		options = append(options, gtrace.ServerOption(tracer))
	}
	remote_call.Init()
	grpcServer := grpc.NewServer(options...)
	pb.RegisterCommentServiceServer(grpcServer, &CommentService{})
	pb.RegisterFavoriteServiceServer(grpcServer, &FavoriteService{})
	pb.RegisterHotFeedServer(grpcServer, &HotFeedService{})
	pb.RegisterVideoInfoServer(grpcServer, &VideoInfoService{})
	reflection.Register(grpcServer)

	config := api.DefaultConfig()

	config.Address = setting.Conf.Consul.Host
	consulClient, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("连接consul失败: %s", err.Error())
	}

	// grpc注册服务的健康检查
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())

	registration := &api.AgentServiceRegistration{
		Address: setting.Conf.Consul.ApiAddress,
		Port:    setting.Conf.Consul.ApiPort,
		ID:      fmt.Sprintf("%s", strings.ReplaceAll(uuid.New().String(), "-", "")),
		Name:    setting.Conf.Consul.ServerNames[2],
		Tags:    setting.Conf.Consul.ApiTags,
		Check: &api.AgentServiceCheck{
			Interval:                       setting.Conf.Consul.ApiHealthCheck.Interval,
			Timeout:                        setting.Conf.Consul.ApiHealthCheck.Timeout,
			GRPC:                           fmt.Sprintf("%s:%d", setting.Conf.Ip, port),
			DeregisterCriticalServiceAfter: setting.Conf.Consul.ApiHealthCheck.DeregisterCriticalServiceAfter,
		},
	}
	err = consulClient.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalf("注册服务失败: %s", err.Error())
	}

	fmt.Printf("服务启动成功;PORT:%d\n", port)

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatal("favorite_comment_service start error", err)
	}
}
