package main

import (
"fmt"
"net/http"
"math/rand"
"time"
)
var fortunes =[]string{
"大吉",
"中吉",
"吉",
"凶",
}
var pick int
var fortune string
func main() {
http.HandleFunc("/webfortune", fortunehandler)
http.ListenAndServe(":8080", nil)
}

func fortunehandler(w http.ResponseWriter, r *http.Request){
rand.Seed(time.Now().UnixNano())
pick = rand.Intn(len(fortunes))
fortune =fortunes[pick]
fmt.Fprintln( w, "今の運勢は", fortune, "です")
}
