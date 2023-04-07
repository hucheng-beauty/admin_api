package response

import (
	"reflect"

	"admin_api/internal/enum"
)

type EmptyMessage struct {
	Message string `json:"msg"`
}

func NewEmptyMessage() *EmptyMessage {
	return &EmptyMessage{}
}

type Response struct {
	RetCode int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func NewResponse(retCode int, message string, data interface{}) *Response {
	return &Response{RetCode: retCode, Message: message, Data: data}
}

type Pagination struct {
	TotalCount int `json:"total_count"`
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
}

func NewPagination(total, offset, limit int) *Pagination {
	return &Pagination{
		TotalCount: total,
		Offset:     offset,
		Limit:      limit,
	}
}

type WithPagination struct {
	*Response
	*Pagination
}

func NewResponseWithPagination(response *Response, pagination *Pagination) *WithPagination {
	return &WithPagination{Response: response, Pagination: pagination}
}

func SuccessWithPagination(data interface{}, pagination *Pagination) *WithPagination {
	if v := reflect.ValueOf(data); v.IsNil() {
		data = []interface{}{}
	}
	return NewResponseWithPagination(NewResponse(enum.CodeMapRequest["Success"], NewEmptyMessage().Message, data), pagination)
}

func Error(code int, message string) *Response {
	return NewResponse(code, message, nil)
}
