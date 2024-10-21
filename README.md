# 👾Pure LISP

フルスクラッチで実装をした純LISP処理系です。(恐らく)   
Goで書かれています。    

HelloWorldのソースコード   
```
(define str HelloWorld)
(print str)
```

# ✅仕様
動的スコープです。   
意味的・構文的に不正なソースコードを入れた際、動作未定義です。  
## データ型
データ型には「atom」と「lambda」と「pair」があります。

### atom
|型|説明|
|---|---|
|nil|nilを表します|
|T|Trueを表します|
|Other|define構文で定義されたもの以外のすべてです|
### lambda
ラムダ構文で定義された関数です。

### pair
二つのデータを保持できる型です。   
以下、例    
```
(A B)
(A nil)
(A (B C))   
(A (lambda (arg1 arg2) (eq arg1 arg2)))
```

## 構文
|構文|書き方|動作|
|---|---|---|
|atom|(atom A)|Aのデータ型がatomならT、それ以外ならnilを返します|
|eq|(eq A B)|AとBが等しければ、T、それ以外ならnilを返します|
|car|(car (cons A B))|引数のペアの右側を返します|
|cdr|(cons A B)|引数のペアの左側を返します|
|cons|(cons A B)|ペア(A B)を返します|
|if|(if expr A B)|exprがTならA、それ以外ならBを実行し、実行した方の戻り値を返します|
|quote|(quote (A B C))|構文木をペア型のデータへ変換します。例の場合、(A (B (C nil)))のペアが返されます|
|lambda|(lambda (arg1 arg2) (eq arg1 arg2))|lambda型のデータを作成し、返します|
|print|(print A)|Golangの機能を使用し、Aを標準表示します|

# ⇩nstall
```
go install github.com/PenguinCabinet/PureLISP@latest
```

# 実行
```
PureLISP src-file.lisp
```

# テスト
```
go test
```
## 🎫LICENSE

[MIT](./LICENSE)

## ✍Author

[PenguinCabinet](https://github.com/PenguinCabinet)
