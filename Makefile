# Goプロジェクト用Makefile

# 変数定義
APP_NAME := char5742/go-ecsite-sample
DOCKER_TAG := latest
MIGRATE_TOOL := github.com/み/migrate/v4/cmd/migrate@latest
DB_HOST := localhost
DB_PORT := 5435
DB_USER := postgres
DB_PASSWORD := postgres
DB_NAME := app

# タスク定義
.PHONY: all clean build run test

# デフォルトタスク
.DEFAULT_GOAL := help

# ビルド
build:
	cd docker &&
	docker build  -t $(APP_NAME):$(DOCKER_TAG) --target deploy .

build-local:
	cd docker &&
	docker build  --no-cache

# 実行
run: 
	docker compose up -d

# 停止
stop:
	docker compose down

# テスト実行
test:
	go test -race -shuffle=on -coverprofile=coverage.out ./...

# 静的解析
lint:
	golangci-lint run

# 生成
# 現状はmockの生成のみ
generate:
	go generate ./...

# マイグレーション用のファイルを作成
# name: バージョン名
db-migrate-create:
	go run ${MIGRATE_TOOL} create -ext sql -dir db/migrations -seq $(name)

# マイグレーションを実行
db-migrate-up:
	go run cmd/migrate/main.go up

# マイグレーションを戻す
db-migrate-down:
	go run cmd/migrate/main.go down


# ヘルプ
help:
	@echo 'Usage: make [task]'
	@echo ''
	@echo 'Tasks:'
	@echo '  build        Dockerイメージをビルドします'
	@echo '  run          Dockerコンテナを起動します'
	@echo '  stop         Dockerコンテナを停止します'
	@echo '  test         テストを実行します'
	@echo '  lint         静的解析を実行します'
	@echo '  generate     生成を実行します'
	@echo '  help         ヘルプを表示します'
	@echo '  db-migrate-create  マイグレーションファイルを作成します'
	@echo '  db-migrate-up      マイグレーションを実行します'
	@echo '  db-migrate-down    マイグレーションを戻します'
	@echo ''
