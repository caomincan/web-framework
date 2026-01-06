/*
 * Copyright (c) 2026 mincancao. All rights reserved.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

/*
* @Description:
* @Author: mincancao
* @Date: 2026/1/6 21:31
* @LastEditTime: 2026/1/6 21:31
* @LastEditors: mincancao
 */

package version

import (
	"github.com/caomincan/web-framework/api/common"
	"github.com/caomincan/web-framework/internal/const/web"
	"github.com/gogf/gf/v2/net/ghttp"
)

// set var through ldflags
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

type buildInfo struct {
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	BuildDate string `json:"buildDate"`
}

type BuildInfoRes struct {
	common.ResponseBase
	Data *buildInfo `json:"data"`
}

func BuildInfo(r *ghttp.Request) {
	r.SetCtxVar("Code", "200")
	requestID := r.GetCtxVar(web.ContextKeyRequestID).String()
	res := BuildInfoRes{
		ResponseBase: common.ResponseBase{
			RequestID: requestID,
			Code:      0,
			Message:   "success",
			Success:   true,
		},
		Data: &buildInfo{
			Version:   version,
			Commit:    commit,
			BuildDate: date,
		},
	}
	r.Response.WriteJson(res)
}
