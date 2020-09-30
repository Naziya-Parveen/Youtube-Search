package main

import (
	"fmt"
	"github.com/Naziya-Parveen/youtube-stats/websocket"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome")
}

func search(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	go websocket.Writer(ws)
}
func setupRoutes() {
	http.HandleFunc("/search", search)
	http.HandleFunc("/", homePage)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
func main() {

	fmt.Println("getting started")
	// items, err := youtube.GetVideos()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(items)
	setupRoutes()
}
