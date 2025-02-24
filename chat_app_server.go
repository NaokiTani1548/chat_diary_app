package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// HandlerFunc 型の関数を定義
func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/index.html")
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w, nil)
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

// HandlerFunc 型の関数を定義
func fuga(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "fuga")
}

func initDB() {
	var err error
	DB_Connection, err := sql.Open("mysql", os.Getenv("DB_ROLE")+":"+os.Getenv("DB_PASS")+"@/"+os.Getenv("DB_NAME"))

	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}

	// 実際に接続できるかテスト
	if err = DB_Connection.Ping(); err != nil {
		log.Fatalf("DB ping error: %v", err)
	}
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
	http.HandleFunc("/fuga", fuga)

	server.ListenAndServe()
}

// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// type HelloHandler struct{}

// // *HelloHandler がインターフェース http.Handler を実装
// func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Helloooooooooooooo, world!")
// }

// func main() {
// 	handler := HelloHandler{}

// 	server := http.Server{
// 		Addr:    ":8080",
// 		Handler: &handler, //http.Handler 型(インターフェース)を期待
// 	}
// 	server.ListenAndServe()
// }
