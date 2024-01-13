package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	port := 8084
	app := gin.Default()
	app.GET("/api/startup/test", welcomeHandler)
	app.Run(fmt.Sprintf(":%d", port))
	

}

func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Service is running")
}


func welcomeHandler(c *gin.Context){
	c.JSON(200, gin.H{
		"message":"Service is currently running",
	})


}
