# 概要
フロントエンド部分

# 起動に関して
本アプリは`yarn`をパッケージ管理に利用している

## ローカルで起動
1. 依存パッケージをインストール
    ```
    yarn install
    ```
2. developmentモードで起動
    ```
    yarn serve
    ```
## dockerを用いて起動
`Dockerfile`のある階層で実行
```
docker build -t frontend_image .

docker container run -it -p 8080:8080 --name frontend_container frontend_image
```

## その他
コードのビルド
```
yarn build
```

コードの自動整形
```
yarn lint
```