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

Reactの場合
```bash
  docker exec -it front ash
  npm run dev
```

