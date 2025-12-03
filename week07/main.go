package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("public/")))
	http.HandleFunc("/cal02", radiocalcHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}

}

func radiocalcHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}

	a, _ := strconv.Atoi(r.FormValue("a"))
	b, _ := strconv.Atoi(r.FormValue("b"))

	switch r.FormValue("cal0") {
	case "+":
		fmt.Fprintln(w, a+b)
	case "-":
		fmt.Fprintln(w, a-b)
	case "*":
		fmt.Fprintln(w, a*b)
	case "/":
		if b == 0 {
			fmt.Fprintln(w, "計算できません")
			break
		}
		fmt.Fprintln(w, float64(a)/float64(b))
	}
}
