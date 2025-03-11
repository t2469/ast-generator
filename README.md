# AST-Generator
## 抽象構文木(AST)の可視化
![ast-generator](images/ast-generator.png)

## クローンと初期設定

### 1. リポジトリをクローン

```bash
git clone https://github.com/t2469/AST-Generator.git
cd AST-Generator
```

### 2. Docker Composeでの起動

必要な場合のみビルドオプションを付ける  
コンテナを停止する際は `docker compose down`

```bash
docker compose up -d [--build]
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
#### DBへのアクセス

```bash
docker exec -it db bash
```

#### MySQLへのログイン

パスワードは `<MYSQL_PASSWORD>` を入力

```bash
mysql -u <MYSQL_USER> -p
```

---
## サンプルコード

### Bash

```bash
#!/bin/bash
echo "Hello, World!"
```

### C

```c
#include <stdio.h>
int main() {
    printf("Hello, World!\n");
    return 0;
}
```

### C++

```cpp
#include <iostream>
int main() {
    std::cout << "Hello, World!" << std::endl;
    return 0;
}
```

### CSS

```css
body {
    background-color: #fff;
    color: #333;
}
```

### Dockerfile

```dockerfile
FROM alpine:latest
CMD ["echo", "Hello, World!"]
```

### Go

```go
package main
import "fmt"
func main() {
    fmt.Println("Hello, World!")
}
```

### HTML

```html
<!DOCTYPE html>
<html>
<head>
    <title>Hello, World!</title>
</head>
<body>
    <h1>Hello, World!</h1>
</body>
</html>
```

### Java

```java
public class Main {
    public static void main(String[] args) {
        System.out.println("Hello, World!");
    }
}
```

### JavaScript

```javascript
console.log("Hello, World!");
```

### Kotlin

```kotlin
fun main() {
    println("Hello, World!")
}
```

### PHP

```php
<?php
echo "Hello, World!";
?>
```

### Python

```python
print("Hello, World!")
```

### Ruby

```ruby
puts "Hello, World!"
```

### Rust

```rust
fn main() {
    println!("Hello, World!");
}
```

### SQL

```sql
SELECT 'Hello, World!';
```

### YAML

```yaml
message: "Hello, World!"
```
