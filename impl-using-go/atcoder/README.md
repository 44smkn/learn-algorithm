# 知見まとめ

<!-- TODO: 目次を挿入する-->

## modは計算途中で何度挟んでもよい

答えが大きくなる場合に「その量を 1000000007 で割ったあまりを求めよ」という問題が出されることがある。  
割る数は`1000000007`以外でもよく、ただその数が手頃な素数であるよう。  

<u>足し算・引き算・掛け算の場合</u>  

**計算の途中過程で積極的にあまりをとってよい**

[playground](https://play.golang.org/p/TMqY4ZR9yMe)

```go
const mod = 1000000007

func main() {

    a := 11111
    b := 12345
    c := 98765

    fmt.Println("まとめてmod: ", a*b*c%mod)    // 130265846
    fmt.Println("都度mod　 : ", a*b%mod*c%mod) // 130265846
}
```

### 該当する問題

- [abc177c](https://atcoder.jp/contests/abc177/tasks/abc177_c)

### 参考文献

- [「1000000007 で割ったあまり」の求め方を総特集！ 〜 逆元から離散対数まで 〜](https://qiita.com/drken/items/3b4fdf0a78e7a138cd9a)