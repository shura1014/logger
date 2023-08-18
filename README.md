# logger

go语言日志库

# 使用

```shell
go get -u -v github.com/shura/logger
```

# 颜色
![01.png](img/01.png)

![02.png](img/02.png)

# 快速使用

### 带颜色
```go
todo := context.TODO()
logger := Default("test")
logger.Warn(todo, "%s %s", "wendell", "25")
logger.Debug(todo, "%s %s", "wendell", "25")
logger.Info(todo, "%s %s", "wendell", "25")
logger.Error(todo, "%s %s", "wendell", "25")
```

结果
![img.png](img/03.png)

### 关闭颜色
```go
func TestLogger(t *testing.T) {
	todo := context.TODO()
	logger := Default("test")
	logger.DisColor()
	logger.Warn(todo, "%s %s", "wendell", "25")
	logger.Debug(todo, "%s %s", "wendell", "25")
	logger.Info(todo, "%s %s", "wendell", "25")
	logger.Error(todo, "%s %s", "wendell", "25")
}
```

![img.png](img/04.png)

### 设置级别
```go
func TestLoggerLevel(t *testing.T) {
	todo := context.TODO()
	logger := Default("test")
	logger.Level = InfoLevel
	logger.Debug(todo, "%s %s", "wendell", "25")
	logger.Info(todo, "%s %s", "wendell", "25")
	logger.Error(todo, "%s %s", "wendell", "25")
}

=== RUN   TestLoggerLevel
[test] 2023-08-17 17:51:56 shura/logger/logger_test.go:22 INFO wendell 25
[test] 2023-08-17 17:51:56 shura/logger/logger_test.go:23 ERROR  Error Cause by:
wendell 25
--- PASS: TestLoggerLevel (0.00s)
PASS
```



### 自定义格式
```go
func TestValue(t *testing.T) {
	ctx := context.WithValue(context.Background(), "STAFF_ID", "TEST01")
	logger := Default("test")
	logger.SetContentFormat(`{[prefix::yellow]} {time:Y-M-D H:m:s::blue} {[param.STAFF_ID::red]} {line} {level} {msg}`)
	logger.Level = InfoLevel
	logger.Debug(ctx, "%s %s", "wendell", "25")
	logger.Info(ctx, "%s %s", "wendell", "25")
	logger.Error(ctx, "%s %s", "wendell", "25")
}


=== RUN   TestValue
[test] 2023-08-17 17:52:41 [TEST01] shura/logger/logger_test.go:32 INFO wendell 25
[test] 2023-08-17 17:52:41 [TEST01] shura/logger/logger_test.go:33 ERROR  Error Cause by:
wendell 25
--- PASS: TestValue (0.00s)
PASS
```



### 设置输入到文件
```go
func TestLoggerToFile(t *testing.T) {
	todo := context.TODO()
	logger := Default("test")
	logger.SetFilePath("./logger.log")
	logger.ForceColor()
	logger.Debug(todo, "%s %s", "wendell", "25")
	logger.Info(todo, "%s %s", "wendell", "25")
	logger.Error(todo, "%s %s", "wendell", "25")
}
```

### Json格式输出
```go
func TestLoggerJsonFormat(t *testing.T) {
	todo := context.TODO()
	logger := Default("test")
	//logger.SetFilePath("./logger.log")
	logger.SetLogFormat(JsonFormat)
	logger.Debug(todo, "%s %s", "wendell", "25")
	logger.Info(todo, "%s %s", "wendell", "25")
	logger.Error(todo, "%s %s", "wendell", "25")
}



=== RUN   TestLoggerJsonFormat
{"level":"DEBUG","msg":"wendell 25","prefix":"test","time":"2023-08-17 18:08:08"}
{"level":"INFO","msg":"wendell 25","prefix":"test","time":"2023-08-17 18:08:08"}
{"level":"ERROR","msg":"wendell 25","prefix":"test","time":"2023-08-17 18:08:08"}
--- PASS: TestLoggerJsonFormat (0.00s)
PASS
```


### 不需要ctx支持
```go
func TestLogger(t *testing.T) {
	logger := simple.Default("test")
	logger.Warn("%s %s", "wendell", "25")
	logger.Debug("%s %s", "wendell", "25")
	logger.Info("%s %s", "wendell", "25")
	logger.Error("%s %s", "wendell", "25")
}
```

### 默认的LOG
```go
func TestDefaultLog(t *testing.T) {
	todo := context.TODO()
	Log.Debug(todo, "%s %s", "wendell", "25")
	Log.Info(todo, "%s %s", "wendell", "25")
	Log.Error(todo, "%s %s", "wendell", "25")
}


=== RUN   TestDefaultLog
[DEFAULT] 2023-08-18 09:39:37 shura/logger/logger_test.go:60 DEBUG wendell 25
[DEFAULT] 2023-08-18 09:39:37 shura/logger/logger_test.go:61 INFO wendell 25
[DEFAULT] 2023-08-18 09:39:37 shura/logger/logger_test.go:62 ERROR  Error Cause by:
wendell 25
--- PASS: TestDefaultLog (0.00s)
PASS
```


### 自定义日志输出格式
实现接口
```go
type LoggingFormatter interface {
	Formatter(param *LoggingFormatterParam) string
	SetContentFormatStr(f string)
	SetContents(contents []*Content)
}

logger.SetLogFormat(f)
```

### 自定义日志格式输出
```go
func TestContentHandlerLog(t *testing.T) {
	todo := context.TODO()
	Log.SetContentFormat(`{time:Y/M/D H:m:s::red}`)
	Log.Debug(todo, "%s %s", "wendell", "25")
	Log.Info(todo, "%s %s", "wendell", "25")
	Log.Error(todo, "%s %s", "wendell", "25")
}

=== RUN   TestContentHandlerLog
2023/08/18 09:47:05
2023/08/18 09:47:05
2023/08/18 09:47:05
--- PASS: TestContentHandlerLog (0.00s)
PASS
```


### 自定义内容处理

```go
// 定义内容处理函数
var fun ContentFunc
fun = func(msg ...any) any {
    a := msg[0].(string)
    return a + "::hello"
}

// 注册函数
RegistryContentFunc("hello", fun)


// 使用 
Log.SetContentFormat(`{time:Y/M/D H:m:s::hello}`)
Log.Debug(todo, "")
Log.Info(todo, "")
Log.Error(todo, "")


// 结果
=== RUN   TestContentHandlerLog
2023/08/18 09:50:10::hello
2023/08/18 09:50:10::hello
2023/08/18 09:50:10::hello
--- PASS: TestContentHandlerLog (0.00s)
PASS
```



