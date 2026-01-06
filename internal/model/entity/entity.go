/*
 * Copyright (c) 2026 mincancao. All rights reserved.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

/*
* @Description:
* @Author: mincancao
* @Date: 2026/1/6 21:59
* @LastEditTime: 2026/1/6 21:59
* @LastEditors: mincancao
 */

package entity

type TaskFcn func()

type Task interface {
	Start()
	Stop()
	GetName() string
}
