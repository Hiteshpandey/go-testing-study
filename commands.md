# Commands for testing

```
$ go test .
```
```
$ go test -v .
```
```
$ go test -cover .
```
```
$ go test -coverprofile=coverage.out .
```
```
$ go tool cover -html=coverage.out
```
```
$ go tool cover -html coverage.out -o cover.html
```