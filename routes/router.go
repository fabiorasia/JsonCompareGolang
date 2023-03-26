package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestRouter(route *gin.Engine) {
	testRouter := route.Group("/test")
	testRouter.POST("/test", postAlbums)
}

func postAlbums(c *gin.Context) {
	type jMap map[string]interface{}
	var inputJson jMap
	if err := c.BindJSON(&inputJson); err != nil {
		return
	}
	outputJson := make(jMap)
	subObjectMap := inputJson["subObject"].(map[string]interface{})
	var subObject jMap = subObjectMap
	outputJson["Integer"] = inputJson["int1"].(float64) +
		inputJson["int10"].(float64) +
		subObject["int5"].(float64) +
		subObject["int15"].(float64)
	outputJson["String"] = inputJson["string1"].(string) +
		inputJson["string10"].(string) +
		subObject["string5"].(string) +
		subObject["string15"].(string)
	outputJson["Boolean"] = inputJson["boolean1"].(bool) &&
		inputJson["boolean10"].(bool) &&
		subObject["boolean5"].(bool) &&
		subObject["boolean15"].(bool)

	c.IndentedJSON(http.StatusOK, outputJson)
}
