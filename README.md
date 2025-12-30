# Troll

  [![Troll Release](https://github.com/theztd/troll/actions/workflows/release_golang.yml/badge.svg)](https://github.com/theztd/troll/actions/workflows/release_golang.yml)
  [![Helm Release](https://github.com/theztd/troll/actions/workflows/release_helm.yml/badge.svg)](https://github.com/theztd/troll/actions/workflows/release_helm.yml)
  [![Docker image Release](https://github.com/theztd/troll/actions/workflows/release_docker.yml/badge.svg)](https://github.com/theztd/troll/actions/workflows/release_docker.yml)


Troll is a very simple webserver returning defined response with configurable delay and a few more features.

**INFO**

 * Report Issues: **https://github.com/theztd/troll/issues/new**
 * Http listen port: **8080** (changable via **http.ADDR** env)
 * TCP proxy listen port: **9999** (changable via **tcp.ADDR** env)
 * Info endpoint URL: **GET /_healthz/info**
 * Readines URL: **GET /_healthz/ready**
 * Livenes URL: **GET /_healthz/alive**
 * Metrics URL: **GET /_healthz/metrics**
 * Headers debug URL: **GET /v1/headers**

**Quick Links**

 * [DEMO application (Try it!!!)](https://troll.check-this.link)



## Purpouse
 * Testing API endpoint (configurable via YAML).
 * Responding with random predefined delay.
 * Serve static files from FS.
 * Ready to use API backend for FE app prototyping.
 * Demo CI/CD with GOlang application

## Features:
 * Define minimal wait interval for response
 * Define document root for serving static content
 * Aplication name could be defined
 * Listen port could be set via ENV
 * Print received json data to log and respond in json
 * Log with basic request_id
 * 404 page dumping request to log
 * Generate 503 randomly (simulate errors)
 * Fill RAM with each request (simulate mem leaks)
 * Generate CPU load on requests with ?heavy=cpu param
 * Ready delay for testing canary releases and readyness check
 * Tcp Proxy with random delay and random error generator
 * Routes could be defined via YAML config and can do SQL query, redis query and shell commands

## 🧱 Build

```bash
env GOOS=target-OS GOARCH=target-architecture go build -o troll cmd/troll/main.go
```

## 🚀 RUN and Operate


### 🛠️ Config
```bash
troll -help
  -name string
        Define custom application name. (NAME) (default "troll")
  -log_level string
        Define LOG_LEVEL (default "info")
  -ready_delay int
        Simulate long application init [sec]. (READY_DELAY) (default 5)
  -config string
        Configure api endpoint. (CONFIG_FILE)

  -http.addr string
        Define address and port where the application listen. (HTTP_ADDR) (default ":8080")
  -http.error_rate int
        Returns 503. Set 1 - 10, where 10 = 100% error rate. (HTTP_ERROR_RATE)
  -http.fill_cpu int
        Generate stress on CPU with each request. It also works as a delay for request [milisecodns]. (HEAVY_CPU)
  -http.fill_ram int
        Fill ram with each request [bytes]. (HEAVY_RAM)
  -http.req_delay int
        Minimal delay before response on request [miliseconds]. (REQUEST_DELAY)
  -http.root string
        Define document root for serving files. (DOC_ROOT) (default "./public")

  -tcp.addr string
        Define address and port where the tcp proxy listens. (TCP_ADDR) (default ":9999")
  -tcp.dest_addr string
        Define address and port where to send tcp proxy requests. (TCP_DEST_ADDR) (default "127.0.0.1:8080")
  -tcp.error_rate int
        Simulate random error rate.  Set 1 - 10, where 10 = 100% error rate. (TCP_ERROR_RATE)
  -tcp.max_delay int
        Simulate long response max delay [miliseconds]. (TCP_MAX_DELAY) (default 5000)
  -tcp.min_delay int
        Simulate long response minimal delay [miliseconds]. (TCP_MIN_DELAY) (default 100)

```

### 📦 Helm install

```bash
helm repo add troll https://theztd.github.io/troll/
helm install troll-test  troll/troll
```


### 🤖 Custom API definition

By editing config_api.yaml you can change /v2 endpoints and his responses (return code including).

The default structure is:
```yaml
---
name: Inventory
description: Our company inventory includes employees and equipment
version: 2022-09-09
game:
  # Enable and configure GAME UI
  route: /game
  templatePath: ./__game_template.html
  backends:
  - http://service-a
  - http://service-b
  - http://service-c
endpoints:
- path: /machines
  kind: basic
  method: GET
  code: 200
  response: "We have plenty of machines in our factory as you can see as follows..."
- path: /machines/add
  kind: basic
  method: POST
  code: 200
  response: "New machine has been added"
```

### 🌿 Dependencies

 | Name | Url | Notes |
 |---|---|---|
 | FS:  | ./config.yaml | Not required. |
 

### 📈 Monitoring

#### Health check example
 ```json
# Content-Type: "application/health+json"
{
      "notes":"Troll is ...",
      "status":"pass",
      "version":"1.0.0"
}
```

#### Metrics example
```prometheus
# HELP gin_request_duration the time server took to handle the request.
# TYPE gin_request_duration histogram
gin_request_duration_bucket{uri="",le="0.1"} 300
gin_request_duration_bucket{uri="",le="0.3"} 300
gin_request_duration_bucket{uri="",le="1.2"} 300
gin_request_duration_bucket{uri="",le="5"} 300
gin_request_duration_bucket{uri="",le="10"} 300
gin_request_duration_bucket{uri="",le="+Inf"} 300
gin_request_duration_sum{uri=""} 0.41419708799999966
gin_request_duration_count{uri=""} 300
gin_request_duration_bucket{uri="/v1/:item/*id",le="0.1"} 61
gin_request_duration_bucket{uri="/v1/:item/*id",le="0.3"} 184
gin_request_duration_bucket{uri="/v1/:item/*id",le="1.2"} 300
gin_request_duration_bucket{uri="/v1/:item/*id",le="5"} 300
gin_request_duration_bucket{uri="/v1/:item/*id",le="10"} 300
gin_request_duration_bucket{uri="/v1/:item/*id",le="+Inf"} 300
gin_request_duration_sum{uri="/v1/:item/*id"} 72.53894113400001
gin_request_duration_count{uri="/v1/:item/*id"} 300
gin_request_duration_bucket{uri="/v1/status",le="0.1"} 100
gin_request_duration_bucket{uri="/v1/status",le="0.3"} 100
gin_request_duration_bucket{uri="/v1/status",le="1.2"} 100
gin_request_duration_bucket{uri="/v1/status",le="5"} 100
gin_request_duration_bucket{uri="/v1/status",le="10"} 100
gin_request_duration_bucket{uri="/v1/status",le="+Inf"} 100
gin_request_duration_sum{uri="/v1/status"} 0.003233080999999999
gin_request_duration_count{uri="/v1/status"} 100
# HELP gin_request_total all the server received request num.
# TYPE gin_request_total counter
gin_request_total 700
# HELP gin_request_uv_total all the server received ip num.
# TYPE gin_request_uv_total counter
gin_request_uv_total 1
# HELP gin_response_body_total the server send response body size, unit byte
# TYPE gin_response_body_total counter
gin_response_body_total 322135
# HELP gin_uri_request_total all the server received request num with every uri.
# TYPE gin_uri_request_total counter
gin_uri_request_total{code="200",method="GET",uri="/v1/:item/*id"} 300
gin_uri_request_total{code="200",method="GET",uri="/v1/status"} 100
gin_uri_request_total{code="404",method="GET",uri=""} 300
# HELP go_sync_mutex_wait_total_seconds_total Approximate cumulative time goroutines have spent blocked on a sync.Mutex or sync.RWMutex. This metric is useful for identifying global changes in lock contention. Collect a mutex or block profile using the runtime/pprof package for more detailed contention data.
# TYPE go_sync_mutex_wait_total_seconds_total counter
go_sync_mutex_wait_total_seconds_total 0.000130336
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 10
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 2
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
```

### 🔎 Logs

Change log level by env LOG_LEVEL

**Format**
```log
2025/12/30 16:23:03 INFO: Loading configuration from .env file.
Starting application, give me 1 sec.. DONE

2025/12/30 16:23:04 INFO TcpProxy: Listening on :9999 and serving content from 127.0.0.1:8080
2025/12/30 16:23:04 WARN: Config file is not defined, but continue with defaults..
2025/12/30 16:23:04 INFO: Initialize default routes 🏗️  ...

Available routes:
  ▶︎ GET    /
  ▶︎ GET    /_healthz/info
  ▶︎ GET    /_healthz/alive
  ▶︎ GET    /_healthz/ready
  ▶︎ GET    /_healthz/metrics
  ▶︎ GET    /v1/_healthz/metrics
  ▶︎ GET    /v1/headers
  ▶︎ GET    /v1/:item/*id
  ▶︎ GET    /ws
  ▶︎ GET    /websocket
  ▶︎ GET    /public/*filepath
  ▶︎ GET    /static/*filepath
  ▶︎ HEAD   /public/*filepath
  ▶︎ HEAD   /static/*filepath
  ▶︎ POST   /v1/:item/*id


2025/12/30 16:23:04 INFO: Running in mode: "info" and listening on address :8080. 😈 Enjoy!
2025/12/30 16:23:14 INFO [AuditLog]: GET   /_metrics      404   From: 127.0.0.1 UA: vm_promscrape
2025/12/30 16:23:24 INFO [AuditLog]: GET   /_metrics      404   From: 127.0.0.1 UA: vm_promscrape

```

### Backup

* There is nothing to backup


## Contribute

Thank you for your interest in this project! Everyone is welcome to help.

### ✅ Commit rules
We use a simple commit style: type(scope): short message

Here are the types you can use:

| Type     | Description                        | Example Commit Message                         |
|----------|------------------------------------|------------------------------------------------|
| **feat** | New feature                        | feat(api): add /v2/status route              |
| **fix**  | Small fix or bug fix               | fix(handler): check for nil input            |
| **ref**  | Refactor (clean up or improve)     | ref(server): simplify router setup           |
| **doc**  | Documentation                      | doc(readme): add build instructions          |
| **ci**   | CI/CD or pipeline changes          | ci(k8s): append canary release ingress       |
| **test** | Testing or temporary experiments   | test(api): log request body for debugging    |

### 🎭 Helm release

 * Do changes in ./Helm directory
 * Don't forget to increase version in Chart.yaml file
 * Push your changes
