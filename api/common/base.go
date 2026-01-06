/*
 * Copyright (c) 2026 mincancao. All rights reserved.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

/*
* @Description:
* @Author: mincancao
* @Date: 2026/1/6 21:12
* @LastEditTime: 2026/1/6 21:12
* @LastEditors: mincancao
 */

package common

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type ResponseBase struct {
	RequestID string `json:"requestID"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Success   bool   `json:"success"`
}

func BuildInvalidParameterRes(requestID, cause string) *DataRes {
	return &DataRes{
		ResponseBase: ResponseBase{
			RequestID: requestID,
			Code:      400,
			Success:   false,
			Message:   cause,
		},
		Data: nil,
	}
}

func BuildMissingParameterRes(requestID, cause string) *DataRes {
	return &DataRes{
		ResponseBase: ResponseBase{
			RequestID: requestID,
			Code:      400,
			Success:   false,
			Message:   cause,
		},
		Data: nil,
	}
}

type DataRes struct {
	ResponseBase
	Data interface{} `json:"data"`
}

func BuildErrorDataRes(requestID string, err error) *DataRes {
	if err == nil {
		panic("should not build error response when error is nil")
	}
	returnCode := 500
	code := gerror.Code(err)
	if code != gcode.CodeNil {
		returnCode = code.Code()
	}
	return &DataRes{
		ResponseBase: buildErrorBaseRes(requestID, err.Error(), returnCode),
		Data:         nil,
	}
}

func BuildDataRes(requestID string, data interface{}) *DataRes {
	return &DataRes{
		ResponseBase: buildSuccessRes(requestID),
		Data:         data,
	}
}

func buildSuccessRes(requestID string) ResponseBase {
	return ResponseBase{
		RequestID: requestID,
		Code:      200,
		Success:   true,
		Message:   "success",
	}
}

func buildErrorBaseRes(requestID, message string, code int) ResponseBase {
	return ResponseBase{
		RequestID: requestID,
		Code:      code,
		Message:   message,
		Success:   false,
	}
}
