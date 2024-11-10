package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TodoList struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Task     string `json:"task"`
	Priority int    `json:"priority"`
	Status   bool   `json:"status"`
	Deadline string `json:"deadline"`
}

// TableName メソッドでテーブル名を明示的に指定
func (TodoList) TableName() string {
	return "todoitems" // データベース上のテーブル名を指定
}

var db *gorm.DB

func main() {
	// .envファイルの読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 環境変数から接続情報を取得
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	// DSNの作成
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)
	fmt.Println("DSN:", dsn)

	// データベース接続
	var errDB error
	db, errDB = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errDB != nil {
		log.Fatalf("Failed to connect to the database: %v", errDB)
	} else {
		log.Println("Database connection established")
	}

	// テーブルのマイグレーション
	if err := db.AutoMigrate(&TodoList{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration completed")

	// エンドポイントの定義
	http.HandleFunc("/todos",withCORS(func(w http.ResponseWriter, r *http.Request){
		switch r.Method{
		case "GET":
			getTodos(w,r)
		case "POST":
			addTodo(w,r)
		case "DELETE":
			deleteTodo(w,r)
		default:
			http.Error(w, "Method not allow", http.StatusMethodNotAllowed)
		}
	}))

	// サーバーの起動
	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// CORS対応ミドルウェア
func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		// OPTIONSリクエストの場合、ヘッダーのみ返して終了
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// OPTIONS以外のリクエストは次のハンドラに処理を渡す
		next(w, r)
	}
}

// GETリクエストでTodoリストを取得
func getTodos(w http.ResponseWriter, r *http.Request) {
	var todos []TodoList
	result := db.Find(&todos); 
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// POSTリクエストで新しいTodoを追加
func addTodo(w http.ResponseWriter, r *http.Request) {
	var todo TodoList
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if result := db.Create(&todo); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// DELETEリクエストで指定されたIDのTodoを削除
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	
	// リクエストURLからIDを取得
	idStr := r.URL.Query().Get("id") // クエリパラメータからIDを取得
	if idStr == "" {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}

	// IDを整数に変換
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// 指定されたIDのTodoを削除
	if result := db.Delete(&TodoList{}, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// レスポンスを返す
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Todo deleted successfully"))
}