###BET
- Test
- Benchmark
- Example

###Definition
- TestFun(t *testing.T)
- BenchmarkFun(b *testing.B)
- ExampleFun()

###Commands
```
go test            (test everything from here on subpackages, use ./...)
go test -cover     (check how much is total test coverage)
go test -coverprofile c.out
go tool cover -html=c.out

go test -bench .
```