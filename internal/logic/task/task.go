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
	"context"
	"sync/atomic"
	"time"

	"github.com/caomincan/web-framework/internal/const/web"
	"github.com/caomincan/web-framework/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/google/uuid"
)

var (
	logger = g.Log("appLogger")
)

type ScheduledTask struct {
	fcn          entity.TaskFcn
	name         string
	rotationTime time.Duration
	running      int64
	stop         chan bool
}

func (s *ScheduledTask) GetName() string {
	return s.name
}

func (s *ScheduledTask) Start() {
	old := int64(0)
	if atomic.CompareAndSwapInt64(&s.running, old, int64(1)) {
		go s.task()
	}
}

func (s *ScheduledTask) Stop() {
	old := int64(1)
	if atomic.CompareAndSwapInt64(&s.running, old, int64(0)) {
		select {
		case s.stop <- true:
			return
		default:
			return
		}
	}
}

func (s *ScheduledTask) task() {
	timer := time.NewTicker(s.rotationTime)
	defer timer.Stop()
	defer close(s.stop)

	ctx := gctx.New()
	uid := uuid.NewString()
	ctx = context.WithValue(ctx, web.ContextKeyAction, s.name)
	ctx = context.WithValue(ctx, web.ContextKeyRequestID, uid)
	ctx = context.WithValue(ctx, web.ContextKeyContextID, "background_task")
	ctx = context.WithValue(ctx, web.ContextKeyCallerID, s.name)
	ctx = context.WithValue(ctx, web.ContextKeyCallerType, "api")
	for {
		select {
		case <-timer.C:
			if !s.isRunning() {
				logger.Info(ctx, g.Map{"Message": "client is stopping, timer happens early"})
				return
			}
			logger.Debug(ctx, g.Map{"Message": "scheduled task is running", "TaskName": s.name})
			s.fcn()
		case <-s.stop:
			logger.Info(ctx, g.Map{"Message": "scheduled task is stopping", "TaskName": s.name})
			return
		}
	}
}

func (s *ScheduledTask) isRunning() bool {
	return atomic.LoadInt64(&s.running) != int64(0)
}

func NewScheduledTask(name string, rotationTime time.Duration, fcn entity.TaskFcn) entity.Task {
	if fcn == nil {
		panic("nil task fcn")
	}
	return &ScheduledTask{
		fcn:          fcn,
		name:         name,
		rotationTime: rotationTime,
		stop:         make(chan bool, 1),
		running:      int64(0),
	}
}
