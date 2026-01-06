/*
 * Copyright (c) 2026 mincancao. All rights reserved.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

/*
* @Description:
* @Author: mincancao
* @Date: 2026/1/6 21:33
* @LastEditTime: 2026/1/6 21:33
* @LastEditors: mincancao
 */

package web

const (
	ContextKeyRequestID   = "RequestID"
	ContextKeyCode        = "Code"
	ContextKeyAction      = "Action"
	ContextKeyContextID   = "ContextID"
	ContextKeyCallerID    = "CallerID"
	ContextKeyCallerType  = "CallerType"
	ContextKeyForwardedIP = "ForwardedIP"

	HeaderKeyContentType     = "Content-Type"
	HeaderKeyContentLanguage = "Content-Language"
	HeaderKeyContextID       = "CONTEXT-ID"
	HeaderKeyRequestID       = "X-RPC-RequestID"
	HeaderKeyCallerType      = "X-RPC-CallerType"
	HeaderKeyCallerID        = "X-RPC-CallerID"
	HeaderKeyForwardedIP     = "X-Forwarded-For"

	HeaderValueAPPJson         = "application/json"
	HeaderValueLanguageEnglish = "en"

	EnvKeyEnvType    = "ENV_TYPE"
	EnvKeyServerMode = "SERVER_MODE"
)
