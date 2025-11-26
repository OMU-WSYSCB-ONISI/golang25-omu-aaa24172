package main

import (
"fmt"
"net/http"
"time"
)

func main() {
http.HandleFunc("/info", info)
http.ListenAndServe(":8080", nil)
}

func info(w http.ResponseWriter, r *http.Request){
current :=time.Now().Format("21:23")
burauza :=r.Header.Get("User-Agent")
fmt.Fprintln( w, "現在時刻", current,"今使っているブラウザは",burauza, "です")
}
