/*
 * Copyright (c) 2026 mincancao. All rights reserved.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

/*
* @Description:
* @Author: mincancao
* @Date: 2026/1/6 21:20
* @LastEditTime: 2026/1/6 21:20
* @LastEditors: mincancao
 */

package cmd

import (
	"context"

	"github.com/caomincan/web-framework/internal/server"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	Main = &gcmd.Command{
		Name:        "main",
		Brief:       "start http server",
		Description: "this is the command entry for starting your server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			return server.GetDefaultHttpServer().Start()
		},
		Strict: false,
	}
)

func RunHttpServer() {
	Main.Run(gctx.GetInitCtx())
}
