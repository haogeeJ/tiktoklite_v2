package main

import (
	gtrace "TikTokLite_v2/common/grpc_jaeger"
	"TikTokLite_v2/video/dal"
	"TikTokLite_v2/video/pb"
	"TikTokLite_v2/video/remote_call"
	"TikTokLite_v2/video/service"
	"TikTokLite_v2/video/setting"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"strings"
)

var options = []grpc.ServerOption{
	grpc.MaxRecvMsgSize(135181930),
	grpc.MaxSendMsgSize(135181930),
}

func main() {
	if err := setting.Init("./config/config.yaml"); err != nil {
		fmt.Printf("init setting failed, err: %v \n", err)
		return
	}
	dal.Init()

	port := setting.Conf.Port
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("listen user rpc-service error", err)
	}
	//获取tracer
	tracer, _, err := gtrace.NewJaegerTracer(setting.Conf.Jaeger.ServiceName, setting.Conf.Jaeger.Host)
	if err != nil {
		fmt.Printf("new tracer err: %+vn", err)
		os.Exit(-1)
	}
	if tracer != nil {
		options = append(options, gtrace.ServerOption(tracer))
	}
	//调用其他rpc服务前的初始化，其实就是先和各rpc服务建立conn
	remote_call.Init()
	grpcServer := grpc.NewServer(options...)
	pb.RegisterVideoServiceServer(grpcServer, &VideoService{})
	reflection.Register(grpcServer)
	config := api.DefaultConfig()
	//设置consul服务的地址
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
		Name:    setting.Conf.Consul.ServerNames[0],
		Tags:    setting.Conf.Consul.ApiTags,
		Check: &api.AgentServiceCheck{
			Interval:                       setting.Conf.Consul.ApiHealthCheck.Interval,
			Timeout:                        setting.Conf.Consul.ApiHealthCheck.Timeout,
			GRPC:                           fmt.Sprintf("%s:%d", setting.Conf.Ip, port), //检查服务的地址请求
			DeregisterCriticalServiceAfter: setting.Conf.Consul.ApiHealthCheck.DeregisterCriticalServiceAfter,
		},
	}
	err = consulClient.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalf("注册服务失败: %s", err.Error())
	}

	fmt.Printf("服务启动成功;PORT:%d\n", port)

	service.BuildHotFeed(context.Background())
	service.UpdateUnLoginFeed(context.Background())
	//定时更新hotfeed和推送
	go func() {
		ticker := time.NewTicker(time.Minute * 30)
		defer ticker.Stop()
		for range ticker.C {
			go service.BuildHotFeed(context.Background())
			go service.CheckAliveUserAndPushHotFeed(context.Background())
			go service.UpdateUnLoginFeed(context.Background())
		}
	}()

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatal("grpcServe start error", err)
	}
}
