# glimit
限制 go 协程并发数

# 示例

```go
g := New(3)
for i := 0; i < 15; i++ {
    g.Add()
    go func(g *Limit, i int) {
        fmt.Println(i)
        time.Sleep(2 * time.Second)
        g.Done()
    }(g, i)
}
g.Wait()
```