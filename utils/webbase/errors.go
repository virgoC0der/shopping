package webbase

type CommonResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

var (
	ErrOK = &CommonResp{
		Code:    0,
		Message: "成功",
	}
	ErrSystemBusy = &CommonResp{
		Code:    10001,
		Message: "系统繁忙，请稍后再试",
	}
	ErrInputParams = &CommonResp{
		Code:    10002,
		Message: "客户端输入错误",
	}
	ErrAuthFailed = &CommonResp{
		Code:    10003,
		Message: "认证失败",
	}
	ErrNotLogin = &CommonResp{
		Code:    10004,
		Message: "未登录",
	}
	ErrNoPermission = &CommonResp{
		Code:    10005,
		Message: "无权限访问",
	}
)
