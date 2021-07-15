# 概要
バックエンド部分

# 起動に関して
## ローカルで起動
1. 依存パッケージをインストール
    ```
    go mod download
    ```
2. 起動
    ```
    go run server.go
    ```

## dockerを用いて起動
`Dockerfile`のある階層で実行
```
docker build -t backend_image .

docker container run -it -p 1323:1323 --name backend_container backend_image
```

## その他
コードのビルド
```
go build
```