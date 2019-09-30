a tool that converts yaml to json 
```
$ ./y2j foo.yaml # or ./y2j < foo.yaml
```

`env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o y2j` to build

static binary available in [releases](https://github.com/mt-caret/y2j/releases)
