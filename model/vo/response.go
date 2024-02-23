package vo

import (
	"errors"

	"github.com/ncuhome/story-cook/pkg/myErrors"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Error(err error, status ...int) *Response {
	if err == nil {
		err = errors.New("internal error")
	}

	code := myErrors.ERROR
	if len(status) > 0 {
		code = status[0]
	}

	return &Response{
		Code: code,
		Msg:  err.Error(),
		Data: nil,
	}
}

func Success() *Response {
	return &Response{
		Code: myErrors.SUCCESS,
		Msg:  "操作成功",
	}
}

func SuccessWithData(data any) *Response {
	return &Response{
		Code: myErrors.SUCCESS,
		Msg:  "操作成功",
		Data: data,
	}
}

// DataList 带有总数的Data结构
type DataList struct {
	Item  any   `json:"item"`
	Total int64 `json:"total"`
}

// List 带有总数的列表构建器
func List(items any, total int64) *Response {
	return &Response{
		Code: myErrors.SUCCESS,
		Data: DataList{
			Item:  items,
			Total: total,
		},
		Msg: "查询列表成功",
	}
}
