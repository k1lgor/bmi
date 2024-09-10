package web

import (
	"net/http"
	"strconv"

	"bmi/pkg/bmi"
	"bmi/pkg/input"

	"github.com/gin-gonic/gin"
)

func StartWeb() {
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"bmi": "", "error": ""})
	})

	router.POST("/calculate", func(c *gin.Context) {
		weightStr := c.PostForm("weight")
		heightStr := c.PostForm("height")
		weightUnit := c.PostForm("weightUnit")
		heightUnit := c.PostForm("heightUnit")

		weight, height, err := input.ParseInput(weightStr, heightStr, weightUnit, heightUnit)
		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{"bmi": "", "error": err.Error()})
			return
		}

		bmiValue := bmi.CalculateBMI(weight, height)
		bmiResult := bmi.InterpretBMI(bmiValue)
		c.HTML(http.StatusOK, "index.html", gin.H{"bmi": strconv.FormatFloat(bmiValue, 'f', 2, 64), "category": bmiResult, "error": ""})
	})

	router.Run(":8080")
}
