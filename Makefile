# Goプロジェクト用Makefile

# 変数定義
APP_NAME := char5742/go-ecsite-sample
DOCKER_TAG := latest

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
	@echo '  help         ヘルプを表示します'
	@echo ''
