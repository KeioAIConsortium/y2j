a tool that converts yaml to json

```
$ ./y2j foo.yaml # or cat foo.yaml | ./y2j
```

`env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o y2j` to build
