* 早めのテストで大事、範囲が広がれば広がるほど、テストがつらくなる
* 言語処理系の場合、入出力が分かりやすくて、TDDもやりやすかった
* 静的スコープより動的スコープの方がシンプルだという話だったが、実装してみると、確かに簡単だった
```
仮引数に実引数を入れる
関数呼び出し
仮引数のデータを消去
```
でいけるため

* quoteとconsの違いがよくわからんかった。用途と同じだし、シュガー構文にしか見えない
(quote A B C D)は
(cons A (cons B (cons C (cons D nil))))に書き換えられるのでは？
なんで必須みたいな扱いになっているんだろう

* quoteとlambdaって書き換え可能じゃないか？
```
(define f (lambda (A) (+ A 1)))
(f 1)
```
```
(define lambda (quote (+A 1))
```
