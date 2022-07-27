# Troll

Is a very simple webserver returning response.


## Purpouse

Testing API endpoint, responding with random predefined delay on all paths... 

### Features:
 * define minimal wait interval for response
 * define document root for serving static content
 * aplication name could be defined
 * listen port could be set via ENV
 * print received json data to log and respond in json
 * log with basic request_id
 
## RUN

```bash
troll -help

  -name string
        Define custom application name (default "troll")
  -root string
        Define document root for serving files (default "./public")
  -wait int
        Minimal wait time before each request

```

### Listen port

Application accept env PORT and the default value is **:8080**