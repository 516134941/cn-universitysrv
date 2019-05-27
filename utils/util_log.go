package utils

import (
	"net/http"
	"time"

	"github.com/go-xweb/log"
)

// LogStat 记录日志
func LogStat(logName string, r *http.Request, t1 time.Time) {
	log.Debugf("%v: requestURI:%v Post:%v cost:%v\n",
		logName, r.RequestURI, r.Form.Encode(), time.Since(t1))
}
