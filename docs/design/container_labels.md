# Container Labels

- `kappa.min` : minimum number of containers that should be running
- `kappa.max` : minimum number of containers that should be running
- `kappa.rate` : refresh rate for the metric in seconds
- `kappa.metric` : the key for the metric to use as configured in the mapping


## `docker-compose.yml` example

```
app:
    image: nginx
    labels:
        - "kappa.min=1"
        - "kappa.max=10"
        - "kappa.max=5"
        - "kappa.metric=queue"
```
