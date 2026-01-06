/*
 * Copyright (c) 2026 mincancao. All rights reserved.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

/*
* @Description:
* @Author: mincancao
* @Date: 2026/1/6 22:17
* @LastEditTime: 2026/1/6 22:17
* @LastEditors: mincancao
 */

package task

import (
	"fmt"

	"github.com/caomincan/web-framwork/internal/model/entity"
)

type Manager struct {
	tasks []entity.Task
}

func (m *Manager) AddTask(task entity.Task) {
	if m.tasks != nil {
		m.tasks = append(m.tasks, task)
	}
}

func (m *Manager) StartAll() {
	fmt.Println("start all tasks in background...")
	fmt.Println("total size of task:", len(m.tasks))
	for _, task := range m.tasks {
		fmt.Println("= start task:", task.GetName())
		task.Start()
	}
}

func (m *Manager) StopAll() {
	fmt.Println("stop all tasks in background...")
	fmt.Println("total size of task:", len(m.tasks))
	for _, task := range m.tasks {
		fmt.Println("= stop task:", task.GetName())
		task.Stop()
	}
}

func NewTaskManager() *Manager {
	return &Manager{
		tasks: make([]entity.Task, 0),
	}
}
