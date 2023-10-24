package main

import (
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

// globally available
var router = gin.New()
var m = ginmetrics.GetMonitor()

var VERSION string = "1.1.2"
var LOG_LEVEL string
var NAME string
var DOC_ROOT string
var WAIT int = 0
var ADDRESS string
var V2_PATH string
var FAIL_FREQ int = 0
var FILL_RAM int = 0
