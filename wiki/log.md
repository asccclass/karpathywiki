# 📜 系統操作日誌 (System Log)
> ⚠️ **APPEND-ONLY**：本檔案僅供追加紀錄，嚴禁修改或刪除歷史條目。

- **[2025-05-14 10:00]** `INIT` 系統初始化完成。建立 `AGENTS.md`、`schema/`、`index.md` 與 `log.md` 基礎架構。
- **[2025-05-14 10:05]** `CREATE` 新增概念條目 [[llm-wiki]] (`wiki/concepts/llm-wiki.md`)。
  - `CREATE` 新增工具條目 [[claude-code]] (`wiki/tools/claude-code.md`)。
  - `UPDATE` 更新 `wiki/index.md` 新增 2 筆索引。
- **[2025-05-14 10:30]** `INGEST` 實作 Go + SherryServer WebSocket 知識庫查詢系統。
  - `CREATE` 新增專案結構模板 `schema/project.md`。
  - `CREATE` 新增工具條目 [[sherryserver]] (`wiki/tools/sherryserver.md`)。
  - `CREATE` 新增專案條目 [[llm-wiki-web]] (`wiki/projects/llm-wiki-web.md`)。
  - `UPDATE` 更新 `wiki/index.md` 新增 1 筆工具與 1 筆專案索引。