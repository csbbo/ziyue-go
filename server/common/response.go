package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

type Response struct {
	Err  interface{} `json:"err"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func RemoveFrontStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	resp := &Response{Err: nil, Msg: nil, Data: data}
	c.JSON(200, resp)
}

func HandleErr(c *gin.Context, errType string, msg interface{}) {
	errs, ok := msg.(validator.ValidationErrors)
	var resp *Response
	if !ok {
		fmt.Println(reflect.TypeOf(msg))
		resp = &Response{Err: errType, Msg: msg, Data: nil}
	} else {
		trans, _ := c.Get("trans")
		translator, _ := trans.(ut.Translator)
		transMsg := RemoveFrontStruct(errs.Translate(translator))
		resp = &Response{Err: errType, Msg: transMsg, Data: nil}
	}
	c.JSON(200, resp)
}

func ResponseError(c *gin.Context, msg interface{}) {
	HandleErr(c, "err", msg)
}

func ServerError(c *gin.Context, msg interface{}) {
	if msg == nil {
		HandleErr(c, "server-err", "服务器错误")
	} else {
		HandleErr(c, "server-err", msg)
	}
}