package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.FormValue("p")+"さんの運勢は「"+omikuji()+"」です")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func omikuji() string {
	result := rand.Intn(6) + 1

	switch result {
	case 1:
		return "大吉"
	case 2:
		return "中吉"
	case 3:
		return "小吉"
	case 4:
		return "吉"
	case 5:
		return "凶"
	case 6:
		return "大凶"
	default:
		return "エラー"
	}
}
