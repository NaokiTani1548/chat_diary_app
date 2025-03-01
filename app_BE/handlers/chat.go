package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"
	"public_chatter/app_BE/models"
	"strconv"
	"strings"
)

func ApiChatHistoryHandler(w http.ResponseWriter, r *http.Request) {

	DB_Connection, err := getDBConnection()
	if err != nil {
		http.Error(w, "DB接続エラー", http.StatusInternalServerError)
		log.Println("DB接続失敗:", err)
		return
	}
	defer DB_Connection.Close()

	cmd := `SELECT chat_id, chat_detail, author_id FROM chat`
	rows, err := DB_Connection.Query(cmd)
	if err != nil {
		http.Error(w, "データ取得エラー", http.StatusInternalServerError)
		log.Println("データ取得失敗:", err)
		return
	}
	defer rows.Close()

	var chats []models.Chat
	for rows.Next() {
		var chat models.Chat
		if err := rows.Scan(&chat.Chat_id, &chat.Chat_detail, &chat.Author_id); err != nil {
			http.Error(w, "データパースエラー", http.StatusInternalServerError)
			log.Println("データ処理失敗:", err)
			return
		}
		chats = append(chats, chat)
	}
	json.NewEncoder(w).Encode(chats)

}

func EditChat(w http.ResponseWriter, r *http.Request) {
	DB_Connection, err := getDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer DB_Connection.Close()

	sub := strings.TrimPrefix(r.URL.Path, "/api/chat_edit")
	_, id := filepath.Split(sub)

	I, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid chat ID", http.StatusBadRequest)
		return
	}
	var chat models.PostChat
	err2 := json.NewDecoder(r.Body).Decode(&chat)
	if err2 != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		log.Println("パース失敗:", err)
		return
	}
	cmd := "UPDATE chat SET chat_detail = ? WHERE chat_id = ?"

	_, execErr := DB_Connection.Exec(cmd, chat.Chat_detail, I)
	if execErr != nil {
		http.Error(w, "Failed to edit chat", http.StatusInternalServerError)
		return
	}

	log.Println("Edit成功:", id, ":", chat)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chat)
}

func DeleteChat(w http.ResponseWriter, r *http.Request) {
	DB_Connection, err := getDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer DB_Connection.Close()
	sub := strings.TrimPrefix(r.URL.Path, "/api/chat_delete")
	_, id := filepath.Split(sub)
	cmd := "DELETE FROM chat WHERE chat_id = ?"
	Id_num, _ := strconv.Atoi(id)
	_, err2 := DB_Connection.Exec(cmd, Id_num)
	if err2 != nil {
		http.Error(w, "Failed to delete chat", http.StatusInternalServerError)
		log.Println("DELETE失敗:", err)
		return
	}
	log.Println("DELETE成功:", id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Id_num)
}

func GetPostChat(w http.ResponseWriter, r *http.Request) {
	DB_Connection, err := getDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer DB_Connection.Close()

	var chat models.PostChat
	err2 := json.NewDecoder(r.Body).Decode(&chat)
	if err2 != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		log.Println("パース失敗:", err)
		return
	}

	cmd := `INSERT INTO chat(chat_detail,author_id)VALUES(?,1)`
	_, err3 := DB_Connection.Exec(cmd, chat.Chat_detail)
	if err3 != nil {
		http.Error(w, "Failed to insert chat", http.StatusInternalServerError)
		log.Println("INSERT失敗:", err)
		return
	}
	log.Println("INSERT成功:", chat.Chat_detail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chat)
}
