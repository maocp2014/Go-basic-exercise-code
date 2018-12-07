package main

import (
    "fmt"
    "log"
    "net/http"
)

// 自定义类型String
type String string

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}
// 实现接口方法ServeHTTP
func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, string(s))
}

// 实现接口方法ServeHTTP
func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, fmt.Sprintf("%v", s))
} 
func main() {
    // 设置url
    http.Handle("/string", String("I'm a frayed knot."))
    http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})

    e := http.ListenAndServe("localhost:4000", nil)
    if e != nil {
        log.Fatal(e)
    }
}