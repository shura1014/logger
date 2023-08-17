package logger

import (
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

var excludeDir string
var bashDir string

func init() {
	_, file, _, _ := runtime.Caller(0)
	bashDir = sourceDir(file)
}

func sourceDir(file string) string {
	dir := filepath.Dir(file)
	excludeDir = dir
	dir = filepath.Dir(dir)
	s := filepath.Dir(dir)
	return filepath.ToSlash(s) + "/"
}

// GetFileNameAndLine
// skip需要跳过多少行
// 默认跳过本方法以及logger中的方法，如果还自行包装了日志方法，需要计算跳过的数量
func GetFileNameAndLine(skip int) string {
	start := 2 + skip
	for i := start; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok && (!strings.HasPrefix(file, excludeDir) || strings.HasSuffix(file, "_test.go")) {
			return strings.TrimPrefix(file+":"+strconv.FormatInt(int64(line), 10), bashDir)
		}
	}

	return ""
}
