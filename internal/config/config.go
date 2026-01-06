/*
 * Copyright (c) 2026 mincancao. All rights reserved.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

/*
* @Description:
* @Author: mincancao
* @Date: 2026/1/6 21:47
* @LastEditTime: 2026/1/6 21:47
* @LastEditors: mincancao
 */

package config

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/caomincan/web-framwork/internal/const/web"
	"github.com/caomincan/web-framwork/internal/logic/task"
	"github.com/caomincan/web-framwork/internal/logic/utils"
	"github.com/caomincan/web-framwork/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/google/uuid"
)

var (
	defaultConfig *ServerConfig
	once          sync.Once
	logger        = g.Log("appLogger")
)

func InitDefaultConfig(path string) {
	once.Do(func() {
		tmp, err := NewServerConfig(path)
		if err != nil {
			panic(err)
		}
		defaultConfig = tmp
	})
}

func GetDefaultConfig() *ServerConfig {
	if defaultConfig == nil {
		panic("config not initialized")
	}
	return defaultConfig
}

func NewServerConfig(path string) (*ServerConfig, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	data := newDefaultData()

	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	res := &ServerConfig{
		data:              data,
		filePath:          path,
		isReadFileSuccess: true,
	}
	return res, nil
}

type config struct {
	Env        string `json:"env"`
	ServerMode string `json:"serverMode"`
}

type ServerConfig struct {
	filePath          string
	isReadFileSuccess bool
	data              *config
	md5Hash           string
	mutex             sync.RWMutex
}

func newDefaultData() *config {
	return &config{
		Env:        "dev",
		ServerMode: "center",
	}
}

func (c *ServerConfig) GetEnv() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.data.Env
}

func (c *ServerConfig) GetServerMode() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.data.ServerMode
}

func (c *ServerConfig) SetEnv(env string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data.Env = env
}

func (c *ServerConfig) SetServerMode(mode string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data.ServerMode = mode
}

func (c *ServerConfig) Dump() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	bytes, err := json.Marshal(c.data)
	if err != nil {
		return err.Error()
	}
	return string(bytes)
}

func (c *ServerConfig) CalculateMD5() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	jsonStr, err := json.Marshal(c.data)
	if err != nil {
		logger.Info(context.Background(), g.Map{
			"Message":      "marshal config json error",
			"ErrorMessage": utils.SafeGetErrorMessage(err),
		})
		return
	}
	hash := md5.New()
	hash.Write(jsonStr)
	md5Hash := hex.EncodeToString(hash.Sum(nil))
	c.md5Hash = md5Hash
}

func (c *ServerConfig) GetListenUpdateTask() entity.Task {
	return task.NewScheduledTask("listen_config_change", 2*time.Second, c.update)
}

func (c *ServerConfig) update() {
	ctx := gctx.New()
	uid := uuid.NewString()
	ctx = context.WithValue(ctx, web.ContextKeyAction, "listen_config_change")
	ctx = context.WithValue(ctx, web.ContextKeyRequestID, uid)
	ctx = context.WithValue(ctx, web.ContextKeyContextID, "background_task")
	ctx = context.WithValue(ctx, web.ContextKeyCallerID, "listen_config_change")
	ctx = context.WithValue(ctx, web.ContextKeyCallerType, "api")

	file, err := os.ReadFile(c.filePath)
	if err != nil {
		logger.Info(ctx, g.Map{
			"Message":      "read config file error",
			"ErrorMessage": utils.SafeGetErrorMessage(err),
		})
		return
	}
	data := newDefaultData()
	err = json.Unmarshal(file, &data)
	if err != nil {
		logger.Info(ctx, g.Map{
			"Message":      "unmarshal config file error",
			"ErrorMessage": utils.SafeGetErrorMessage(err),
		})
	}
	// mark some filed not to compare diff
	{
		alreadyUnlock := int64(0)
		c.mutex.RLock()
		defer func() {
			old := int64(0)
			if atomic.CompareAndSwapInt64(&alreadyUnlock, old, int64(1)) {
				c.mutex.RUnlock()
			}
		}()
		// mark no change field if has
		old := int64(0)
		if atomic.CompareAndSwapInt64(&alreadyUnlock, old, int64(1)) {
			c.mutex.RUnlock()
		}
	}
	jsonStr, err := json.Marshal(data)
	if err != nil {
		logger.Info(ctx, g.Map{
			"Message":      "marshal config file error",
			"ErrorMessage": utils.SafeGetErrorMessage(err),
		})
		return
	}
	// calculate md5
	hash := md5.New()
	hash.Write(jsonStr)
	md5Hash := hex.EncodeToString(hash.Sum(nil))
	hasChange := false
	{
		c.mutex.Lock()
		defer c.mutex.Unlock()
		if md5Hash != c.md5Hash {
			hasChange = true
			logger.Info(ctx, g.Map{
				"Message": "md5Hash not match update config",
				"MD5Hash": md5Hash,
			})
			c.data = data
			c.md5Hash = md5Hash
			//c.setAppLogLevel()
		}
	}

	if hasChange {
		logger.Info(ctx, g.Map{
			"Message": "update config file success",
		})
		bytes, err := json.Marshal(data)
		if err != nil {
			return
		}
		logger.Info(ctx, string(bytes))
	} else {
		logger.Debug(ctx, g.Map{
			"Message": "config no change to update",
		})
	}
}
