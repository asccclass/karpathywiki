# File: my-llm-wiki/wiki/tools/sherryserver.md
---
title: SherryServer
category: tool
tags: [golang, web-server, http, websocket]
created: 2025-05-14
updated: 2025-05-14
sources: [https://github.com/asccclass/sherryserver]
---

# SherryServer

> **定位**：由 `asccclass` 開發的 Go 語言輕量級 Web 伺服器框架，封裝了標準路由、靜態檔案服務與常用中介軟體，適合快速建構微服務與 Web 應用。

## 核心功能與特色
- **快速初始化**：透過 `sherryserver.NewServer(port, docRoot, templateRoot)` 即可快速啟動具備靜態目錄與模板支援的 HTTP 伺服器。
- **高擴充性路由**：公開 `server.Router` 介面，可無縫掛載自訂的 `http.HandlerFunc` 或整合 `gorilla/websocket` 實現雙向即時通訊。

## 使用工作流 / 最佳實踐
- 在 [[llm-wiki-web]] 專案中擔任後端核心伺服器，負責提供前端 HTML 介面並透過 `/ws` 端點即時檢索 `wiki/` 目錄下的 Markdown 知識條目。

## 相關連結 (See Also)
- 應用專案：[[llm-wiki-web]]
- 核心概念：[[llm-wiki]]