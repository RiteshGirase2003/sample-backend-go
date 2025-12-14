package response


type moduleType int

const (
	PING moduleType = iota 
	
)

var ModuleEnum = map[moduleType]string{ 
	PING:  "ping", 
	
}

type ErrorType int

const (
	OK ErrorType = 200

	// 4xx Client Errors
	BadRequest          ErrorType = 400
	Unauthorized        ErrorType = 401
	Forbidden           ErrorType = 403
	NotFound            ErrorType = 404
	MethodNotAllowed    ErrorType = 405
	Conflict            ErrorType = 409
	UnprocessableEntity ErrorType = 422
	TooManyRequests     ErrorType = 429

	// 5xx Server Errors
	InternalServerError ErrorType = 500
	NotImplemented      ErrorType = 501
	BadGateway           ErrorType = 502
	ServiceUnavailable  ErrorType = 503
	GatewayTimeout      ErrorType = 504
)

var ErrorMessage = map[ErrorType]string{
	OK: "ok",

	BadRequest:          "bad request",
	Unauthorized:        "unauthorized",
	Forbidden:           "forbidden",
	NotFound:            "not found",
	MethodNotAllowed:    "method not allowed",
	Conflict:            "conflict",
	UnprocessableEntity: "unprocessable entity",
	TooManyRequests:     "too many requests",

	InternalServerError: "internal server error",
	NotImplemented:      "not implemented",
	BadGateway:          "bad gateway",
	ServiceUnavailable: "service unavailable",
	GatewayTimeout:     "gateway timeout",
}


type Response struct {
	Service string      `json:"service"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}


func (r *Response) ResponseWithData(
	service moduleType,
	status ErrorType,
	data interface{},
) {
	r.Service = ModuleEnum[service]
	r.Status = int(status)
	r.Message = ErrorMessage[status]
	r.Data = data
}

func (r *Response) ResponseWithError(
	service moduleType,
	status ErrorType,
	errorMsg string,
) {
	msg := errorMsg
	if msg == "" {
		msg = ErrorMessage[status]
	}

	r.Service = ModuleEnum[service]
	r.Status = int(status)
	r.Message = msg
	r.Data = nil
}


