package server

import (
	"embed"
	"github.com/genleel/go-sync/server/controller"
	"github.com/gin-gonic/gin"
	"io/fs"
	"log"
	"net/http"
	"strings"
)

//go:embed frontend/dist/*
var FS embed.FS

func Run() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	staticFiles, _ := fs.Sub(FS, "frontend/dist")
	r.GET("/uploads/:path", controller.UploadsController)
	r.GET("/api/v1/addresses", controller.AddressController)
	r.GET("/api/v1/qrcodes", controller.QrcodeController)
	r.POST("/api/v1/texts", controller.TextController)
	r.POST("/api/v1/files", controller.FileController)
	r.StaticFS("/static", http.FS(staticFiles))
	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		if strings.HasPrefix(path, "/static/") {
			reader, err := staticFiles.Open("index.html")
			if err != nil {
				log.Fatal(err)
			}
			defer reader.Close()
			stat, err := reader.Stat()
			if err != nil {
				log.Fatal(err)
			}
			c.DataFromReader(http.StatusOK, stat.Size(), "text/html;charset=utf-8", reader, nil)
		} else {
			c.Status(http.StatusNotFound)
		}
	})
	r.Run(":27149")

}
