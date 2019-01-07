package server


import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "ZCache/zcache_rpc/zcacherpc"
	"sync"
	"github.com/sirupsen/logrus"
	"ZCache/data"
	"fmt"
)


type ZcacheServer struct {
	savedData []*pb.Data
	mu sync.Mutex
}

func (s *ZcacheServer) GetValue(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	fmt.Println("GetValue",data)
	node ,err := zdata.CoreGet(data.Key)
	if err != nil {
		logrus.Warningf("CoreGet Failed[Err:%s]",err.Error())
		return nil, err
	}
	resp = &pb.Data{Key:node.Key,Value:node.Value}
	return
}

func (s *ZcacheServer) GetValues(data *pb.Data, stream pb.ZacheProto_GetValuesServer)(err error){
	fmt.Println("GetValues",data)
	head ,err := zdata.CoreGetAll()
	if err != nil {
		logrus.Warningf("CoreGet Failed[Err:%s]",err.Error())
		return  err
	}
	for head != nil{
		if err = stream.Send(&pb.Data{Key:head.Key,Value:head.Value}); err != nil {
			logrus.Warnf("stream.Send Failed[Data:%+v, Err:%s]",head, err.Error())
			continue
		}
		head = head.Next
	}
	return  nil
}

func (s *ZcacheServer) SetValue(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	fmt.Println("SetValue",data)
	_ ,err = zdata.CoreAdd(data.Key,data.Value)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{Key:data.Key,Value:data.Value}
	}
	return
}


func (s *ZcacheServer) DeleteValue(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	fmt.Println("DeleteValue",data)
	_ ,err = zdata.CoreDelete(data.Key)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{Key:data.Key,Value:data.Value}
	}
	return
}

func GrpcInit(){
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatal("failed to listen: %s", err.Error())
		return
	}

	s := grpc.NewServer()
	pb.RegisterZacheProtoServer(s, &ZcacheServer{})

	s.Serve(lis)
}
