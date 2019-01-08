package main

import (
	"google.golang.org/grpc"
	"ZCache/tool/logrus"
	pb "ZCache/zcache_rpc/zcacherpc"
	"context"
	"time"
	"io"
	"fmt"
)

func main(){
	fmt.Println("MMMM")
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		logrus.Warningf("fail to dial: %s", err.Error())
		fmt.Println("PPPPP",err)
	}
	defer conn.Close()
	client := pb.NewZacheProtoClient(conn)
	SetValue(client, pb.Data{Key:"YYY",Value:"ZZZ"})
	SetValue(client, pb.Data{Key:"UUU",Value:"OOO"})
	GetKey(client, pb.Data{Key:"UUU",Value:""})
	GetKeys(client, pb.Data{Key:"",Value:""})
	DeleteKey(client, pb.Data{Key:"UUU",Value:""})
	GetKeys(client, pb.Data{Key:"",Value:""})
}


func GetKey(client pb.ZacheProtoClient, data pb.Data)(resp *pb.Data, err error){
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer cancel()
	if resp, err = client.InGetKey(ctx,&data); err != nil{
		logrus.Warningf("GetValue Failed[data:%+v, err:%+s]",data, err.Error())
		return
	}

	return

}

func GetKeys(client pb.ZacheProtoClient, data pb.Data)(resp []pb.Data, err error){
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer cancel()
	stream, err := client.InGetKeys(ctx,&data)
	if err != nil{
		logrus.Warningf("GetValue Failed[data:%+v, err:%+s]",data, err.Error())
		return
	}
	for {
		serverData ,err := stream.Recv()
		if err == io.EOF{
			break
		}
		if err != nil{
			logrus.Warningf("stream.Recv Failed[Err:%s]",err.Error())
		}
		resp = append(resp, *serverData)
	}

	return
}

func SetValue(client pb.ZacheProtoClient, data pb.Data)(resp *pb.Data, err error){
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer cancel()
	if resp, err = client.InSetValue(ctx, &data); err != nil{
		logrus.Warningf("SetValue Failed[data:%+v, err:%+v]",data, err)
		return
	}
	return
}

func DeleteKey(client pb.ZacheProtoClient, data pb.Data)(resp *pb.Data, err error){
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer cancel()
	if resp, err = client.InDeleteKey(ctx, &data); err != nil{
		logrus.Warningf("SetValue Failed[data:%+v, err:%+s]",data, err.Error())
		return
	}
	return
}
