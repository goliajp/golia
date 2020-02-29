package libgin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"time"
)

type HttpLog struct {
	Server        string        `json:"server"`
	Host          string        `json:"host"`
	ClientIP      string        `json:"client_ip"`
	RequestMethod string        `json:"request_method"`
	RequestBody   string        `json:"request_body"`
	RequestURI    string        `json:"request_uri"`
	ResponseCode  int           `json:"response_code"`
	ResponseBody  string        `json:"response_body"`
	Latency       time.Duration `json:"latency"`
}

func (l *HttpLog) Json() string {
	j, err := json.Marshal(l)
	if err != nil {
		log.Fatalln(err)
	}
	return string(j)
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LogMiddleware(svrName string, LogHandler func(l *HttpLog) error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqBody, _ := ctx.GetRawData()
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
		respBodyWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = respBodyWriter
		s := time.Now()
		ctx.Next()
		l := HttpLog{
			Server:        svrName,
			ClientIP:      ctx.ClientIP(),
			Host:          ctx.Request.Host,
			RequestMethod: ctx.Request.Method,
			RequestBody:   string(reqBody),
			RequestURI:    ctx.Request.RequestURI,
			ResponseCode:  ctx.Writer.Status(),
			ResponseBody:  respBodyWriter.body.String(),
			Latency:       time.Since(s),
		}
		err := LogHandler(&l)
		if err != nil {
			fmt.Println(err)
		}
	}
}
