package http

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func RequestWithBody(method string, path string, body []byte, token string, g *gin.Engine) (*httptest.ResponseRecorder, error) {

	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)

	httpReq, err := http.NewRequest(method, path, bytes.NewBuffer(body))
	if token != "" {
		httpReq.Header.Add("x-token", token)
	}

	c.Request = httpReq

	g.HandleContext(c)

	return rr, err
}
