/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-08 06:58:48
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-05-18 13:58:57
 * @FilePath     : /v2/go-common-v2-dh-log/dhLog_test.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package dhlog

import (
	"testing"
)

func TestInfo(t *testing.T) {
	Info("测试info日志")
	Info("测试info日志: %v %v", "哈哈", "呵呵")
}

func TestError(t *testing.T) {
	Error("测试error日志")
}
