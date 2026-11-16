---
title: LLM-Wiki Web 即時查詢系統
category: project
tags: [golang, websocket, pkm, web-ui, sherryserver]
created: 2025-05-14
updated: 2025-05-14
sources: [main.go, www/html/index.html]
---

# LLM-Wiki Web 即時查詢系統

> **專案目標**：為 [[llm-wiki]] 提供一個基於 WebSocket 的低延遲 Web 視覺化檢索介面，支援全文即時搜尋與 `[[雙向連結]]` 點擊跳轉。

## 系統架構與技術棧
- **後端引擎**：使用 Go 語言搭配 [[sherryserver]] 架設 HTTP 伺服器，並結合 `gorilla/websocket` 處理長連線查詢請求。
- **前端介面**：採用 Tailwind CSS 與 `marked.js`，在瀏覽器端即時解析 Markdown 並將 `[[slug]]` 語法動態轉譯為 WebSocket 查詢事件。

## 核心功能
- **即時全文檢索 (`action: search`)**：使用者輸入關鍵字時，透過 WebSocket 即時掃描 `wiki/` 目錄下的檔名與內文。
- **雙向連結導覽 (`action: get`)**：點擊文章內的 `[[知識條目]]` 標籤，無須重新整理頁面即可載入目標條目。
- **全域地圖與日誌檢視**：一鍵調閱 `index.md` 與 `log.md`，掌握第二大腦全貌。

## 相關連結 (See Also)
- 核心概念：[[llm-wiki]]
- 使用工具：[[sherryserver]]、[[claude-code]]