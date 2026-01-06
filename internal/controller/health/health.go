/*
 * Copyright (c) 2026 mincancao. All rights reserved.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

/*
* @Description:
* @Author: mincancao
* @Date: 2026/1/6 21:29
* @LastEditTime: 2026/1/6 21:29
* @LastEditors: mincancao
 */

package health

import "github.com/gogf/gf/v2/net/ghttp"

func Health(r *ghttp.Request) {
	r.SetCtxVar("Code", "200")
	r.Response.Writeln("ok")
}
