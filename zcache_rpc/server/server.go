package server


import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "zCache/zcache_rpc/zcacherpc"
	"sync"
	"github.com/sirupsen/logrus"
	"zCache/data"
	"fmt"
	"strconv"
	"zCache/routes"
	"zCache/tool"
	"zCache/types"
	"errors"
	"zCache/global"
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
	return
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

	return
}

func (s *ZcacheServer) InImport(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	err = zdata.CoreImport()
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{}
	}

	return
}

func (s *ZcacheServer) InDeleteKeys(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	err = zdata.CoreDeleteAll()
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{}
	}
	return
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

	return
}

func (s *ZcacheServer) InGetKeyNum(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	num, err := zdata.CoreGetKeyNum()
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{Value:string(num)}
	}
	return
}

func (s *ZcacheServer) InKeyIncr(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	_, err = zdata.CoreInDecr(data.Key,data.Value)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{}
	}

	return
}

func (s *ZcacheServer) InKeyIncrBy(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	_, err = zdata.CoreInDecr(data.Key,data.Value)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{}
	}
	return
}


func (s *ZcacheServer) InKeyDecr(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	_, err = zdata.CoreInDecr(data.Key,data.Value)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{}
	}
	return
}

func (s *ZcacheServer) InKeyDecrBy(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	_, err = zdata.CoreInDecr(data.Key,data.Value)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{}
	}
	return
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

	return
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
	return
}
func (s *ZcacheServer) Get(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	fmt.Println("GetValue",data)
	value ,err := routes.Get(data.Key)
	if err != nil {
		logrus.Warningf("CoreGet Failed[Err:%s]",err.Error())
		return nil, err
	}
	resp = &pb.Data{Key:data.Key,Value:value}
	return
}
func (s *ZcacheServer) DeleteKey(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	err = routes.Delete(data.Key)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{Key:data.Key,Value:data.Value}
	}
	return
}
func (s *ZcacheServer) SetValue(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	err = routes.Set(data.Key,data.Value)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{Key:data.Key,Value:data.Value}
	}
	return
}
func (s *ZcacheServer) UpdateValue(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	err = routes.Update(data.Key,data.Value)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{Key:data.Key,Value:data.Value}
	}
	return
}
func (s *ZcacheServer) GetKeys(data *pb.Data, stream pb.ZacheProto_GetKeysServer)(err error){
	nodeList ,err := routes.GetAll()
	if err != nil {
		logrus.Warningf("CoreGet Failed[Err:%s]",err.Error())
		return  err
	}
	for _, node := range nodeList{
		if err = stream.Send(&pb.Data{Key:node.Key,Value:node.Value}); err != nil {
			logrus.Warnf("stream.Send Failed[Data:%+v, Err:%s]",node, err.Error())
			continue
		}
	}
	return
}
func (s *ZcacheServer) Export(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	err = routes.Flush()
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{}
	}
	return
}
func (s *ZcacheServer) Import(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	err = routes.Import()
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{}
	}
	return
}
func (s *ZcacheServer) DeleteKeys(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	err = routes.DeleteAll()
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{}
	}
	return
}
func (s *ZcacheServer) Expension(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	err = routes.Expension(data.Value)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{}
	}
	return
}
func (s *ZcacheServer) GetKeyNum(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	num, err := routes.GetKeyNum()
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{Key:"",Value:string(num)}
	}
	return
}
func (s *ZcacheServer) Incr(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	err = routes.Incr(data.Key)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{}
	}
	return
}
func (s *ZcacheServer) IncrBy(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	err = routes.IncrBy(data.Key,data.Value)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{}
	}
	return
}
func (s *ZcacheServer) Decr(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	err = routes.Decr(data.Key)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{}
	}
	return
}
func (s *ZcacheServer) DecrBy(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	err = routes.DecrBy(data.Key,data.Value)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{}
	}
	return
}
func (s *ZcacheServer) ImportFromRedis(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	err = routes.ImportFromRedis()
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{}
	}
	return
}

func (s *ZcacheServer) ExportToRedis(ctx context.Context, data *pb.Data)(resp *pb.Data, err error){
	err = routes.ExportToRedis()
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.Data{}
	}
	return
}

func (s *ZcacheServer) PreDeleteKey(ctx context.Context, data *pb.Data)(resp *pb.CommitIDMsg, err error){
	commitID, err := tool.GetHashIndex("Delete" + data.Key)
	preReq := types.ProcessingRequest{
		CommitID: commitID,
		Req:      types.ReqType_DELETE,
		Key:      data.Key,
		Value:    "",
		Next:     nil,
	}
	err = tool.AddInternalReq(preReq)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.CommitIDMsg{CommitID:string(commitID)}
	}
	return
}

func (s *ZcacheServer) PreSetValue(ctx context.Context, data *pb.Data)(resp *pb.CommitIDMsg, err error){
	commitID, err := tool.GetHashIndex("Set" + data.Key + data.Value)
	preReq := types.ProcessingRequest{
		CommitID: commitID,
		Req:      types.ReqType_PUT,
		Key:      data.Key,
		Value:    data.Value,
		Next:     nil,
	}
	err = tool.AddInternalReq(preReq)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.CommitIDMsg{CommitID:string(commitID)}
	}
	return
}
func (s *ZcacheServer) PreUpdateValue(ctx context.Context, data *pb.Data)(resp *pb.CommitIDMsg, err error){
	commitID, err := tool.GetHashIndex("Update" + data.Key + data.Value)
	preReq := types.ProcessingRequest{
		CommitID: commitID,
		Req:      types.ReqType_POST,
		Key:      data.Key,
		Value:    data.Value,
		Next:     nil,
	}
	err = tool.AddInternalReq(preReq)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.CommitIDMsg{CommitID:string(commitID)}
	}
	return
}
func (s *ZcacheServer) PreDeleteKeys(ctx context.Context, data *pb.Data)(resp *pb.CommitIDMsg, err error){
	//err = routes.ExportToRedis()
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.CommitIDMsg{}
	}
	return
}
func (s *ZcacheServer) PreExpension(ctx context.Context, data *pb.Data)(resp *pb.CommitIDMsg, err error){
	//err = routes.ExportToRedis()
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.CommitIDMsg{}
	}
	return
}
func (s *ZcacheServer) PreIncr(ctx context.Context, data *pb.Data)(resp *pb.CommitIDMsg, err error){
	commitID, err := tool.GetHashIndex("Incr" + data.Key)
	preReq := types.ProcessingRequest{
		CommitID: commitID,
		Req:      types.ReqType_INCR,
		Key:      data.Key,
		Value:    "",
		Next:     nil,
	}
	err = tool.AddInternalReq(preReq)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.CommitIDMsg{CommitID:string(commitID)}
	}
	return
}
func (s *ZcacheServer) PreIncrBy(ctx context.Context, data *pb.Data)(resp *pb.CommitIDMsg, err error){

	commitID, err := tool.GetHashIndex("IncrBy" + data.Key)
	preReq := types.ProcessingRequest{
		CommitID: commitID,
		Req:      types.ReqType_INCRBY,
		Key:      data.Key,
		Value:    data.Value,
		Next:     nil,
	}
	err = tool.AddInternalReq(preReq)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.CommitIDMsg{CommitID:string(commitID)}
	}
	return
}
func (s *ZcacheServer) PreDecr(ctx context.Context, data *pb.Data)(resp *pb.CommitIDMsg, err error){
	commitID, err := tool.GetHashIndex("Decr" + data.Key)
	preReq := types.ProcessingRequest{
		CommitID: commitID,
		Req:      types.ReqType_DECR,
		Key:      data.Key,
		Value:    "",
		Next:     nil,
	}
	err = tool.AddInternalReq(preReq)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.CommitIDMsg{CommitID:string(commitID)}
	}
	return
}
func (s *ZcacheServer) PreDecrBy(ctx context.Context, data *pb.Data)(resp *pb.CommitIDMsg, err error){
	commitID, err := tool.GetHashIndex("DecrBy" + data.Key)
	preReq := types.ProcessingRequest{
		CommitID: commitID,
		Req:      types.ReqType_DECRBY,
		Key:      data.Key,
		Value:    "",
		Next:     nil,
	}
	err = tool.AddInternalReq(preReq)
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.CommitIDMsg{CommitID:string(commitID)}
	}
	return
}
func (s *ZcacheServer) PreImportFromRedis(ctx context.Context, data *pb.Data)(resp *pb.CommitIDMsg, err error){
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.CommitIDMsg{}
	}
	return
}
func (s *ZcacheServer) Commit(ctx context.Context, data *pb.CommitIDMsg)(resp *pb.CommitIDMsg, err error){
	commitID := data.CommitID
	global.GlobalVar.GInternalLock.Lock()
	defer global.GlobalVar.GInternalLock.Unlock()

	var toDORequest  *types.ProcessingRequest = nil
	if(string(global.GlobalVar.GPreDoReqList.CommitID) == commitID){
		toDORequest = global.GlobalVar.GPreDoReqList
		global.GlobalVar.GPreDoReqList = global.GlobalVar.GPreDoReqList.Next
	} else{
		var tmpPre *types.ProcessingRequest = global.GlobalVar.GPreDoReqList
		for tmpPre.Next != nil && string(tmpPre.Next.CommitID) != commitID{
			tmpPre = tmpPre.Next
		}
		if tmpPre.Next != nil && string(tmpPre.Next.CommitID) == commitID{
			toDORequest = tmpPre.Next
			tmpPre.Next = tmpPre.Next.Next
		}

	}


	if toDORequest == nil {
		err = errors.New("no such job")
	} else {
		err = zdata.DoCommit(toDORequest)
	}
	if err != nil{
		return nil ,err
	} else {
		resp = &pb.CommitIDMsg{}
	}
	return
}
func (s *ZcacheServer) Drop(ctx context.Context, data *pb.CommitIDMsg)(resp *pb.CommitIDMsg, err error){
	commitID := data.CommitID
	global.GlobalVar.GInternalLock.Lock()
	defer global.GlobalVar.GInternalLock.Unlock()

	if(string(global.GlobalVar.GPreDoReqList.CommitID) == commitID){
		global.GlobalVar.GPreDoReqList = global.GlobalVar.GPreDoReqList.Next
	} else {
		var tmpPre *types.ProcessingRequest = global.GlobalVar.GPreDoReqList
		for tmpPre.Next != nil && string(tmpPre.Next.CommitID) != commitID {
			tmpPre = tmpPre.Next
		}
		if tmpPre.Next != nil && string(tmpPre.Next.CommitID) == commitID {
			tmpPre.Next = tmpPre.Next.Next
		}
	}
		if err != nil{
		return nil ,err
	} else {
		resp = &pb.CommitIDMsg{}
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
