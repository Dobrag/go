package gee
import(
    // "fmt"
    "net/http"
)
// type HandlerFunc func(http.ResponseWriter,*http.Request)
type HandlerFunc func(c *Context)
type Engine struct{
    // router map[string]HandlerFunc
    router *router
}

// func New() *Engine{
//     return &Engine{router:make(map[string]HandlerFunc)}
// }

func New() *Engine{
    return &Engine{router:newRouter()}
}

func (engine *Engine)addRouter(method string,pattern string,handler HandlerFunc){
    // key := methed + "-" + pattern
    // engine.router[key] = handler
    engine.router.addRouter(method,pattern,handler)
}

func (engine *Engine)GET(pattern string,handler HandlerFunc){
    engine.addRouter("GET",pattern,handler)
}

func (engine *Engine)POST(pattern string,handler HandlerFunc){
    engine.addRouter("POST",pattern,handler)
}

func (engine *Engine)RUN(addr string) (err error){
    return http.ListenAndServe(addr,engine)
}

// func (engine *Engine)ServeHTTP(w http.ResponseWriter,req *http.Request){
//     key := req.Method + "-" + req.URL.Path
//     if handler , ok := engine.router[key]; ok {
//         handler(w,req)
//     }else{
//         fmt.Fprintf(w,"404 NOT FOUND %s\n",req.URL)
//     }
// }

func (engine *Engine)ServeHTTP(w http.ResponseWriter,req *http.Request){
     c := newContext(w,req)
     engine.router.handle(c)
}