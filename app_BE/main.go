package main

import (
	"log"
	"net/http"
	"public_chatter/app_BE/handlers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	handlers.InitDB()

	mux := http.NewServeMux()
	mux.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources/"))))
	mux.Handle("/api/chat_post/", handlers.EnableCORS(http.HandlerFunc(handlers.GetPostChat)))
	mux.Handle("/api/chat_delete/", handlers.EnableCORS(http.HandlerFunc(handlers.DeleteChat)))
	mux.Handle("/api/chat_edit/", handlers.EnableCORS(http.HandlerFunc(handlers.EditChat)))
	mux.Handle("/api/chat_history", handlers.EnableCORS(http.HandlerFunc(handlers.ApiChatHistoryHandler)))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
