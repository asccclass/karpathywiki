// File: my-llm-wiki/main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	SherryServer "github.com/asccclass/sherryserver"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

// WikiEntry 代表單篇知識條目的結構
type WikiEntry struct {
	Slug     string `json:"slug"`     // 檔名 (不含 .md)
	Title    string `json:"title"`    // 條目標題
	Category string `json:"category"` // 分類 (concepts/tools/projects/root)
	Path     string `json:"path"`     // 相對路徑
	Content  string `json:"content"`  // Markdown 完整內容
}

// WSRequest 定義前端傳來的 WebSocket 查詢指令
type WSRequest struct {
	Action string `json:"action"` // "search" (關鍵字搜尋) | "get" (讀取指定條目) | "index" (讀取首頁地圖)
	Query  string `json:"query"`  // 搜尋關鍵字或條目 slug
}

// WSResponse 定義回傳給前端的 WebSocket 資料格式
type WSResponse struct {
	Type    string      `json:"type"`    // "results" | "error"
	Message string      `json:"message"` // 狀態訊息
	Results []WikiEntry `json:"results"` // 查詢結果陣列
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允許本地端跨域連線
	},
}

// 搜尋並讀取 ./wiki 目錄下的 Markdown 檔案
func searchWiki(action, query string) ([]WikiEntry, error) {
	var results []WikiEntry
	wikiRoot := "./wiki"
	queryLower := strings.ToLower(strings.TrimSpace(query))

	err := filepath.WalkDir(wikiRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(strings.ToLower(d.Name()), ".md") {
			return nil
		}

		slug := strings.TrimSuffix(d.Name(), ".md")

		// 若 action 為 "index"，只回傳 index.md
		if action == "index" && slug != "index" {
			return nil
		}

		// 若 action 為 "get"，精準比對檔名 (支援 [[slug]] 點擊跳轉)
		if action == "get" && strings.ToLower(slug) != queryLower {
			return nil
		}

		contentBytes, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		content := string(contentBytes)

		// 若 action 為 "search"，比對檔名與內文是否包含關鍵字
		if action == "search" && queryLower != "" {
			if !strings.Contains(strings.ToLower(slug), queryLower) &&
				!strings.Contains(strings.ToLower(content), queryLower) {
				return nil
			}
		}

		// 解析標題與分類
		title := extractTitle(content, slug)
		category := filepath.Base(filepath.Dir(path))
		if category == "wiki" {
			category = "system"
		}

		results = append(results, WikiEntry{
			Slug:     slug,
			Title:    title,
			Category: category,
			Path:     filepath.ToSlash(path),
			Content:  content,
		})
		return nil
	})

	return results, err
}

// 從 YAML Frontmatter 或 H1 標題提取顯示名稱
func extractTitle(content, fallback string) string {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "title:") {
			return strings.TrimSpace(strings.TrimPrefix(line, "title:"))
		}
		if strings.HasPrefix(line, "# ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "# "))
		}
	}
	return fallback
}

// WebSocket 處理器
func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket Upgrade Error:", err)
		return
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		var req WSRequest
		if err := json.Unmarshal(msg, &req); err != nil {
			conn.WriteJSON(WSResponse{Type: "error", Message: "無效的 JSON 格式"})
			continue
		}

		entries, err := searchWiki(req.Action, req.Query)
		if err != nil {
			conn.WriteJSON(WSResponse{Type: "error", Message: err.Error()})
			continue
		}

		conn.WriteJSON(WSResponse{
			Type:    "results",
			Message: fmt.Sprintf("找到 %d 筆結果", len(entries)),
			Results: entries,
		})
	}
}

func main() {
	if err := godotenv.Load("envfile"); err != nil {
		fmt.Println(err.Error())
		return
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	documentRoot := os.Getenv("DocumentRoot")
	if documentRoot == "" {
		documentRoot = "www/html"
	}
	templateRoot := os.Getenv("TemplateRoot")
	if templateRoot == "" {
		templateRoot = "www/template"
	}

	// 確保目錄存在
	os.MkdirAll(documentRoot, 0755)
	os.MkdirAll(templateRoot, 0755)

	// 初始化 SherryServer
	srv, err := SherryServer.NewServer(":"+port, documentRoot, templateRoot)
	if err != nil {
		panic(err)
	}
	router := NewRouter(srv, documentRoot)
	if router == nil {
		fmt.Println("router is nil")
		return
	}

	// 註冊 WebSocket 路由與靜態檔案服務
	router.HandleFunc("/ws", wsHandler)
	// router.PathPrefix("/").Handler(http.FileServer(http.Dir(documentRoot)))

	fmt.Printf("🧠 LLM-Wiki Web 伺服器已啟動：http://localhost:%s/\n", port)
	srv.Server.Handler = router // 若要自訂跨來源保護，可在此包裝 handler
	srv.Start()
}
