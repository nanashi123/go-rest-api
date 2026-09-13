# go-api-practice

Go と [Gorilla Mux](https://github.com/gorilla/mux) による学習用の REST API。記事とコメントを題材に、ルーティング・認証・ミドルウェア・DB 連携・テストを実装。

## 機能

- 記事（article）とコメント（comment）の REST API
- Google ID トークンによる認証ミドルウェア
- ロギング / トレース ID 付与のミドルウェア
- MySQL 連携（`database/sql` + `go-sql-driver/mysql`）
- 独自エラーハンドリング（`apperrors`）
- `net/http/httptest` によるテーブル駆動テスト

## エンドポイント

| メソッド | パス | 説明 |
|---|---|---|
| GET | `/hello` | 動作確認 |
| POST | `/article` | 記事の投稿 |
| GET | `/article/list` | 記事一覧 |
| GET | `/article/{id}` | 記事詳細 |
| POST | `/article/nice` | いいね追加 |
| POST | `/comment` | コメント投稿 |

## ディレクトリ構成

| ディレクトリ | 役割 |
|---|---|
| `api/` | ルーター定義・ミドルウェア |
| `controllers/` | リクエスト/レスポンス処理 |
| `services/` | ビジネスロジック |
| `repositories/` | DB アクセス |
| `models/` | データ型定義 |
| `apperrors/` | アプリ独自エラー |
| `common/` | 共通ユーティリティ |
| `_*` | 学習過程のデモ（JWT 検証、DB 接続、ミドルウェアなど） |

## セットアップ

### 1. MySQL 起動

`docker-compose.yaml` が参照する環境変数を設定して起動。

```powershell
docker compose up -d
```

- `ROOTUSER` / `ROOTPASS` — root ユーザー
- `DATABASE` — 初期データベース名
- `USERNAME` / `USERPASS` — アプリ用ユーザー

テーブル作成・初期データ投入は `_db/createTable.sql` / `_db/insertData.sql`。

### 2. アプリ起動

DB 接続用の環境変数を設定して起動。

- `DB_USER`
- `DB_PASSWORD`
- `DB_NAME`

```powershell
go run .
```

`http://localhost:8080` でアクセス可能。

### 3. テスト

```powershell
go test ./...
```

## 参考教材

[APIを作りながら進む Go中級者への道](https://techbookfest.org/product/jXDAEU1dR53kbZkgtDm9zx?productVariantID=dvjtgpjw8VDTXNqKaanTVi)（技術書典）を参考に実装。教材本文の転載は含まない。
