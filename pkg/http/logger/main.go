package logger

import (
	"io"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func SetLogger(out io.Writer, json bool, level string) {
	//	set time format
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// Log into JSON Format if bool is set to true
	if json {
		log.SetFormatter(&log.JSONFormatter{})
	}

	// redirect output to StdOut
	log.SetOutput(out)

	// Only log the warning severity or above.
	switch level {
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}

// logging middleware fucntion Gin
func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		st := time.Now()
		c.Next()
		et := time.Now()

		//build log messge
		log.WithFields(log.Fields{
			"host":        c.Request.Host,
			"status_code": c.Writer.Status(),
			"req_latency": et.Sub(st),
			"client_ip":   c.ClientIP(),
			"req_method":  c.Request.Method,
			"req_uri":     c.Request.RequestURI,
		}).Info()
	}
}
