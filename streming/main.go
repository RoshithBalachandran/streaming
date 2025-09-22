package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/stream/:audio", func(ctx *gin.Context) {
		filename := ctx.Param("audio")
		filepath := filepath.Join("./", filename)

		file, err := os.Open(filepath)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Failed to open file")
			return
		}

		defer file.Close()

		fi, err := file.Stat()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "Error get file info")
			return
		}
		ctx.Header("content-type", "audio/mpeg")
		ctx.Header("content-DisPosition", "inline")
		ctx.Header("Accept-Ranges", "bytes")

		http.ServeContent(ctx.Writer, ctx.Request, filename, fi.ModTime(), file)
	})
	r.GET("/download/:audio", func(ctx *gin.Context) {
		filename := ctx.Param("audio")
		filepath := filepath.Join("./", filename)

		file, err := os.Open(filepath)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Failed to open file")
			return
		}

		defer file.Close()

		fi, err := file.Stat()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "Error get file info")
			return
		}
		ctx.Header("content-type", "application/octet-stream")
		ctx.Header("content-DisPosition", "attachment;filename="+filename)
		ctx.Header("content-length", string(fi.Size()))

		http.ServeContent(ctx.Writer, ctx.Request, filename, fi.ModTime(), file)
	})

	r.Run(":8050")
}
