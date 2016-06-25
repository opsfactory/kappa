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

Kappa is using [gvt](https://github.com/FiloSottile/gvt) to manage dependencies.

If you really need to add a dependency to kappa this is the way to go:

- Run `go get -u github.com/FiloSottile/gvt`
- Edit your code to import `foo/bar`.
- Run `gvt get foo/bar`
