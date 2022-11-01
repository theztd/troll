package main

import "github.com/gin-gonic/gin"

// globally available
var router = gin.Default()

var VERSION string = "1.0.0"
var NAME string
var DOC_ROOT string
var WAIT int = 0
var ADDRESS string
var V2_PATH string
var FAIL_FREQ int = 0
var FILL_RAM int = 0
