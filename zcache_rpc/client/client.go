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
	GetValue(client, pb.Data{Key:"UUU",Value:""})
	GetValues(client, pb.Data{Key:"",Value:""})
	DeleteValue(client, pb.Data{Key:"UUU",Value:""})
	GetValues(client, pb.Data{Key:"",Value:""})
}


func GetValue(client pb.ZacheProtoClient, data pb.Data)(resp *pb.Data, err error){
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer cancel()
	if resp, err = client.GetValue(ctx,&data); err != nil{
		logrus.Warningf("GetValue Failed[data:%+v, err:%+s]",data, err.Error())
		return
	}

	fmt.Println("AAA",resp)
	return

}

func GetValues(client pb.ZacheProtoClient, data pb.Data)(resp []pb.Data, err error){
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer cancel()
	stream, err := client.GetValues(ctx,&data)
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

	fmt.Println("BBB",resp)
	return
}

func SetValue(client pb.ZacheProtoClient, data pb.Data)(resp *pb.Data, err error){
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer cancel()
	if resp, err = client.SetValue(ctx, &data); err != nil{
		fmt.Println("TTTTT",err)
		logrus.Warningf("SetValue Failed[data:%+v, err:%+v]",data, err)
		return
	}
	return
}

func DeleteValue(client pb.ZacheProtoClient, data pb.Data)(resp *pb.Data, err error){
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer cancel()
	if resp, err = client.DeleteValue(ctx, &data); err != nil{
		logrus.Warningf("SetValue Failed[data:%+v, err:%+s]",data, err.Error())
		return
	}
	return
}
