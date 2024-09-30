package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://google.com"}

	var dao VoucherDao = &InMemoryVoucherDao{}

	r.POST("/add", func(c *gin.Context) {

		jsonData, err := io.ReadAll(c.Request.Body)

		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		entry := &VoucherEntry{}
		err = json.Unmarshal(jsonData, &entry)

		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		dao.CreateVoucherEntry(*entry)

		c.Status(http.StatusCreated)
	})

	r.GET("/list", func(c *gin.Context) {
		fmt.Println(c.ClientIP())
		entries, err := dao.ListAllVoucherEntry()
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		c.JSON(http.StatusOK, entries)
	})

	r.Use(cors.New(config))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
