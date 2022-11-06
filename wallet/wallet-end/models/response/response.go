package response

const SUCCESS = 200
const BAD_REQUEST = 400
const COMMON_FAIL = 400

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func NewResponseWithString(val string) Response {

	res := Response{
		Code:    SUCCESS,
		Message: "",
		Result:  val,
	}
	return res
}

func NewResponseWithPair(key string, value string) Response {
	m := make(map[string]string)
	m[key] = value
	res := Response{
		Code:    SUCCESS,
		Message: "",
		Result:  m,
	}
	return res
}
func NewResponseWithSuccess() Response {

	res := Response{
		Code:    SUCCESS,
		Message: "",
		Result:  nil,
	}
	return res
}

func NewResponseWithData(value interface{}) Response {

	res := Response{
		Code:    SUCCESS,
		Message: "",
		Result:  value,
	}
	return res
}

func ErrorResponse(code int, error error) Response {
	res := Response{
		Code:    code,
		Message: error.Error(),
		Result:  nil,
	}
	return res
}
