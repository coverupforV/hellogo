# leetcode

## 1616

注意考虑check1里ab ba反过来的情况

从a头和b的尾开始，往中间靠，如果两个指针交错过则必能产生回文串

如果没有交错那么要判断a或者b剩下的i到j之间的部分是不是回文串

```Go
func checkPalindromeFormation(a string, b string) bool {
 return check1(a, b) || check1(b, a)
}

func check1(a, b string) bool {
 i, j := 0, len(b)-1
 for i < j && a[i] == b[j] {
  i++
  j--
 }
 return i >= j || check2(a, i, j) || check2(b, i, j)
}

func check2(a string, i, j int) bool {
 for i < j && a[i] == a[j] {
  i++
  j--
 }
 return i >= j
}
```

## 1625

会不了一点呜呜呜

```Go
func findLexSmallestString(s string, a int, b int) string {
 n := len(s)
 ans := s
 for _ = range s {
  s = s[n-b:] + s[:n-b]
  cs := []byte(s)
  for j := 0; j < 10; j++ {
   for k := 1; k < n; k += 2 {
    cs[k] = byte((int(cs[k]-'0')+a)%10 + '0')
   }
   if b&1 == 1 {
    for p := 0; p < 10; p++ {
     for k := 0; k < n; k += 2 {
      cs[k] = byte((int(cs[k]-'0')+a)%10 + '0')
     }
     s = string(cs)
     if ans > s {
      ans = s
     }
    }
   } else {
    s = string(cs)
    if ans > s {
     ans = s
    }
   }
  }
 }
 return ans
}
```
