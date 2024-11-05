package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)



type TodoList struct {
    ID int `gorm:"primaryKey" json:"id"`
    Task  string `json:"task"`
    Priority int `json:"priority"`
    Status  bool `json:"status"`
    Deadline string	`json:"deadline"`
};


var TodoItems = []TodoList{
    { ID: 1, Task: "買い物に行く", Priority: 15, Status: false, Deadline: "2024-11-05" },
    { ID: 2, Task: "ジムで運動する", Priority: 3, Status: true, Deadline: "2024-11-10" },
    { ID: 3, Task: "本を読む", Priority: 0, Status: false, Deadline: "2024-11-05"},
    { ID: 4, Task: "部屋の掃除", Priority: 19, Status: false, Deadline: "2024-11-08" },
    { ID: 5, Task: "仕事のメールを返信", Priority: 1, Status: true, Deadline: "2024-11-05" },
    { ID: 6, Task: "友達とランチ", Priority: 0, Status: false, Deadline: "2024-11-15" },
    { ID: 7, Task: "映画を見る", Priority: 7, Status: false, Deadline: "2024-11-05" },
    { ID: 8, Task: "週末の予定を立てる", Priority: 12, Status: true, Deadline: "2024-11-12" },
    { ID: 9, Task: "歯医者の予約を取る", Priority: 8, Status: false, Deadline: "2024-11-05" },
    { ID: 10, Task: "レシピを考える", Priority: 20, Status: true, Deadline: "2024-11-20" },
    { ID: 11, Task: "日記を書く", Priority: 0, Status: false, Deadline: "2024-11-05" },
    { ID: 12, Task: "銀行に行く", Priority: 4, Status: false, Deadline: "2024-11-18" },
    { ID: 13, Task: "衣替えをする", Priority: 17, Status: true, Deadline: "2024-11-05" },
    { ID: 14, Task: "プレゼン資料を準備", Priority: 2, Status: false, Deadline: "2024-11-14" },
    { ID: 15, Task: "植物に水をやる", Priority: 0, Status: false, Deadline: "2024-11-05" },
    { ID: 16, Task: "ストレッチをする", Priority: 13, Status: true, Deadline: "2024-11-22" },
    { ID: 17, Task: "新しいレシピを試す", Priority: 5, Status: false, Deadline: "2024-11-05" },
    { ID: 18, Task: "オンラインコースを受講", Priority: 18, Status: false, Deadline: "2024-11-25" },
    { ID: 19, Task: "プロジェクトの進捗確認", Priority: 16, Status: true, Deadline: "2024-11-05" },
    { ID: 20, Task: "部屋の模様替えを考える", Priority: 6, Status: false, Deadline: "2024-11-30" },
};

// ※Goではコードの記述順序は関係ないので、上に書いても下に書いても構いません。
func main() {
	http.HandleFunc("/", getTodoItems)
	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getTodoItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    // ここにCORS対応コードを追加します。
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") 
	json.NewEncoder(w).Encode(TodoItems)
}