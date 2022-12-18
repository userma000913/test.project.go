package model

type (
	// 接口响应结果
	ApiResult struct {
		// 错误码
		DmError int `json:"dm_error"`
		// 错误提示语
		ErrorMsg string `json:"error_msg"`
	}

	// 响应基类
	RespBase struct {
		ApiResult
		Data interface{} `json:"data"`
	}
)
