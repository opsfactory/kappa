# CONTRIBUTING

## Build using Docker

This will create a docker image named `kappa:latest` and a binary in `bin/kappa`

```
make build
```

## Build manually

```
go build .
```


## Adding a dependency

Kappa is using [godep](https://github.com/tools/godep) to manage dependencies.

If you really need to add a dependency to kappa this is the way to go:

- Run `go get foo/bar`
- Edit your code to import foo/bar.
- Run godep save ./...
