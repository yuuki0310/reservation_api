# ディレクトリ構成

```
myapp/
├── cmd/
│   └── server/
│       └── main.go            # エントリーポイント (アプリケーションの起動)
├── config/
│   └── config.go              # 設定管理（環境変数や設定ファイルの読み込み）
├── controllers/
│   └── user_controller.go     # コントローラ (ハンドラー関数)
├── models/
│   └── user.go                # モデル (データ構造、ORM)
├── services/
│   └── user_service.go        # ビジネスロジック（サービス層）
├── repositories/
│   └── user_repository.go     # データベースアクセスロジック (リポジトリ層)
├── routes/
│   └── routes.go              # ルーティング (エンドポイントの定義)
├── middlewares/
│   └── auth_middleware.go     # ミドルウェア (認証やロギングなど)
├── utils/
│   └── utils.go               # ユーティリティ関数 (共通のヘルパー関数)
├── Dockerfile                 # Dockerファイル
├── go.mod                     # Go モジュールファイル
└── go.sum                     # Go 依存関係ファイル
```
