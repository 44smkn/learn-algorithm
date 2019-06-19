# Greatest Common Divisor (GCD)

## description

[this article](https://www.geeksforgeeks.org/c-program-find-gcd-hcf-two-numbers/) helps understanding with you.

Greatest Common Divisor is, in Japanese, `最大公約数`.

There is two solutions.

- simple solution
  - find all prime factors of both numbers
    - prime factors are `素因数`
  - find intersection of all factors present in both numbers
    - intersection means `共通する数`
  - return product of elements in the intersection
  - O(n)
- efficient solution
  - use Euclidean algorithm (`ユークリッドの互除法`)
    - when x >= y, gcd(x, y) == gcd(y,x mod y)
    - O(logb)
  - Proof
    - if d is a common divisor of integer A and B, express `a = ld` and `b = md` using natural number `l` and `g`.
    - `a = bq + r`
      - `ld = mdq + r`
      - `r = (l-mq)d` indicates d is divisor of r.
