package middleware

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LogToFile() gin.HandlerFunc {

	log := logs.NewLogger()
	err := log.SetLogger(logs.AdapterFile, `{"filename":"venus.log"}`)
	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		startAt := time.Now()

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		bodyStr := ""
		if c.Request.Body != nil && c.Request.Method == "POST" {
			bodybuf, err := c.GetRawData()
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
			}
			bodyStr = string(bodybuf)
		}

		userID := 0
		id, exist := c.Get("userID")
		if exist {
			userID = int(id.(float64))
		}

		log.Info(fmt.Sprintf("%s - [%s] \"%s %s %s [userID: %d] %d %s \"%s\" request_body:`%s` response_body:`%s` \n",
			c.ClientIP(),
			startAt.Format("2006-01-02 15:04:05"),
			c.Request.Method,
			c.Request.RequestURI,
			c.Request.Proto,
			userID,
			c.Writer.Status(),
			time.Since(startAt),
			c.Request.UserAgent(),
			bodyStr,
			blw.body.String(),
		))
	}
}

func LogToDB() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func LogToES() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func LogToConsole() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
