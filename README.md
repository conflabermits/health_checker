# Health Checker

It checks health!

## What is `health_checker`?

The `health_checker` tool's main goal is to make HTTP requests to web applications, bring back the response, parse it, and pretty-print it.

The web applications are expected to respond with JSON content, and `health_checker` is capable of parsing the response JSON and displaying it in pretty-printed JSON of various lengths (dicted by its `depth` option).

It also has a web server mode where it can start up a simple HTTP server that will host a simple health checker HTML page. The page has a text box to enter the URL and a button that will perform the check and return the response to the page.

## Usage Instructions

Help text:

```bash
$ go run main.go -help
Usage: health_checker-go [options]

  -depth string
        Determine amount/type of data to return (default "dynamic")
  -hostHeader string
        override Host specified in URL
  -port string
        Port to run the local web server (default "8080")
  -url string
        url to check
```

Depth flag to control amount of detail:

```bash
$ go run main.go -depth short -url "http://localhost:48080/outage"
{
    "name": "appname",
    "statusCode": "OUTAGE"
}

$ go run main.go -depth dynamic -url "http://localhost:48080/outage"
{
    "broken_components": [
        {
            "description": "Most important check",
            "essential": true,
            "name": "auth-service",
            "statusCode": "CRITICAL",
            "statusText": "Can't reach auth service, returns 500",
            "uri": "http://localhost:38080/auth-service/health"
        }
    ],
    "name": "appname",
    "statusCode": "OUTAGE"
}

$ go run main.go -depth full -url "http://localhost:48080/outage"
{
    "components": [
        {
            "description": "Most important check",
            "essential": true,
            "name": "auth-service",
            "statusCode": "CRITICAL",
            "statusText": "Can't reach auth service, returns 500",
            "uri": "http://localhost:38080/auth-service/health"
        },
        {
            "description": "Less important check",
            "essential": false,
            "name": "activity-webservice",
            "statusCode": "OK",
            "statusText": null,
            "uri": "http://localhost:38080/activity-service/health"
        },
        {
            "description": "Some other check",
            "essential": true,
            "name": "database",
            "statusCode": "OK",
            "statusText": null,
            "uri": "http://localhost:48080/user-table"
        }
    ],
    "name": "appname",
    "statusCode": "OUTAGE"
}
```

Port to run local web server:

Coming *soon*
