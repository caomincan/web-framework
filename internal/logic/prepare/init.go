/*
 * Copyright (c) 2026 mincancao. All rights reserved.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

/*
* @Description:
* @Author: mincancao
* @Date: 2026/1/6 22:35
* @LastEditTime: 2026/1/6 22:35
* @LastEditors: mincancao
 */

package prepare

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/caomincan/web-framwork/internal/config"
	"github.com/caomincan/web-framwork/internal/const/web"
	"github.com/caomincan/web-framwork/internal/service/task"
	"github.com/gogf/gf/v2/frame/g"
)

const (
	defaultEnv             = "dev"
	defaultDynamicFilePath = "/manifest/mock/dynamic_server_config.json"
)

func InitSystem() {
	ctx := context.Background()
	env := g.Cfg().MustGetWithEnv(ctx, web.EnvKeyEnvType, defaultEnv)
	fmt.Printf("Current Env: %s\n", env.String())
	cur, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dynamicConfig, err := g.Cfg().Get(ctx, "other.dynamic_config_file")
	if err != nil {
		panic(err)
	}
	dynamicConfigPath := dynamicConfig.String()
	if dynamicConfigPath == "${dynamic_config_file}" {
		dynamicConfigPath = cur + defaultDynamicFilePath
	}
	fmt.Println("read config file: " + dynamicConfigPath)
	config.InitDefaultConfig(dynamicConfigPath)
	if !env.IsEmpty() {
		config.GetDefaultConfig().SetEnv(env.String())
	}
	mode, err := g.Cfg().GetWithEnv(ctx, web.EnvKeyServerMode)
	if err == nil && !mode.IsEmpty() {
		config.GetDefaultConfig().SetServerMode(mode.String())
	}
	dump := config.GetDefaultConfig().Dump()
	fmt.Printf("Current dynamic config: %s\n", dump)
	serverMode := config.GetDefaultConfig().GetServerMode()
	if serverMode == "" {
		panic(errors.New("server mode is empty"))
	}
	config.GetDefaultConfig().CalculateMD5()
	InitTasks()
}

func InitTasks() {
	task.GetBackgroundTaskManager().AddTask(config.GetDefaultConfig().GetListenUpdateTask())
}
