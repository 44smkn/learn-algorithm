# learn-algorithm

螺旋本やLeetCodeからアルゴリズムを実装して勉強する本
GoとRustで実装していく

[![GoDoc](https://godoc.org/github.com/shibukawa/my-lib?status.svg)](https://godoc.org/github.com/44smkn/learn-algorithm)

## アルゴリズムの計算量

アルゴリズムの効率を評価する指標の一つとして、O表記法がある。  
大体はnを入力のサイズとした関数で表す。  
例えば、O(n^2)であれば、計算量はn^2に比例することになる。  
一般的には最悪の計算量で見積もることになる。

|n|logn|√n|nlogn|n^2|2^n|n!|
|---|---|---|---|---|---|---|
|5|2|2|10|25|32|120|
|100|6|10|600|10000|10^30|9.3*10^157|
