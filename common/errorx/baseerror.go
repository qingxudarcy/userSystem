package errorx

import "net/http"

const defaultCode = http.StatusInternalServerError

type errorField struct {
	Code string `json:"code"`
	Msg  string `json:"messgae"`
}

type CheckError struct {
	Msg    string                `json:"msg"`
	Fields map[string]errorField `json:"fields"`
}

type CheckErrorResponse struct {
	Msg    string                `json:"msg"`
	Fields map[string]errorField `json:"fields"`
}

func NewCheckError(fieldName, message string) *CheckError {
	fields := make(map[string]errorField)
	fields[fieldName] = errorField{Code: "invalid", Msg: message}
	return &CheckError{Msg: "字段校验失败", Fields: fields}
}

func (err *CheckError) Error() string {
	return err.Msg
}

func (err *CheckError) Data() *CheckErrorResponse {
	return &CheckErrorResponse{Msg: err.Msg, Fields: err.Fields}
}

func ErrorResponse(err error) (int, interface{}) {
	switch e := err.(type) {
	case *CheckError:
		return http.StatusUnprocessableEntity, e.Data()
	default:
		return http.StatusInternalServerError, nil
	}
}
