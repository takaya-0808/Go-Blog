# Go-Blog

## API
+ localhost:8022
+ v1
  + GET: localhost:8022/MyBlog/api/v1 => ユーザの情報を全部取得
  + GET: localhost:8022/MyBlog/api/v1/name => nameのユーザ情報を取得
+ user
  + POST: localhost:8022/MyBlog/api/user/login => ユーザをサイイン
  + POST: localhost:8022/MyBlog/api/user/register => 新たなユーザの登録
  + GET: localhost:8022/MyBlog/api/user/{userid} => 特定のユーザの検索
  + PUT: localhost:8022/MyBlog/api/user/{userid} => 特定のユーザの更新
  + DELETE: localhost:8022/MyBlog/api/user/{userid} => 特定のユーザの削除
+ article
  + GET: localhost:8022/MyBlog/api/article => 記事を全て取得
  + POST: localhost:8022/MyBlog/api/article => 記事の作成
  + GET: localhost:8022/MyBlog/api/article/{:id} => 記事の取得
  + GET: localhost:8022/MyBlog/api/article/titles => タイトル一覧取得
  + PUT: localhost:8022/MyBlog/api/article/{:id} => 記事の更新
  + DELETE: localhost:8022/MyBlog/api/article/{:id} => 記事の削除

## Requset/Responce
+ Requset
  ```
  localhost:8022/MyBlog/api/v1
  ```
+ Responce 
  ```
  {"message": ok,
  Userdata: user data}
  ```
+ Requset 
  ```
  localhost:8022/MyBlog/api/user/login
  ```
+ Responce 
  ```
  {"message":"ok",
  "token":"token_id"}
  ```

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