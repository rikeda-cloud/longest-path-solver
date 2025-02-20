# ビルドステージ
FROM golang:1.23 AS build

WORKDIR /app

COPY ./srcs/ /app/
RUN go mod tidy
RUN go build -o longest-path-solver ./main.go

# 実行ステージ
FROM alpine:latest AS runtime

COPY --from=build /app/longest-path-solver /usr/local/bin/longest-path-solver
ENTRYPOINT ["/usr/local/bin/longest-path-solver"]
