# 前言

近期使用go开发api后台服务，对http模块源码有初步了解，所以整理总结一下。

![](https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1558197105388&di=680d81c712524ca66c99ebe905526b98&imgtype=0&src=http%3A%2F%2Fi0.hdslb.com%2Fbfs%2Farchive%2Fa82faf6879551dbbc48cb7c5e60ee47141673780.jpg)

## server端核心模型包括

* **Server** 类型结构，代表了一个指定端口下的服务。

同时也保存了tls，超时等配置。最重要的是一个handler处理器。

```
type Server struct {
	Addr    string  // TCP address to listen on, ":http" if empty
	Handler Handler // handler to invoke, http.DefaultServeMux if nil
	TLSConfig *tls.Config
	ReadTimeout time.Duration
        //...
}
```

ListenAndServe方法开启监听，等待连接。
server 监听客户端请求，启动goroutine处理。

```
func (srv *Server) Serve(l net.Listener) error {
	//...
	for {
		rw, e := l.Accept()
		// ...
		c := srv.newConn(rw)
		go c.serve(ctx)
	}
}
```
* **Handler** 接口，响应一个http请求
```
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

* ServeMux : http请求多路复用器，实现了 Handler接口，它的实现为: 转发请求到匹配的handler处理。
```
type ServeMux struct {
	mu    sync.RWMutex
	m     map[string]muxEntry
	es    []muxEntry // slice of entries sorted from longest to shortest.
	hosts bool       // whether any patterns contain hostnames
}

// ServeHTTP dispatches the request to the handler whose
// pattern most closely matches the request URL.
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
	if r.RequestURI == "*" {
		if r.ProtoAtLeast(1, 1) {
			w.Header().Set("Connection", "close")
		}
		w.WriteHeader(StatusBadRequest)
		return
	}
	h, _ := mux.Handler(r)
	h.ServeHTTP(w, r)
}
```

* DefaultServeMux:  是ServeMux类型一个实力， 作为一个全局对象导出。当新建的server对象，handler为空时，使用DefaultServeMux。
```
// DefaultServeMux is the default ServeMux used by Serve.
var DefaultServeMux = &defaultServeMux

var defaultServeMux ServeMux
```

可以通过HandleFunc方法，向全局多路复用器中加入路由模式及其handler。
处理请求时根据所有加入的路由模式中匹配。
```
// HandleFunc registers the handler function for the given pattern
// in the DefaultServeMux.
// The documentation for ServeMux explains how patterns are matched.
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}
```

注册pattern, 如果以/结尾，会加入到排序的切片中，在map中匹配不到路由时，使用它，做最长匹配。
```
if pattern[len(pattern)-1] == '/' {
		mux.es = appendSorted(mux.es, e)
	}
```

路由匹配规则，先在map中精确匹配。
匹配失败则 用到排序的slice 进行最长匹配。  

```
func (mux *ServeMux) match(path string) (h Handler, pattern string) {
	// Check for exact match first.
	v, ok := mux.m[path]
	if ok {
		return v.h, v.pattern
	}

	// Check for longest valid match.  mux.es contains all patterns
	// that end in / sorted from longest to shortest.
	for _, e := range mux.es {
		if strings.HasPrefix(path, e.pattern) {
			return e.h, e.pattern
		}
	}
	return nil, ""
}
```

## 总结：
* 1.   go的http server模块同时也包含了处理socker 连接的代码，所以不需要像java需要额外的应用服务器，例如tomcat。 go更易于部署，开发迭代。
我认为一部分原因 是得益于go对并发的原生支持，go的并发模型。go开发时，可以使用goroute简单 并且低代价的方式，即可获得高并发的支持。

* 2. go server中路由匹配 力度太大，一个url匹配一个handler，不区分http方法，也无法在url中加入参数匹配。
所以 路由部分 还是建议使用开源模块
[https://github.com/gorilla/mux](https://github.com/gorilla/mux)
[https://github.com/julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)
mux源码分析参考 : [这里]([https://www.jianshu.com/p/e36fed97d369](https://www.jianshu.com/p/e36fed97d369)
)



> 参考：
> [https://www.godoc.org/net/http](https://www.godoc.org/net/http)


[原文链接](http://bestlang.cn:8080/)
