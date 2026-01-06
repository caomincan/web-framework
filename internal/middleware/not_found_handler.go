/*
 * Copyright (c) 2026 mincancao. All rights reserved.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

/*
* @Description:
* @Author: mincancao
* @Date: 2026/1/6 21:36
* @LastEditTime: 2026/1/6 21:36
* @LastEditors: mincancao
 */

package middleware

import (
	"github.com/caomincan/web-framework/internal/const/web"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func NotFoundHandler(r *ghttp.Request) {
	code := r.GetCtxVar(web.ContextKeyCode).String()
	requestID := r.GetCtxVar(web.ContextKeyRequestID).String()
	// if it is not biz logic to give 404 code then means uri not found
	if code != "404" {
		r.Response.ClearBuffer()
		r.Response.Header().Set(web.HeaderKeyContentType, web.HeaderValueAPPJson)
		r.Response.Header().Set(web.HeaderKeyContentLanguage, web.HeaderValueLanguageEnglish)
		r.Response.WriteJson(g.Map{
			"requestID": requestID,
			"code":      404,
			"message":   "invalid uri.",
			"success":   false,
			"data":      nil,
		})
		r.SetCtxVar(web.ContextKeyCode, "404")
		logger.Info(r.GetCtx(), g.Map{
			"Message": "invalid uri.",
			"Host":    r.Host,
			"URI":     r.RequestURI,
		})
	}
}
