package serializer

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`//错误码
	Data   interface{} `json:"data"`//值
	Msg    string      `json:"msg"`//提示
	Error  string      `json:"error"`//错误
}

