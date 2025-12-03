package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("public/")))
	http.HandleFunc("/BMI", BMIcalcHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
}

func BMIcalcHandler(w http.ResponseWriter, r *http.Request) {
	weight, _ := strconv.Atoi(r.FormValue("weight"))
	height, _ := strconv.Atoi(r.FormValue("height"))

	height_meter := float64(height) / 100.0

	bmi := float64(weight) / (height_meter * height_meter)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//うまくhtml形式と認識されなかったためhtml形式で書いた。
	fmt.Fprintf(w, "<html<body>")
	fmt.Fprintf(w, "<p>BMIは%.2fです。</p>", bmi)
	fmt.Fprintf(w, "</body></html>")
}
