/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-03-28 05:53:53
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-05-29 13:25:12
 * @FilePath     : /inner-user-center-api/data/mycode/dahe/go-common/v2/go-common-v2-dh-log/dhLog.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package dhlog

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"runtime"
)

func DebugAny(params any) {
	encodeData, _ := json.MarshalIndent(params, "", "    ")

	commonWriteLog("Info", string(encodeData))
}

func Info(msg string, args ...any) {
	commonWriteLog("Info", msg, args...)
}

func InfoWithSturct(args interface{}) {
	// 将结构体序列化为 JSON
	jsonBytes, err := json.Marshal(args)
	if err != nil {
		Error("Error serializing struct to JSON: %v", err)
	}

	// 将 JSON 字符串打印到日志
	Info("Serialized struct: %s", jsonBytes)
}

func Error(msg string, args ...any) {
	commonWriteLog("Error", msg, args...)
}

func Warn(msg string, args ...any) {
	commonWriteLog("Warn", msg, args...)
}

func commonWriteLog(LogType, msg string, args ...any) {
	pc, file, line, _ := runtime.Caller(2)

	funcName := fmt.Sprintf("Caller function: %s\n", runtime.FuncForPC(pc).Name())
	fileName := fmt.Sprintf("Caller file: %s\n", file)
	LineNum := fmt.Sprintf("Caller line: %d\n", line)

	logInfo := fmt.Sprintf("\n%s%s%s", funcName, fileName, LineNum)

	msg = fmt.Sprintf(msg, args...)
	switch LogType {
	case "Info":
		slog.Info(logInfo + " " + msg)
	case "Error":
		slog.Error(logInfo + " " + msg)
	case "Warn":
		slog.Warn(logInfo + " " + msg)
	default:
		// This case will execute when number is not 1, 2, or 3
		slog.Info(logInfo + " " + msg)
	}
}
