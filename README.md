# Go Tutorial (Tech with Lana)

https://www.youtube.com/watch?v=yyUHQIec83I&list=WL&index=3

https://pkg.go.dev/

```
go mod init booking-app
# go run main.go
# go run main.go helper.go
go run .
```

Scope and packages
* packages your create are all related to the same module <br>
(import using booking-app.<package> or <package> booking-app.<folder>)
* export variables and functions by upper-casing the first letter
* package scope : at the root of the package, should use full syntax var/const (:= not allowed)

Async management
* create go routines by adding operator `go` before function call
* Synchronize between routines using `sync.WaitGroup` `Add`/`Wait`/`Done`