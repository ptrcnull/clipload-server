package main

import (
	"clipload-server/rand"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"path"
)

var BaseUrl string

func main() {
	r := gin.Default()

	r.GET("/*.png", func(c *gin.Context) {
		c.File(path.Join("img", c.Request.URL.Path[1:]))
	})

	r.POST("/", func(c *gin.Context) {
		data, err := c.GetRawData()
		if err != nil {
			c.AbortWithStatusJSON(400, err)
			return
		}

		filename := rand.RandString(16) + ".png"

		err = ioutil.WriteFile(path.Join("img", filename), data, 0644)
		if err != nil {
			c.AbortWithStatusJSON(500, err)
			return
		}

		c.String(200, path.Join(BaseUrl, filename))
	})

	err := r.Run("0.0.0.0:8012")
	if err != nil {
		log.Panicln(err)
		return
	}
}
