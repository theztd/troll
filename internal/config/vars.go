package config

import (
	"github.com/penglongli/gin-metrics/ginmetrics"
)

// globally available

var Metrics = ginmetrics.GetMonitor()

var VERSION string = "1.2.0"
var LOG_LEVEL string
var NAME string
var DOC_ROOT string
var WAIT int = 0
var ADDRESS string
var V2_PATH string
var FAIL_FREQ int = 0
var FILL_RAM int = 0
var READY_DELAY int = 0
