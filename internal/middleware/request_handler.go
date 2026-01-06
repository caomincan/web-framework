/*
 * Copyright (c) 2026 mincancao. All rights reserved.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

/*
* @Description:
* @Author: mincancao
* @Date: 2026/1/6 21:25
* @LastEditTime: 2026/1/6 21:25
* @LastEditors: mincancao
 */

package middleware

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/caomincan/web-framwork/api/common"
	"github.com/caomincan/web-framwork/internal/const/web"
	"github.com/caomincan/web-framwork/internal/logic/utils"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gvalid"
	"github.com/google/uuid"
)

var (
	errLogger = g.Log("errLogger")
	logger    = g.Log("appLogger")
)

func RequestHandler(r *ghttp.Request) {
	buildContext(r)
	// log begin
	logBeginRequest(r)
	start := time.Now()
	// request process
	r.Middleware.Next()
	// handler res
	if err := r.GetError(); err != nil {
		errLogger.Error(r.GetCtx(), g.Map{
			"Message":      "request handler error",
			"ErrorMessage": utils.SafeGetErrorMessage(err),
		})
		processError(r, err)
	} else if response := r.GetHandlerResponse(); response != nil {
		processNormalResponse(r, response)
	}
	elapsed := time.Since(start)
	// log end
	logEndRequest(r, elapsed)
	// finished
	r.Exit()
}

func buildContext(r *ghttp.Request) {
	ctxID := r.GetHeader(web.HeaderKeyContextID)
	if ctxID == "" {
		ctxID = uuid.NewString()
	}
	r.SetCtxVar(web.ContextKeyContextID, ctxID)
	reqID := r.GetHeader(web.HeaderKeyRequestID)
	if reqID == "" {
		reqID = uuid.New().String()
	}
	r.SetCtxVar(web.ContextKeyRequestID, reqID)
	method := strings.ToUpper(r.Method)
	path := r.URL.Path
	action := method + "#" + path
	r.SetCtxVar(web.ContextKeyAction, action)

	callerType := r.GetHeader(web.HeaderKeyCallerType)
	if callerType == "" {
		callerType = "unknown"
	}
	r.SetCtxVar(web.ContextKeyCallerType, callerType)
	callerID := r.GetHeader(web.HeaderKeyCallerID)
	if callerID == "" {
		callerID = "unknown"
	}
	r.SetCtxVar(web.ContextKeyCallerID, callerID)
	r.SetCtxVar(web.ContextKeyForwardedIP, r.GetHeader(web.HeaderKeyForwardedIP))
}

func logBeginRequest(r *ghttp.Request) {
	uri := r.RequestURI
	host := r.Host
	xxf := r.GetCtxVar(web.ContextKeyForwardedIP).String()
	logger.Info(r.GetCtx(), g.Map{
		"Host":        host,
		"URI":         uri,
		"LogPosition": "BeginRequest",
		"Method":      r.Method,
		"ForwardedIP": xxf,
	})
}

func logEndRequest(r *ghttp.Request, elapsed time.Duration) {
	code := r.GetCtxVar(web.ContextKeyCode).String()
	uri := r.RequestURI
	host := r.Host
	xxf := r.GetCtxVar(web.ContextKeyForwardedIP).String()
	logger.Info(r.GetCtx(), g.Map{
		"Host":        host,
		"URI":         uri,
		"LogPosition": "EndRequest",
		"Method":      r.Method,
		"ForwardedIP": xxf,
		"Code":        code,
		"ElapsedTime": elapsed.Milliseconds(),
	})

}

func processNormalResponse(r *ghttp.Request, resp interface{}) {
	r.Response.ClearBuffer()
	r.Response.Header().Set(web.HeaderKeyContentType, web.HeaderValueAPPJson)
	commonResp, ok := resp.(*common.DataRes)
	if !ok {
		r.SetCtxVar(web.ContextKeyCode, "500")
	} else {
		if commonResp.Success {
			r.SetCtxVar(web.ContextKeyCode, "200")
		} else {
			r.SetCtxVar(web.ContextKeyCode, strconv.Itoa(commonResp.Code))
		}
	}
	r.Response.WriteJson(resp)
}

func processError(r *ghttp.Request, err error) {
	requestID := r.GetCtxVar(web.ContextKeyRequestID).String()
	r.Response.ClearBuffer()
	r.Response.Header().Set(web.HeaderKeyContentType, web.HeaderValueAPPJson)
	// check if error is validation error
	var v gvalid.Error
	if errors.As(err, &v) {
		var errKey, errMsg string
		_, firstItem := v.FirstItem()
		for k, e := range firstItem {
			errKey = k
			errMsg = e.Error()
			break
		}
		var resp *common.DataRes
		if errKey == "required" {
			resp = common.BuildMissingParameterRes(requestID, errMsg)
		} else {
			resp = common.BuildInvalidParameterRes(requestID, errMsg)
		}
		r.SetCtxVar(web.ContextKeyCode, strconv.Itoa(resp.Code))
		r.Response.WriteJson(resp)
	} else {
		buildInternalErrorResp(r, err)
	}
}

func buildInternalErrorResp(r *ghttp.Request, err error) {
	requestID := r.GetCtxVar(web.ContextKeyRequestID).String()
	message := "internal error"
	if err != nil {
		message = message + ": " + err.Error()
	}
	r.Response.WriteJson(&common.DataRes{
		ResponseBase: common.ResponseBase{
			RequestID: requestID,
			Message:   message,
			Success:   false,
			Code:      500,
		},
		Data: nil,
	})
	r.SetCtxVar(web.ContextKeyCode, "500")
}
