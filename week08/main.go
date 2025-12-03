package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("public/")))
	http.HandleFunc("/average", averageHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}

}

func averageHandler(w http.ResponseWriter, r *http.Request) {
	var sum_score float64
	var tt int
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	str_score := r.FormValue("score")
	temp_score := strings.ReplaceAll(str_score, " ", "")
	tokuten := strings.Split(temp_score, ",")
	distri := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := range tokuten {
		tt, _ = strconv.Atoi(tokuten[i])
		if tt == 100 {
			distri[10]++
		}
		for j := 0; j < 10; j++ {
			if tt >= j*10 && tt < (j+1)*10 {
				distri[j]++
				break
			}
		}
		sum_score += float64(tt)
	}
	average := sum_score / float64(len(tokuten))
	fmt.Fprintln(w, "得点一覧:", tokuten)
	fmt.Fprintln(w, "平均点:", average)
	fmt.Fprintln(w, "得点分布")
	for i := 0; i < 10; i++ {
		fmt.Fprint(w, i*10, "点-", (i+1)*10-1, "点")
		for j := 0; j < distri[i]; j++ {
			fmt.Fprint(w, "*")
		}
		fmt.Fprintln(w, " ")
	}
	fmt.Fprint(w, "100点")
	for i := 0; i < distri[10]; i++ {
		fmt.Fprint(w, "*")
	}
}
