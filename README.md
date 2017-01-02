# hotreload-example
Proof-of-concept for hot code reload in go. Currently only works on Linux using go1.8beta2!

## Usage

```
$ go get github.com/YuriyNasretdinov/hotreload-example
$ cd $GOPATH/src/github.com/YuriyNasretdinov/hotreload-example
$ hotreload-example &
$ curl 'http://localhost:9999/example?name=Yuriy'
Hello, Yuriy!
$ cd handlers/plug
$ go build -buildmode plugin -o ../../handlers.so
$ curl 'http://localhost:9999/reload'
OK
$ curl 'http://localhost:9999/example?name=Yuriy'
Hello modified, Yuriy!
```
