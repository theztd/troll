# Troll

[![pipeline status](https://gitlab.com/theztd/troll/badges/main/pipeline.svg)](https://gitlab.com/theztd/troll/-/commits/main)   [![coverage report](https://gitlab.com/theztd/troll/badges/main/coverage.svg)](https://gitlab.com/theztd/troll/-/commits/main)   
[![Go Report Card](https://goreportcard.com/badge/gitlab.com/theztd/troll?style=flat-square)](https://goreportcard.com/report/gitlab.com/theztd/troll)   [![Latest Release](https://gitlab.com/theztd/troll/-/badges/release.svg)](https://gitlab.com/theztd/troll/-/releases)

Troll is a very simple webserver returning defined response with configurable delay and a few more features.

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
 
## RUN

```bash
troll -help

  -name string
        Define custom application name (default "troll")
  -root string
        Define document root for serving files (default "./public")
  -wait int
        Minimal wait time before each request
      
  -v2-path string
        Path to the yaml with custom api definition, example format is in part custom API definition

```

### Listen port

Application accept env PORT and the default value is **:8080**


### Custom API definition

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

