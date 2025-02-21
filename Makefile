BINARY_NAME = longest-path-solver
DOCKER_IMAGE = longest-path-solver:latest

SRC_DIR = ./srcs

.DEFAULT_GOAL := docker-run

# Dockerイメージのビルド
docker-build:
	docker build -t $(DOCKER_IMAGE) .

# Docker コンテナの実行
docker-run: docker-build
	docker run --rm -i $(DOCKER_IMAGE)

# ローカル環境でのビルド(Go言語がインストール済みであることが前提)
build:
	cd $(SRC_DIR) && \
	go mod tidy && \
	go build -o $(BINARY_NAME) cmd/longest-path-solver/main.go

# ローカル環境でのビルド後に実行(Go言語がインストール済みであることが前提)
run: build
	$(SRC_DIR)/$(BINARY_NAME)

# ローカル環境でGoプロジェクト配下のテストを実行
test:
	cd $(SRC_DIR) && go test ./...

# 実行バイナリとDockerイメージの削除
clean:
	@if [ -f $(SRC_DIR)/$(BINARY_NAME) ]; then \
		rm -f $(SRC_DIR)/$(BINARY_NAME); \
	fi
	@if docker image inspect $(DOCKER_IMAGE) > /dev/null 2>&1; then \
		docker image rm -f $(DOCKER_IMAGE); \
	fi
