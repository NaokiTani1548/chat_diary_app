package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Chat struct {
	Chat_id     int
	Chat_detail string
	Author_id   int
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	DB_Connection, _ := sql.Open("mysql", os.Getenv("DB_ROLE")+":"+os.Getenv("DB_PASS")+"@/"+os.Getenv("DB_NAME"))
	defer DB_Connection.Close()
	cmd := `SELECT * FROM chat`
	rows, err := DB_Connection.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	var body []Chat
	for rows.Next() {
		var chat Chat
		err := rows.Scan(&chat.Chat_id, &chat.Chat_detail, &chat.Author_id)
		if err != nil {
			log.Fatalln(err)
		}
		body = append(body, chat)
	}

	t, err := template.ParseFiles("views/index.html")
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w, body)
}
func getPostChat(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	DB_Connection, _ := sql.Open("mysql", os.Getenv("DB_ROLE")+":"+os.Getenv("DB_PASS")+"@/"+os.Getenv("DB_NAME"))
	defer DB_Connection.Close()
	v := r.FormValue("chat")
	cmd := `INSERT INTO chat(chat,author_id)VALUES(?,1)`
	DB_Connection.Exec(cmd, v)
	http.Redirect(w, r, "/", http.StatusFound)
}

func deleteChat(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	DB_Connection, _ := sql.Open("mysql", os.Getenv("DB_ROLE")+":"+os.Getenv("DB_PASS")+"@/"+os.Getenv("DB_NAME"))
	defer DB_Connection.Close()
	cmd := "DELETE FROM chat WHERE chat_id = ?"
	i := r.FormValue("chat_delete")
	var I int
	I, _ = strconv.Atoi(i)
	DB_Connection.Exec(cmd, I)
	http.Redirect(w, r, "/", http.StatusFound)
}

func initDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	DB_Connection, err := sql.Open("mysql", os.Getenv("DB_ROLE")+":"+os.Getenv("DB_PASS")+"@/"+os.Getenv("DB_NAME"))

	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}
	if err = DB_Connection.Ping(); err != nil {
		log.Fatalf("DB ping error: %v", err)
	}

	cmd := `INSERT INTO author (author_id) VALUES (1);`
	DB_Connection.Exec(cmd)
}

func main() {
	initDB()

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources/"))))
	server := http.Server{
		Addr:    ":8080",
		Handler: nil, // DefaultServeMux を使用
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/chat/", getPostChat)
	http.HandleFunc("/chat_delete/", deleteChat)

	server.ListenAndServe()
}
