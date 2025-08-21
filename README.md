# Troll

  [![Troll Release](https://github.com/theztd/troll/actions/workflows/release_golang.yml/badge.svg)](https://github.com/theztd/troll/actions/workflows/release_golang.yml)
  [![Helm Release](https://github.com/theztd/troll/actions/workflows/release_helm.yml/badge.svg)](https://github.com/theztd/troll/actions/workflows/release_helm.yml)
  [![Docker image Release](https://github.com/theztd/troll/actions/workflows/release_docker.yml/badge.svg)](https://github.com/theztd/troll/actions/workflows/release_docker.yml)


Troll is a very simple webserver returning defined response with configurable delay and a few more features.

**INFO**

 * Report Issues: **https://github.com/theztd/troll/issues/new**
 * Listen port: **8080** (changable via **PORT** env)
 * Health Url: **/_healthz/ready.json**
 * Readines URL: **/_healthz/ready.json**
 * Livenes URL: **/_healthz/ready.json**
 * Metrics URL **/_healthz/metrics/**

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

## 🧱 Build

```bash
env GOOS=target-OS GOARCH=target-architecture go build -o troll cmd/troll/main.go
```

## 🚀 RUN and Operate

```bash
troll -help
  -addr string
        Define address and port where the application listen. (ADDRESS) (default ":8080")
  -config string
        Configure api endpoint. (CONFIG_FILE) (default "./config.yaml")
  -dsn string
        Define database DSN (default "postgresql://chuvicka:XciF3j5tLlMPVxlqBWlzjg@my-lab-3925.8nj.gcp-europe-west1.cockroachlabs.cloud:26257/chuvicka?sslmode=verify-full")
  -fail int
        Returns 503. Set 1 - 10, where 10 = 100% error rate. (FAIL_FREQ)
  -fill-cpu int
        Generate stress on CPU with each request. It also works as a delay for request. Set in milisecodns. (HEAVY_CPU)
  -fill-ram int
        Fill ram with each request. Set number in bytes. (HEAVY_RAM)
  -log string
        Define LOG_LEVEL (default "info")
  -name string
        Define custom application name. (NAME) (default "troll")
  -ready-delay int
        Simulate long application init [sec]. (READY_DELAY) (default 5)
  -root string
        Define document root for serving files. (DOC_ROOT) (default "./public")
  -wait int
        Minimal wait time before each request. (REQUEST_DELAY)

```

### 📦 Helm install

```bash
helm repo add troll https://theztd.github.io/troll/
helm install troll-test  troll/troll
```

### 🛠️ Config

Configuration is possible via ENV variables

 |Key| Default | description |
 |---|:---:|---|
 | ADDRESS | :8080 | Set **127.0.0.1:8080** to listen only on localhost |
 | LOG_LEVEL | info | Valid options are: error, warning, info, debug |

More options are posible via arguments. Configuration of v2 api is possible as follows... 

### 🤖 Custom API definition

By editing v2_api.yaml you can change /v2 endpoints and his responses (return code including).

The default structure is:
```yaml
---
name: Inventory
description: Our company inventory includes employees and equipment
version: 2022-02-09
endpoints:
- path: /employee
  method: GET
  code: 200
  response: "List of our employee..."
```

### 🌿 Dependencies

 | Name | Url | Notes |
 |---|---|---|
 | FS:  | ./v2_api.yaml | v2 config file have to be available. |
 

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
2023/04/18 16:41:11 [application:nodeName] INFO: Running env devel
2023/04/18 16:41:12 [application:nodeName] DEBUG: Metrics DB migration...
2023/04/18 16:41:12 [application:nodeName] DEBUG: Auth DB migration...
2023/04/18 16:41:12 [application:nodeName] INFO: Waiting for request at address :8080
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


