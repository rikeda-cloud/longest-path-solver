# longest-path-solver

## 実行方法
### Golangがインストールされている場合
```
make run
```
OR
```
cd srcs/
go mod tidy
go build -o longest-path-solver cmd/longest-path-solver/main.go
./longest-path-solver
```

### Golangがインストールされていない場合
```
make docker-run
```
OR
```
docker build -t longest-path-solver:latest .
docker run --rm -i longest-path-solver:latest
```

## ファイル構成
```
.
├── Dockerfile
├── Makefile
├── README.md
└── srcs
    ├── cmd
    │   └── longest-path-solver
    │       └── main.go                                          # アプリケーションのエントリーポイント
    ├── go.mod
    └── internal
        ├── algorithm
        │   ├── benchmark_test.go                                # ベンチマークテストを定義
        │   ├── dfs.go                                           # 深さ優先探索に関連するロジックを定義
        │   ├── find_longest_path_by_dfs.go                      # DFSを用いて最長経路を見つけるロジックを定義
        │   ├── find_longest_path_by_dfs_goroutine.go            # DFS&Goroutineを用いて最長経路を見つけるロジックを定義
        │   ├── find_longest_path_test.go
        │   └── path_node.go                                     # DFS時にStackで管理されるデータ構造を定義
        ├── graph
        │   ├── edge_id.go
        │   ├── graph.go                                         # 隣接リスト構造体を定義
        │   ├── interface.go                                     # グラフに関連するインターフェースを定義
        │   └── map_based_graph.go                               # マップベースの隣接リスト構造体を定義
        ├── input
        │   ├── convert_graph_inputs_to_graph.go                 # 入力処理用構造体スライスからグラフ構造体に変換する処理を定義
        │   ├── convert_graph_inputs_to_graph_test.go
        │   ├── convert_graph_inputs_to_map_based_graph_test.go
        │   ├── graph_input.go                                   # 1行分の入力をParse、入力処理用構造体へ変換する処理を定義
        │   ├── graph_input_test.go
        │   ├── parse_graph_inputs.go                            # 全行の入力をParse、入力処理用構造体のスライスへ変換する処理を定義
        │   └── parse_graph_inputs_test.go
        ├── output
        │   └── output_result.go
        └── stack
            ├── stack.go
            └── stack_test.go
```
