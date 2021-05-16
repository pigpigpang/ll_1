package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ll1_analyse/ll1"
)

type Data struct {
	Input string `json:"input"`
	Temp  string `json:"temp"`
}

func main() {
	r := gin.Default()
	r.POST("/analyse", func(c *gin.Context) {

		var parsedBody Data
		if err := c.BindJSON(&parsedBody); err != nil {
			fmt.Println(err)
			fmt.Println(parsedBody)
			c.JSON(400, gin.H{"message": "参数错误 非可识别Json"})
			return
		}
		input := parsedBody.Input
		temp := parsedBody.Temp
		firstmap, followmap, table, b, yesset, noset,ans := ll1.Analyse(input,temp)
		if b {
			c.JSON(200, gin.H{
				"message": "不是ll(1)文法，有冲突",
				"first":   firstmap,
				"follow":  followmap,
				"table":   table,
				"yesset":  yesset,
				"noset":   noset,
				"ans": ans,
			})
		} else {
			c.JSON(200, gin.H{
				"message": "是ll(1)文法",
				"first":   firstmap,
				"follow":  followmap,
				"table":   table,
				"yesset":  yesset,
				"noset":   noset,
				"ans": ans,
			})
		}

	})
	r.Run(":5433") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
