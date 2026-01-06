/*
 * Copyright (c) 2026 mincancao. All rights reserved.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

/*
* @Description:
* @Author: mincancao
* @Date: 2026/1/6 21:11
* @LastEditTime: 2026/1/6 21:11
* @LastEditors: mincancao
 */

package main

import (
	"github.com/caomincan/web-framwork/internal/cmd"
	"github.com/caomincan/web-framwork/internal/logic/prepare"
	"github.com/caomincan/web-framwork/internal/service/task"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {
	prepare.InitSystem()
	cmd.RunHttpServer()
	task.GetBackgroundTaskManager().StartAll()
	defer task.GetBackgroundTaskManager().StopAll()
	// block and listening
	ghttp.Wait()

}
