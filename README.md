## クローンと初期設定

### 1. リポジトリをクローン

```bash
git clone https://github.com/t2469/AST-Generator.git
cd AST-Generator
```

### 2. Docker Composeでの起動

必要な場合のみビルドオプションをつける

コンテナを停止する際はdocker compose down
```bash
dockerc compose up -d [--build]
```

### 3. 各サーバーの起動

#### Reactの場合
```bash
  docker exec -it front ash
  npm run dev
```
#### Ginの場合
```bash
docker exec -it back bash
go run main.go
```

### その他
#### dbへのアクセス
```bash
docker exec -it db bash # dbコンテナへ入る
```

#### mysqlへのログイン

パスワードは<MYSQL_PASSWORD>を入力
```bash
mysql -u <MYSQL_USER> -p
```