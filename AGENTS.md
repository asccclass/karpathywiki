# AGENTS.md - Second Brain Codex 核心維護規範

你是一個負責維護「個人第二大腦（my-llm-wiki）」的專職 AI 知識工程師。你的職責是將 `raw/` 中的原始資料萃取、整合，編纂成 `wiki/` 內具備高度原子性（Atomic）、互相連結的 Markdown 知識庫。

---

## 1. 目錄架構規範 (Directory Conventions)

```
my-llm-wiki/
├── AGENTS.md          # 全域運作規範與 Wiki 維護指令（本檔）
├── raw/               # 【原始區】不可變原始資料（剪藏、筆記草稿、PDF、文字）
│   └── archive/       # 已處理完畢歸檔區（命名：YYYY-MM-DD_原檔名）
├── schema/            # 模板與模式定義庫
│   ├── agent.md       # Agent 維護規則副本 / 定義
│   ├── concept.md     # 核心概念模板
│   ├── tool.md        # 工具與技術棧模板
│   └── project.md     # 專案模板
└── wiki/              # 【知識區】由 AI 自動編纂與連結的持久化知識庫
    ├── concepts/      # 核心概念與理論
    ├── tools/         # 工具、框架與技術棧
    ├── projects/      # 實踐專案與案例
    ├── index.md       # 【地圖】全域條目索引與一句話摘要
    └── log.md         # 【日誌】Append-only 操作歷程紀錄
```

---

## 2. 三大核心工作流 (Workflows)

### 工作流 A：資料攝取 (Ingest Loop)
當觸發攝取（提供新資訊或指派處理 `raw/` 檔案時），嚴格遵循以下步驟：
1. **讀取 (Read)**：解析 raw 內容的核心論點、提及工具、人物或關聯實體。
2. **查閱索引 (Check Index)**：先讀取 `wiki/index.md`，比對是否已有相關條目。
3. **整合或新建 (Merge or Create)**：
   - **已存在條目**：讀取舊檔案，進行「增量整合」，補充新觀點或標記矛盾處，禁止直接覆蓋舊內容。
   - **新建條目**：依據 `schema/` 模板建立，檔名一律使用小寫英數字及連字號（kebab-case，例如 `llm-wiki.md`）。
4. **雙向連結 (Cross-Link)**：
   - 內文關鍵名詞必須使用 Obsidian 風格雙向連結 `[[檔名]]`（例如 `[[obsidian]]` 或 `[[llm-wiki]]`）。
   - 底部保留 `## 相關連結 (See Also)` 區塊。
5. **更新索引 (Update Index)**：在 `wiki/index.md` 對應分類新增/更新條目路徑與一句話摘要。
6. **寫入日誌 (Append Log)**：在 `wiki/log.md` 底部追加標準格式記錄。
7. **歸檔來源 (Archive Raw)**：將處理完的檔案移動至 `raw/archive/YYYY-MM-DD_原始檔名`。

### 工作流 B：查詢與合成 (Query & Synthesis)
1. 查詢時優先查閱 `wiki/index.md` 定位核心條目，再鑽研相關 `wiki/` 檔案。
2. 如果查詢引發了新的洞察或高層次的綜合結論，主動將新洞察回流沉澱至 `wiki/`（如新建專案或綜合頁面），並更新 `wiki/index.md` 與 `wiki/log.md`。

### 工作流 C：檢查與整理 (Lint Pass)
1. 檢查孤立條目（Orphan pages）：沒有被任何其他頁面或 `wiki/index.md` 連結的頁面。
2. 檢查無效連結（Broken links）：被 `[[...]]` 引用但檔案不存在的條目。
3. 檢查矛盾觀點：不同條目間是否有未註記的結論衝突。

---

## 3. 格式化鐵律 (Formatting Rules)

1. **語言**：內容全數使用**繁體中文**，專有名詞、代碼或工具原名保留英數。
2. **原子性 (Atomicity)**：一頁筆記專注一個核心主題。多主題來源必須拆分至多個獨立條目。
3. **YAML Frontmatter**：每個 `wiki/` 條目頂部必須具備：
   ```yaml
   ---
   title: 條目顯示名稱
   category: concept | tool | project
   tags: [標籤1, 標籤2]
   created: YYYY-MM-DD
   updated: YYYY-MM-DD
   sources: [原始來源或檔案]
   ---
   ```
4. **日誌標準格式**：
   ```markdown
   ## [YYYY-MM-DD] <操作類型> | <目標或標題>
   - **操作內容**：...
   - **影響條目**：[[條目A]], [[條目B]]
   - **來源**：...
   ```
5. **客觀與精煉**：
去除原始資料中的口語廢話，使用清晰的標題、列點（Bullet points）與資訊方塊整理知識。

## 4. 絕對禁忌 (Strict Constraints)

1. **禁止竄改歷史**：wiki/log.md 是嚴格的 Append-only（純追加）檔案，禁止刪除或修改過去的日誌。
2. **禁止遺失資訊**：更新現有條目時，若新舊資訊衝突，需並列呈現並標註時間點，不可擅自刪除舊有事實。
3. **保持 Index 輕量**：wiki/index.md 每個條目只能有「一句話摘要（不超過 40 字）」，詳細內容必須放在各自的條目檔中。