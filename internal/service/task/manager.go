/*
 * Copyright (c) 2026 mincancao. All rights reserved.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

/*
* @Description:
* @Author: mincancao
* @Date: 2026/1/6 22:33
* @LastEditTime: 2026/1/6 22:33
* @LastEditors: mincancao
 */

package task

import (
	"sync"

	"github.com/caomincan/web-framework/internal/logic/task"
	"github.com/caomincan/web-framework/internal/model/entity"
)

var (
	defaultManager BackgroundTaskManager
	onceTask       sync.Once
)

type BackgroundTaskManager interface {
	AddTask(task entity.Task)
	StartAll()
	StopAll()
}

func init() {
	onceTask.Do(func() {
		defaultManager = task.NewTaskManager()
	})
}

func GetBackgroundTaskManager() BackgroundTaskManager {
	return defaultManager
}
