# 前言

log模块的源码比较简单，适合作为学习go标准库源码的开端。

![](http://k.zol-img.com.cn/sjbbs/7692/a7691515_s.jpg)

## Logger类型结构

```
type Logger struct {
	mu     sync.Mutex // ensures atomic writes; protects the following fields
	prefix string     // prefix to write at beginning of each line
	flag   int        // properties
	out    io.Writer  // destination for output
	buf    []byte     // for accumulating text to write
}
```

* 1. out log输出的目的

##　构造函数

```
func New(out io.Writer, prefix string, flag int) *Logger {
	return &Logger{out: out, prefix: prefix, flag: flag}
}
```

## std 全局变量

```
var std = New(os.Stderr, "", LstdFlags)

```
std为非导出的全局变量，只能通过Printf Println等方法使用。

## Printf Println

```
// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	std.Output(2, fmt.Sprintf(format, v...))
}

// Println calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func Println(v ...interface{}) {
	std.Output(2, fmt.Sprintln(v...))
}

func Fatal(v ...interface{}) {
	std.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

```

这两个方法都使用std 作为logger

## 总结

log模块代码比较简单，但是通过源码的学习基本可以看出go中一个模块的组织形式。

* 1.首先定义核心的模型，Logger, 通常是导出的。 这样便于使用方定制结构的内容。

* 2. 提供构造方法 ，通常为New 参数应该是常用的属性。如果使用需要设置其他属性，可以直接使用 导出的模型定义。

* 3. 创建一个默认的对象，通常是非导出的 ，提供后面的方法使用。

* 4. 创建一些导出方法，这些方法会使用默认对象的相应方法实现。这样对于想使用默认配置的使用者，可以通过直接调用这些方法实现。 而不需要创建对象。

