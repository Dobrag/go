package main 
import (
    // "fmt"
    "net/http"
    "project/gee"
)

func main(){
    r := gee.New()
    r.GET("/",func(w http.ResponseWriter,req *http.Request){
        fmt.Fprintf(w,"URL.Path=%q\n",req.URL.Path)
    })
    r.GET("/hello",func(w http.ResponseWriter,req *http.Request){
        for k ,v := range(req.Header){
            fmt.Fprintf(w,"Hearder%q=%q\n",k,v)
        }
    })
    r.GET("/xiaomin",func(w http.ResponseWriter,req *http.Request){
        for k ,v := range(req.Header){
            fmt.Fprintf(w,"Hearder%q=%q\n",k,v)
        }
    })
    r.RUN(":9999")
}
func main(){
    r := gee.New()
    r.GET("/",func(c *gee.Context){
        c.HTML(http.StatusOK,"<h1>Hello Gee</h1>")
    })

    r.GET("/hello",func(c *gee.Context){
        c.String(http.StatusOK,"hello $s,you at %s\n",c.Query("name"),c.Path)
    })

    r.POST("/login", func(c *gee.Context) {
        c.JSON(http.StatusOK, gee.H{
            "username": c.PostForm("username"),
            "password": c.PostForm("password"),
        })
    })

    r.RUN(":9999")
}

// package main 
// import(
//    "sync"
//    "time"
//    "fmt"
// )
// var m sync.Mutex
// var set = make(map[int]bool,100)
// func printOnce(num int){
//     // m.Lock()
//     // a , b := set[100]
//     // fmt.Println(a , b)
//     if _,exist := set[num]; !exist{
//         fmt.Println(num)
//     }
//     set[num] = true
//     // m.Unlock()
// } 

// func main(){
//     for i := 0 ;i < 10;i++{
//         go printOnce(100)
//     }
//     time.Sleep(time.Second)
//      a , b := set[100]
//     fmt.Println(a , b)
// }