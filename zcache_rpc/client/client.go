package client

import (
	"google.golang.org/grpc"
	"zCache/tool/logrus"
	pb "zCache/zcache_rpc/zcacherpc"
	"context"
	"time"
	"io"
	"fmt"
)

func main(){
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


func PreDecr(client pb.ZacheProtoClient, data pb.Data,commitID chan string)(err error){
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer cancel()
	resp, err := client.PreDecr(ctx,&data);
	if err != nil{
		logrus.Warningf("PreDecr Failed[data:%+v, err:%+s]",data, err.Error())
		return
	}

	commitID <- resp.CommitID

	return

}
func PreDecrBy(client pb.ZacheProtoClient, data pb.Data,commitID chan string)(err error){
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer cancel()
	resp, err := client.PreDecrBy(ctx,&data);
	if err != nil{
		logrus.Warningf("PreDecrBy Failed[data:%+v, err:%+s]",data, err.Error())
		return
	}

	commitID <- resp.CommitID

	return

}

func PreIncr(client pb.ZacheProtoClient, data pb.Data,commitID chan string)(err error){
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer cancel()
	resp, err := client.PreIncr(ctx,&data);
	if err != nil{
		logrus.Warningf("PreIncr Failed[data:%+v, err:%+s]",data, err.Error())
		return
	}

	commitID <- resp.CommitID

	return

}
func PreIncrBy(client pb.ZacheProtoClient, data pb.Data,commitID chan string)(err error){
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer cancel()
	resp, err := client.PreIncrBy(ctx,&data);
	if err != nil{
		logrus.Warningf("PreIncrBy Failed[data:%+v, err:%+s]",data, err.Error())
		return
	}

	commitID <- resp.CommitID

	return

}

func PreSet(client pb.ZacheProtoClient, data pb.Data,commitID chan string)(err error){
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer cancel()
	resp, err := client.PreSetValue(ctx,&data);
	if err != nil{
		logrus.Warningf("PreSetValue Failed[data:%+v, err:%+s]",data, err.Error())
		return
	}

	commitID <- resp.CommitID

	return

}

func PreDelete(client pb.ZacheProtoClient, data pb.Data,commitID chan string)(err error){
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer cancel()
	resp, err := client.PreDeleteKey(ctx,&data);
	if err != nil{
		logrus.Warningf("PreDeleteKey Failed[data:%+v, err:%+s]",data, err.Error())
		return
	}

	commitID <- resp.CommitID

	return

}
func PreUpdate(client pb.ZacheProtoClient, data pb.Data,commitID chan string)(err error){
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer cancel()
	resp, err := client.PreUpdateValue(ctx,&data);
	if err != nil{
		logrus.Warningf("PreDUpdate Failed[data:%+v, err:%+s]",data, err.Error())
		return
	}

	commitID <- resp.CommitID

	return

}

func CommitJob(client pb.ZacheProtoClient, data pb.CommitIDMsg)(err error){
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer cancel()
	_, err = client.Commit(ctx,&data);
	if err != nil{
		logrus.Warningf("Commit Failed[data:%+v, err:%+s]",data, err.Error())
		return
	}
	return

}

func DropJob(client pb.ZacheProtoClient, data pb.CommitIDMsg)(err error){
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
	defer cancel()
	_, err = client.Drop(ctx,&data);
	if err != nil{
		logrus.Warningf("Drop Failed[data:%+v, err:%+s]",data, err.Error())
		return
	}


	return

}