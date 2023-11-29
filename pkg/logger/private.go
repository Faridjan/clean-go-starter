package logger

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

var (
	Red     = color("\033[1;31m")
	Green   = color("\033[1;32m")
	Yellow  = color("\033[1;33m")
	Magenta = color("\033[1;35m")
)

func color(colorString string) func(...interface{}) string {
	sprint := func(kv ...interface{}) string {
		return colorString + fmt.Sprint(kv...) + "\033[0m"
	}
	return sprint
}

func body(level, clr string, kv ...interface{}) string {
	var result strings.Builder

	result.WriteString("level=")
	result.WriteString(level)
	result.WriteString(" svc=")
	result.WriteString(svcNameSingleton)
	result.WriteString(" clr=")
	result.WriteString(clr)

	if len(kv) > 0 {
		kvStr := make([]string, 0, len(kv))
		for i, keyValue := range kv {
			str := fmt.Sprint(keyValue)
			if (i+1)%2 != 0 {
				str = str + ": "
			} else {
				str = str + "; "
			}

			kvStr = append(kvStr, str)
		}

		message := strings.Join(kvStr, "")
		message = strings.TrimRight(message, "; ")
		message = strings.TrimRight(message, ": ")

		result.WriteString(" msg=")
		result.WriteString(`"` + message + `"`)
	}

	return result.String()
}

func caller(depth int) string {
	pc, file, line, _ := runtime.Caller(depth)
	fName := runtime.FuncForPC(pc).Name()

	idx := strings.LastIndexByte(file, '/')
	idxF := strings.LastIndexByte(fName, '.')

	return file[idx+1:] + "." + fName[idxF+1:] + ":" + strconv.Itoa(line)
}
