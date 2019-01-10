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
	"strconv"
)


type ZcacheServer struct {
	savedData []*pb.Data
	mu sync.Mutex
}

func (s *ZcacheServer) InGetKey(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	fmt.Println("GetValue",data)
	node ,err := zdata.CoreGet(data.Key)
	if err != nil {
		logrus.Warningf("CoreGet Failed[Err:%s]",err.Error())
		return nil, err
	}
	resp = &pb.Data{Key:node.Key,Value:node.Value}
	return
}

func (s *ZcacheServer) InGetKeys(data *pb.Data, stream pb.ZacheProto_InGetKeysServer)(err error){
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

func (s *ZcacheServer) InSetValue(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	_ ,err = zdata.CoreAdd(data.Key,data.Value)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{Key:data.Key,Value:data.Value}
	}
	return
}


func (s *ZcacheServer) InDeleteKey(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	_ ,err = zdata.CoreDelete(data.Key)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{Key:data.Key,Value:data.Value}
	}
	return
}

func (s *ZcacheServer) InExport(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	err = zdata.CoreFlush()
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{}
	}

	return nil , nil
}

func (s *ZcacheServer) InImport(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	err = zdata.CoreImport()
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{}
	}

	return nil , nil
}

func (s *ZcacheServer) InDeleteKeys(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	err = zdata.CoreDeleteAll()
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{}
	}
	return nil , nil
}

func (s *ZcacheServer) InExpension(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	size, err := strconv.ParseInt(data.Value,10,64)
	if err != nil {
		return nil ,err
	}
	err = zdata.CoreExpension(size)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{}
	}

	return nil , nil
}

func (s *ZcacheServer) InGetKeyNum(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	num, err := zdata.CoreGetKeyNum()
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{Value:string(num)}
	}
	return nil , nil
}

func (s *ZcacheServer) InKeyIncr(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	_, err = zdata.CoreInDecr(data.Key,data.Value)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{}
	}

	return nil , nil
}

func (s *ZcacheServer) InKeyIncrBy(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	_, err = zdata.CoreInDecr(data.Key,data.Value)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{}
	}
	return nil , nil
}


func (s *ZcacheServer) InKeyDecr(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	_, err = zdata.CoreInDecr(data.Key,data.Value)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{}
	}
	return nil , nil
}

func (s *ZcacheServer) InKeyDecrBy(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	_, err = zdata.CoreInDecr(data.Key,data.Value)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{}
	}
	return nil , nil
}

func (s *ZcacheServer) InCommit(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	commitID, err := strconv.ParseInt(data.Value,10,64)
	if err != nil {
		return nil ,err
	}
	err = zdata.CoreCommit(commitID)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{}
	}

	return nil , nil
}

func (s *ZcacheServer) InDrop(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	commitID, err := strconv.ParseInt(data.Value,10,64)
	if err != nil {
		return nil ,err
	}
	err = zdata.CoreDrop(commitID)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{}
	}
	return nil , nil
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
