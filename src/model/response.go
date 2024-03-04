package model

import (
	"time"

	"github.com/achwanyusuf/carrent-lib/pkg/errormsg"
	"github.com/achwanyusuf/carrent-lib/pkg/logger"
	"github.com/gin-gonic/gin"
)

type TransactionInfo struct {
	RequestURI    string    `json:"request_uri"`
	RequestMethod string    `json:"request_method"`
	RequestID     string    `json:"request_id"`
	Timestamp     time.Time `json:"timestamp"`
	ErrorCode     int64     `json:"error_code,omitempty"`
	Cause         string    `json:"cause,omitempty"`
}

type Response struct {
	TransactionInfo TransactionInfo `json:"transaction_info"`
	Code            int64           `json:"status_code"`
	Message         string          `json:"message,omitempty"`
	Translation     *Translation    `json:"translation,omitempty"`
}

type EmptyResponse struct {
	Response
}

func (r *EmptyResponse) Transform(ctx *gin.Context, log logger.Logger, code int, err error) int {
	r.Response = Response{
		TransactionInfo: TransactionInfo{
			RequestURI:    ctx.Request.RequestURI,
			RequestMethod: ctx.Request.Method,
			RequestID:     ctx.GetHeader("x-request-id"),
			Timestamp:     time.Now(),
		},
		Code: int64(code),
	}
	if err != nil {
		getErrMsg := errormsg.GetErrorData(err)
		r.Response.TransactionInfo.ErrorCode = getErrMsg.Code
		log.Error(ctx, errormsg.WriteErr(err))
		r.Response.Code = getErrMsg.WrappedMessage.StatusCode
		r.Response.Message = getErrMsg.WrappedMessage.Message
		translation := Translation(getErrMsg.WrappedMessage.Translation)
		r.Response.Translation = &translation
	}

	return int(r.Response.Code)
}

type Translation struct {
	EN string `json:"en"`
}

type SingleCarResponse struct {
	Response
	Data Car `json:"data"`
}

func (r *SingleCarResponse) Transform(ctx *gin.Context, log logger.Logger, code int, err error) int {
	r.Response = Response{
		TransactionInfo: TransactionInfo{
			RequestURI:    ctx.Request.RequestURI,
			RequestMethod: ctx.Request.Method,
			RequestID:     ctx.GetHeader("x-request-id"),
			Timestamp:     time.Now(),
		},
		Code: int64(code),
	}
	if err != nil {
		getErrMsg := errormsg.GetErrorData(err)
		r.Response.TransactionInfo.ErrorCode = getErrMsg.Code
		log.Error(ctx, errormsg.WriteErr(err))
		r.Response.Code = getErrMsg.WrappedMessage.StatusCode
		r.Response.Message = getErrMsg.WrappedMessage.Message
		translation := Translation(getErrMsg.WrappedMessage.Translation)
		r.Response.Translation = &translation
	}

	return int(r.Response.Code)
}

type CarsResponse struct {
	Response
	Data       []Car      `json:"data"`
	Pagination Pagination `json:"pagination"`
}

func (r *CarsResponse) Transform(ctx *gin.Context, log logger.Logger, code int, err error) int {
	r.Response = Response{
		TransactionInfo: TransactionInfo{
			RequestURI:    ctx.Request.RequestURI,
			RequestMethod: ctx.Request.Method,
			RequestID:     ctx.GetHeader("x-request-id"),
			Timestamp:     time.Now(),
		},
		Code: int64(code),
	}
	if err != nil {
		getErrMsg := errormsg.GetErrorData(err)
		r.Response.TransactionInfo.ErrorCode = getErrMsg.Code
		log.Error(ctx, errormsg.WriteErr(err))
		r.Response.Code = getErrMsg.WrappedMessage.StatusCode
		r.Response.Message = getErrMsg.WrappedMessage.Message
		translation := Translation(getErrMsg.WrappedMessage.Translation)
		r.Response.Translation = &translation
	}

	if len(r.Data) == 0 {
		r.Data = []Car{}
	}

	return int(r.Response.Code)
}

type SingleOrderResponse struct {
	Response
	Data Order `json:"data"`
}

func (r *SingleOrderResponse) Transform(ctx *gin.Context, log logger.Logger, code int, err error) int {
	r.Response = Response{
		TransactionInfo: TransactionInfo{
			RequestURI:    ctx.Request.RequestURI,
			RequestMethod: ctx.Request.Method,
			RequestID:     ctx.GetHeader("x-request-id"),
			Timestamp:     time.Now(),
		},
		Code: int64(code),
	}
	if err != nil {
		getErrMsg := errormsg.GetErrorData(err)
		r.Response.TransactionInfo.ErrorCode = getErrMsg.Code
		log.Error(ctx, errormsg.WriteErr(err))
		r.Response.Code = getErrMsg.WrappedMessage.StatusCode
		r.Response.Message = getErrMsg.WrappedMessage.Message
		translation := Translation(getErrMsg.WrappedMessage.Translation)
		r.Response.Translation = &translation
	}

	return int(r.Response.Code)
}

type OrdersResponse struct {
	Response
	Data       []Order    `json:"data"`
	Pagination Pagination `json:"pagination"`
}

func (r *OrdersResponse) Transform(ctx *gin.Context, log logger.Logger, code int, err error) int {
	r.Response = Response{
		TransactionInfo: TransactionInfo{
			RequestURI:    ctx.Request.RequestURI,
			RequestMethod: ctx.Request.Method,
			RequestID:     ctx.GetHeader("x-request-id"),
			Timestamp:     time.Now(),
		},
		Code: int64(code),
	}
	if err != nil {
		getErrMsg := errormsg.GetErrorData(err)
		r.Response.TransactionInfo.ErrorCode = getErrMsg.Code
		log.Error(ctx, errormsg.WriteErr(err))
		r.Response.Code = getErrMsg.WrappedMessage.StatusCode
		r.Response.Message = getErrMsg.WrappedMessage.Message
		translation := Translation(getErrMsg.WrappedMessage.Translation)
		r.Response.Translation = &translation
	}

	if len(r.Data) == 0 {
		r.Data = []Order{}
	}

	return int(r.Response.Code)
}
