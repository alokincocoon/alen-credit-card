package router

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func isCardValid(cardNumber string) bool {
	sum := 0
	var valid bool = false
	for i := len(cardNumber) - 1; i >= 0; i-- {
		cardno, err := strconv.Atoi(string(cardNumber[i]))
		if err != nil {
			log.Println(err.Error())
			// return "error:" + err.Error()
			return valid
		}
		if i%2 == 0 {
			cardno *= 2
			if cardno > 9 {
				cardno -= 9
			}
		}
		sum += cardno
	}
	if sum%10 == 0 {
		valid = true
	}
	return valid
}

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/validate-card", func(c *gin.Context) {

		var data struct {
			CardNumber string `json:"cardNumber"`
		}

		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data.CardNumber = (strings.ReplaceAll(data.CardNumber, " ", ""))

		if len(data.CardNumber) != 16 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid card number"})
			return
		}

		var cardValid bool = isCardValid(data.CardNumber)
		if cardValid {
			c.JSON(http.StatusOK, gin.H{
				"message": "Valid card number",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid card number!",
			})
		}
	})

	return r
}
