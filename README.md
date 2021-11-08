# Go-Blog

## API
+ localhost:8022
+ GET: localhost:8022/MyBlog/api/v1 => ユーザの情報を全部取得
+ GET: localhost:8022/MyBlog/api/v1/name => nameのユーザ情報を取得
+ POST: localhost:8022/MyBlog/api/auth/login => ユーザをサイイン
+ POST: localhost:8022/MyBlog/api/auth/register => 新たなユーザの登録

## usege

docker-composeの起動
```
docker-compose build
docker-compose up -d
```
testを行う
```
go test ~
```
* 現状、POSTのtestの実行ができていない
## Reference