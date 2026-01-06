/*
 * Copyright (c) 2026 mincancao. All rights reserved.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

/*
* @Description:
* @Author: mincancao
* @Date: 2026/1/6 21:22
* @LastEditTime: 2026/1/6 21:22
* @LastEditors: mincancao
 */

package server

import (
	"sync"

	"github.com/caomincan/web-framework/internal/controller/health"
	"github.com/caomincan/web-framework/internal/controller/version"
	"github.com/caomincan/web-framework/internal/middleware"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	defaultServer *ghttp.Server
	once          sync.Once
)

func GetDefaultHttpServer() *ghttp.Server {
	once.Do(func() {
		defaultServer = NewHttpServer()
	})
	return defaultServer
}

func NewHttpServer() *ghttp.Server {
	s := g.Server()
	// not print router map
	s.SetDumpRouterMap(false)
	// set global handler
	s.Use(middleware.RequestHandler)
	// register api
	s.Group("/api/v1", func(group *ghttp.RouterGroup) {
		group.Middleware()
		group.Bind()
	})
	// register health handler
	if !hasRegisterPath(s, "/health") {
		s.BindHandler("/health", health.Health)
	}
	// register build info handler
	if !hasRegisterPath(s, "/buildinfo") {
		s.BindHandler("/buildinfo", version.BuildInfo)
	}
	// bind global not found handler
	s.BindStatusHandler(404, middleware.NotFoundHandler)
	return s
}

func hasRegisterPath(s *ghttp.Server, pattern string) bool {
	for _, r := range s.GetRoutes() {
		if r.Route == pattern {
			return true
		}
	}
	return false
}
