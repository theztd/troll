package config

import (
	"github.com/penglongli/gin-metrics/ginmetrics"
)

// globally available

var Metrics = ginmetrics.GetMonitor()

var VERSION string = "1.6.2"
var LOG_LEVEL string
var NAME string
var DOC_ROOT string
var WAIT int = 0
var ADDRESS string
var CONFIG_FILE string
var FAIL_FREQ int = 0
var HEAVY_CPU int = 0
var HEAVY_RAM int = 0
var READY_DELAY int = 0
var GAME bool = false
var DSN string
var HOSTNAME string
