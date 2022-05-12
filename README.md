# CoopeLunch

ランチを通じて、チームのコミュニケーションを生み出すアプリ

## アーキテクチャ

![CoopeMeal アーキテクチャ](https://user-images.githubusercontent.com/63759669/165153571-40f2d270-f0ff-46c8-a5b5-830a7d462f67.png)

## 立ち上げ方

**初回の場合**

1. `docker-compose build`
2. `docker-compose up -d`
3. `docker-compose exec db bash -c "chmod 0775 docker-entrypoint-initdb.d/init-database.sh"`
4. `docker-compose exec db bash -c "./docker-entrypoint-initdb.d/init-database.sh"`

**2 回目以降**

1. `docker-compose up -d`

Docker を停止するときは以下を入力
`docker-compose stop`
