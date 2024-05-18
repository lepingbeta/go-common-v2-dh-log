/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-08 06:58:48
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-05-18 14:08:17
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

func TestDebugAny(t *testing.T) {
	// Person 结构体定义了一个人的基本信息
	type Person struct {
		Name  string
		Age   int
		Email string
	}

	// 创建一个 Person 实例并初始化
	person := Person{
		Name:  "张三",
		Age:   30,
		Email: "zhangsan@example.com",
	}

	DebugAny(person)
}

func TestError(t *testing.T) {
	Error("测试error日志")
}
