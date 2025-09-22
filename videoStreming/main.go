package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)


func main(){
	r:=gin.Default()
	r.GET("/stream/:video",func(ctx *gin.Context) {
		filename:=ctx.Param("video")
		filepath:=filepath.Join("./",filename)

		file,err:=os.Open(filepath)
		if err!=nil{
			ctx.JSON(http.StatusInternalServerError,"Failed to open file")
			return
		}
		defer file.Close()

		fi,err:=file.Stat()
		if err!=nil{
			ctx.JSON(http.StatusInternalServerError,"Error while info vide")
			return
		}
		ctx.Header("content-type","video/mp4")
		ctx.Header("content-Disposition","inline")
		ctx.Header("accept-Ranges","bytes")
		http.ServeContent(ctx.Writer,ctx.Request,filename,fi.ModTime(),file)
	})

	r.GET("/download/:video",func(ctx *gin.Context) {
		filename:=ctx.Param("video")
		filepath:=filepath.Join("./",filename)

		file,err:=os.Open(filepath)
		if err!=nil{
			ctx.JSON(http.StatusInternalServerError,"Failed to open file")
			return
		}
		defer file.Close()

		fi,err:=file.Stat()
		if err!=nil{
			ctx.JSON(http.StatusInternalServerError,"Error while info vide")
			return
		}
		ctx.Header("content-type","applocation/octet-stream")
		ctx.Header("content-Disposition","attachment;filename="+filename)
		ctx.Header("content-length",string(fi.Size()))
		http.ServeContent(ctx.Writer,ctx.Request,filename,fi.ModTime(),file)
	})

	r.Run(":8050")
}