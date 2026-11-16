# LLM Wiki 說明文檔

## 系統架構

```
my-llm-wiki/
├── AGENTS.md        #【核心】AI系統提示詞與維護法則（原claude.md，亦可視為agent.md）
├── raw/             # 【原始區】人類唯一的輸入端：隨手記、剪藏、PDF、語音轉文字
│   └── archive/     # 【歸檔區】(新增) 經 AI 攝取(Ingest)完畢的原始檔案會移至此處
├── schema/          # 【結構規範】定義不同類型筆記的標準 Markdown 模板
│   ├── concept.md   # 概念類條目模板
│   ├── tool.md      # 工具類條目模板
│   └── project.md   # 專案類條目模板
└── wiki/            # 【知識區】由 AI 全權維護的結構化知識網路
    ├── concepts/    # 核心概念、理論、心智模型
    ├── tools/       # 軟體工具、程式庫、技術棧、硬體
    ├── projects/    # (新增) 進行中的個人專案、研究主題
    ├── index.md     # 【全域地圖】所有條目的分類目錄與「一句話摘要」
    └── log.md       # 【系統日誌】時間軸操作紀錄（嚴格限制：純追加 Append-only）
```

## 初始化

### 方法一：直接執行下列指令

```
Implement this exact idea files:
https://gist.github.com/karpathy/442a6bf555914893e9891c11519de94f as my complete second brain.
Guide me step-by-step: create the AGENTS.md schema file with full rules and change the file name 
to agent.md. Set up index.md and log.md, define folder conventions and show me the first ingest 
example. From now on, every interaction follows the schema.

使用架構如下（如果有額外的需要可以自己增加）：

my-llm-wiki/
├── AGENTS.md        #【核心】AI系統提示詞與維護法則（原claude.md，亦可視為agent.md）
├── raw/             # 【原始區】人類唯一的輸入端：隨手記、剪藏、PDF、語音轉文字
│   └── archive/     # 【歸檔區】(新增) 經 AI 攝取(Ingest)完畢的原始檔案會移至此處
├── schema/          # 【結構規範】定義不同類型筆記的標準 Markdown 模板
│   ├── concept.md   # 概念類條目模板
│   ├── tool.md      # 工具類條目模板
│   └── project.md   # 專案類條目模板
└── wiki/            # 【知識區】由 AI 全權維護的結構化知識網路
    ├── concepts/    # 核心概念、理論、心智模型
    ├── tools/       # 軟體工具、程式庫、技術棧、硬體
    ├── projects/    # (新增) 進行中的個人專案、研究主題
    ├── index.md     # 【全域地圖】所有條目的分類目錄與「一句話摘要」
    └── log.md       # 【系統日誌】時間軸操作紀錄（嚴格限制：純追加 Append-only）
```

### 方法二：直接使用已經建立好的 repo

* 下載 Github repo

```
git clone https://github.com/asccclass/emptywiki

cd kaparthywiki
```

* 將要檢索的原始 markdown 檔案放入 raw 目錄下
* 執行下列指令：

```
處理 raw/  或  ingest 新資料
```

## 進階建議：

1. 當 wiki/ 超過數百頁時，單靠 index.md 閱讀會吃較多 context，建議整合 [qmd](https://github.com/tobi/qmd) 做本地 BM25 + Vector 混合檢索 CLI。
2. 可配置 Obsidian 外掛 Dataview，直接以 YAML Frontmatter 自動生成動態表格。

## 參考資料
* [llm_wiki](https://github.com/nashsu/llm_wiki/tree/main)