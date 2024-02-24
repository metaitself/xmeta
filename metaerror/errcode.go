package metaerror

// GRPC Status
var (
	ErrOK                 = New(0, "成功")
	ErrCanceled           = New(1, "操作已取消")
	ErrUnknown            = New(2, "未知错误")
	ErrInvalidArgument    = New(3, "无效参数")
	ErrDeadlineExceeded   = New(4, "超过最后期限")
	ErrNotFound           = New(5, "数据不存在")
	ErrAlreadyExists      = New(6, "已经存在")
	ErrPermissionDenied   = New(7, "权限不足")
	ErrResourceExhausted  = New(8, "资源耗尽")
	ErrFailedPrecondition = New(9, "操作被拒绝")
	ErrAborted            = New(10, "操作被中止")
	ErrOutOfRange         = New(11, "操作超过有效范围")
	ErrUnimplemented      = New(12, "此操作未被支持")
	ErrInternalException  = New(13, "发生错误，请稍后重试")
	ErrUnavailable        = New(14, "当前服务不可用")
	ErrDataLoss           = New(15, "数据丢失")
	ErrUnauthenticated    = New(16, "表示请求没有有效的操作认证凭证")
)

// Base error
var (
	ErrMarshal   = New(101, "打包数据失败")
	ErrUnmarshal = New(102, "解析数据出错")
)
