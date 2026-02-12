package config

import (
	"github.com/penglongli/gin-metrics/ginmetrics"
)

// globally available

var Metrics = ginmetrics.GetMonitor()

var VERSION string = "1.7.1"
var LOG_LEVEL string
var NAME string
var DOC_ROOT string
var REQUEST_DELAY int = 0
var ADDRESS string
var TCP_ADDRESS string
var TCP_DEST_ADDRESS string
var CONFIG_FILE string
var ERROR_RATE int = 0
var HEAVY_CPU int = 0
var HEAVY_RAM int = 0
var READY_DELAY int = 0
var GAME bool = false
var DSN string
var HOSTNAME string
var BADASS bool
var TCP_MIN_DELAY int
var TCP_MAX_DELAY int
var TCP_ERROR_RATE int
