package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ginServe(){
	r := gin.Default()
	r.POST("/v1/item", func(context *gin.Context) {
		rsp := Client()
		if rsp != nil {
			context.JSON(http.StatusOK, rsp)
		}else{
			context.JSON(http.StatusInternalServerError, "")
		}
	})
	r.Run("localhost:8881")
}

func main(){
	var f string
	flag.StringVar(&f, "f", "client", "client or server ?")
	flag.Parse()

	if f == "client"{
		ginServe()
	}else{
		Serve()
	}
}
