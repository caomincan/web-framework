/*
 * Copyright (c) 2026 mincancao. All rights reserved.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

/*
* @Description:
* @Author: mincancao
* @Date: 2026/1/6 22:04
* @LastEditTime: 2026/1/6 22:04
* @LastEditors: mincancao
 */

package utils

import "strings"

func SafeGetErrorMessage(err error) string {
	if err == nil {
		return ""
	}
	return strings.Replace(err.Error(), "\"", "'", -1)
}
