package client

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"master/scaffold_master/etcd"
	lightweight "master/scaffold_master/pb"
	"time"
)

func client_user() {
	newResolver := etcd.NewResolver([]string{
		"127.0.0.1:2379",
	}, "user_etcd")
	resolver.Register(newResolver)

	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	addr := fmt.Sprintf("%s:///%s", newResolver.Scheme(), "g.srv.mail" /*user_etcd经测试，这个可以随便写，底层只是取scheme对应的Build对象*/)

	dialContext, err := grpc.DialContext(ctx, addr, grpc.WithInsecure(),
		//指定初始化round_robin => balancer (后续可以自行定制balancer和 register、resolver 同样的方式)
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
		grpc.WithBlock())

	// 这种方式也行
	//conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBalancerName("round_robin"))
	//conn, err := grpc.Dial(":8972", grpc.WithInsecure())

	if err != nil {
		log.Error("grpc dial err:", err)
		return
	}

	/*conn, err := grpc.Dial(
	        fmt.Sprintf("%s://%s/%s", "consul", GetConsulHost(), s.Name),
	        //不能block => blockkingPicker打开，在调用轮询时picker_wrapper => picker时若block则不进行robin操作直接返回失败
	        //grpc.WithBlock(),
	        grpc.WithInsecure(),
	        //指定初始化round_robin => balancer (后续可以自行定制balancer和 register、resolver 同样的方式)
	        grpc.WithBalancerName(roundrobin.Name),
	        //grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
	    )
		if err != nil {
		panic(err)
	}

	*/
	send_main(lightweight.NewUserServiceClient(dialContext))
	set_user(lightweight.NewUserServiceClient(dialContext))
}

// 测试发送邮件
func send_main(cc lightweight.UserServiceClient) {
	mail, err := cc.SendMail(context.TODO(), &lightweight.MailRequest{
		Mail: "qq@mail.com",
		Text: "test,test",
	})

	if err != nil {
		log.Error("grpc sendmail err", err)
		return

	}
	log.Info(mail)
}

// 获取用户信息
func set_user(cc lightweight.UserServiceClient) {
	info, err := cc.GetUserInfo(context.TODO(), &lightweight.User{
		Id:    1,
		Name:  "红尘",
		Age:   18,
		Phone: "110",
		Sex:   "男",
		Addr:  "成都",
	})

	if err != nil {
		log.Error("grpc set user err=", err)
		return
	}
	log.Info(info)
}
