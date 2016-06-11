# Kappa config

## Example

```yaml
backend: docker
backend_config:
  TLSCACert: "",
  TLSCert: "",
  TLSKey: "",
  AllowInsecure: false,
metrics:
  queue_length: "/usr/local/bin/queue_length.sh",
  reqsec:  "/usr/local/bin/reqsec.sh",
```

# Config Sections


## Metrics

```
metrics:
  queue_length: "/usr/local/bin/queue_length.sh",
  reqsec:  "/usr/local/bin/reqsec.sh",
```

## Backends

### Docker

```
backend: docker
backend_config:
  TLSCACert: "",
  TLSCert: "",
  TLSKey: "",
  AllowInsecure: false,
```

### Kubernetes

TODO: define backend_config

```
backend: kubernetes
backend_config:
  TO DEFINE
```


### Mesosphere

TODO: define backend_config

```
backend: mesosphere
backend_config:
  TO DEFINE
```

