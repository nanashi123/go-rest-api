# go-api-practice

Go と [Gorilla Mux](https://github.com/gorilla/mux) を使って実装した、学習用の REST API です。記事とコメントを題材に、ルーティング・認証・ミドルウェア・DB 連携・テストといった Web API の基本要素を一通り実装しています。

## 主な機能

- 記事（article）とコメント（comment）の REST API
- Google ID トークンによる認証ミドルウェア
- ロギング / トレース ID 付与のミドルウェア
- MySQL との連携（`database/sql` + `go-sql-driver/mysql`）
- 独自のエラーハンドリング（`apperrors`）
- `net/http/httptest` とテーブル駆動によるテスト

## エンドポイント

| メソッド | パス | 説明 |
|---|---|---|
| GET | `/hello` | 動作確認用 |
| POST | `/article` | 記事の投稿 |
| GET | `/article/list` | 記事一覧の取得 |
| GET | `/article/{id}` | 記事詳細の取得 |
| POST | `/article/nice` | いいねの追加 |
| POST | `/comment` | コメントの投稿 |

## ディレクトリ構成

| ディレクトリ | 役割 |
|---|---|
| `api/` | ルーター定義とミドルウェア |
| `controllers/` | リクエスト/レスポンスのハンドリング |
| `services/` | ビジネスロジック |
| `repositories/` | DB アクセス |
| `models/` | データ型の定義 |
| `apperrors/` | アプリケーション独自エラー |
| `common/` | 共通ユーティリティ |
| `_*` | 学習過程のデモ（JWT 検証、DB 接続、ミドルウェアなど） |

## セットアップ

### 1. MySQL を起動

`docker-compose.yaml` で使う環境変数を用意した上で、コンテナを起動します。

```powershell
docker compose up -d
```

必要な環境変数（値はローカルで設定してください）:

- `ROOTUSER` / `ROOTPASS` — MySQL の root ユーザー
- `DATABASE` — 初期データベース名
- `USERNAME` / `USERPASS` — アプリ用ユーザー

テーブル作成と初期データ投入には `_db/createTable.sql` / `_db/insertData.sql` を利用します。

### 2. アプリを起動

アプリが DB 接続に使う環境変数を設定します。

- `DB_USER`
- `DB_PASSWORD`
- `DB_NAME`

```powershell
go run .
```

起動後、`http://localhost:8080` でアクセスできます。

### 3. テスト

```powershell
go test ./...
```

## 参考教材

- [APIを作りながら進む Go中級者への道](https://techbookfest.org/product/jXDAEU1dR53kbZkgtDm9zx?productVariantID=dvjtgpjw8VDTXNqKaanTVi)（技術書典）

本リポジトリは学習目的で上記教材を参考に実装したものです。教材本文の転載は含みません。
