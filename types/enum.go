package types

type ClusterHealthType int8

const (
	CLUSTER_HEALTH_TYPE_UNKNOWN ClusterHealthType = iota
	CLUSTER_HEALTH_TYPE_UNHEALTH  //不健康
	CLUSTER_HEALTH_TYPE_SUBHEALTH  //亚健康
	CLUSTER_HEALTH_TYPE_HEALTH  //健康

)

type OperationType int8

const (
	OPERATION_TYPE_UNKNOW OperationType = iota
	OPERATION_TYPE_SET  //不健康
	OPERATION_TYPE_GET  //亚健康
	OPERATION_TYPE_POST  //不健康
	OPERATION_TYPE_DELETE  //不健康

)


type LogType int8

const (
	LogType_UNKNOW LogType = iota
	LogType_DEBUG
	LogType_INFO
	LogType_WARNING
	LogType_ERROR
)

const (
	HttpGet = "GET"
	HttpDelete = "DELETE"
	HttpPOST = "POST"
	HttpPut = "PUT"
)

type ReqType int8
const (
	ReqType_GET  ReqType = iota
	ReqType_DELETE
	ReqType_PUT
	ReqType_POST
	ReqType_INCR
	ReqType_INCRBY
	ReqType_DECR
	ReqType_DECRBY
)