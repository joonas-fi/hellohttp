[![Build Status](https://img.shields.io/travis/joonas-fi/hellohttp.svg?style=for-the-badge)](https://travis-ci.org/joonas-fi/hellohttp)
[![Download](https://img.shields.io/docker/pulls/joonas/hellohttp.svg?style=for-the-badge)](https://hub.docker.com/r/joonas/hellohttp/)

A hello world HTTP server with Prometheus metrics.

When deploying, specify (service-level) ENV variable `METRICS_ENDPOINT=/metrics` to
optionally expose metrics automatically with
[promswarmconnect](https://github.com/joonas-fi/hellohttp).


Example
-------

```
$ docker service create \
  --name hellohttp \
  --env "METRICS_ENDPOINT=/metrics" \
  --network yourNetwork \
  --publish 8012:80 \
  "joonas/hellohttp:VERSION"
```

Replace VERSION with a tag from Docker Hub link from top of this document.

You can now access http://localhost:8012/

Metrics are at http://localhost:8012/metrics (look for `requests_served`).

If you're using promswarmconnect, `requests_served` should now be queryable from Prometheus.
