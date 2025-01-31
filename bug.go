```go
func main() {
    var m map[string]int
    fmt.Println(m["key"]) // This will not panic, but will print 0.
}
```