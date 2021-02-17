package router

import (
	"BackendGo/auth"
	"BackendGo/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

func RequestApi(c *gin.Context, query string) {
	response, err := http.Get(query)
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	reader := response.Body
	defer reader.Close()
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")
	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, nil)
}

func queries(s *server.Server) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := auth.ExtractTokenID(c.Request)
		if err != nil {
			_ = c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		user, err := s.Controller.GetUserByID(id)
		if err != nil {
			_ = c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		datas := c.Query("datas")

		fmtUrl := fmt.Sprintf("http://%s/api/v1/query", s.PrometheusURL)
		u, _ := url.Parse(fmtUrl)
		values, _ := url.ParseQuery(u.RawQuery)
		str := fmt.Sprintf("%s{owner=\"%s\"}", datas, user.UUID.String())
		values.Set("query", str)
		u.RawQuery = values.Encode()
		RequestApi(c, fmt.Sprintf("%v", u))
	}
}

func queryExporter(s *server.Server) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := auth.ExtractTokenID(c.Request)
		if err != nil {
			_ = c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		user, err := s.Controller.GetUserByID(id)
		if err != nil {
			_ = c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		datas := c.Query("datas")
		exporter := c.Param("exporter")

		userExporters, err := s.Controller.GetExportersFromUser(user)
		if err != nil {
			_ = c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		var foundExporter string
		for _, userExporter := range userExporters {
			foundExporter = userExporter.UUID.String()
			if userExporter.UUID.String() == exporter {
				break
			}
			foundExporter = ""
		}
		if foundExporter == "" {
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		fmtUrl := fmt.Sprintf("http://%s/api/v1/query", s.PrometheusURL)
		u, _ := url.Parse(fmtUrl)
		values, _ := url.ParseQuery(u.RawQuery)
		str := fmt.Sprintf("%s{exporter=\"%s\"}", datas, foundExporter)
		values.Set("query", str)
		u.RawQuery = values.Encode()
		RequestApi(c, fmt.Sprintf("%v", u))
	}
}
