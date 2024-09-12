package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func IndexController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "pong",
	})
}

func CallJavaDemo(c *gin.Context) {
	getJava(1)
	getJava(2)
	getJava(3)
	getJava(4)
	getJava(5)
}

func getJava(num int) {
	domain := os.Getenv("JAVA_DEMO_DOMAIN")
	resp, err := http.Get(fmt.Sprintf("http://%s/test?count=%d", domain, num))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Sprintf("%s, %s", resp.Body, "1234456563")
}
