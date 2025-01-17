/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"bk-bcs/bcs-common/common/blog"
	"bk-bcs/bcs-common/common/conf"
	"bk-bcs/bcs-common/common/license"
	"bk-bcs/bcs-services/bcs-netservice/app"
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//loading configuration file
	cfg := app.NewConfig()
	conf.Parse(cfg)
	license.CheckLicense(cfg.LicenseServerConfig)
	//init logs
	blog.InitLogs(cfg.LogConfig)
	defer blog.CloseLogs()
	//running netservice application
	if err := app.Run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "bcs-netservice running failed: %s\n", err.Error())
		time.Sleep(5 * time.Second)
		return
	}
}
