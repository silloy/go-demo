package src

type ErrCode int64 //错误码
const (
	ERR_CODE_OK             ErrCode = 0 // PROCESS OK
	ERR_CODE_INVALID_PARAMS ErrCode = 1 // 参数无效
	ERR_CODE_TIMEOUT        ErrCode = 2 // 超时
	ERR_CODE_FILE_NOT_EXIST ErrCode = 3 // 文件不存在
	ERR_CODE_CONN_REFUSE    ErrCode = 4 // 连接被拒绝
	ERR_CODE_NET_ABNORMAL   ErrCode = 5 // 网络异常
)
