# 挿入ソート

## description

### 挿入ソートの要素

- 先頭の要素をソート済とする
- 未ソートの部分がなくなるまで、以下の処理を繰り返す
  1. 未ソート部分の先頭から要素を取り出す（これをAとする）
  1. ソート済の部分において、Aより大きい要素を後方へ1ずつ移動する
  1. 空いた位置に取り出したAを挿入する

### 図示

[geeksforgeeks](https://www.geeksforgeeks.org/)から引用

![ソートされる様子](https://media.geeksforgeeks.org/wp-content/uploads/insertionsort.png)

## Time Complexity

O(n*2)
