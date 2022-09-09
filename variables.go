package main

import "github.com/gin-gonic/gin"

// globally available
var router = gin.Default()

var VERSION string = "0.0.1"
var NAME string
var DOC_ROOT string
var WAIT int = 0
var ADDRESS string
