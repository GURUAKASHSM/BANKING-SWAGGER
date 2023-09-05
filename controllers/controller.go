package controllers

import (
	"fmt"
	"mongoapi/models"
	"mongoapi/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Getalldata(c *gin.Context) {
	alltransaction := service.Getalldata()
	c.JSON(http.StatusOK, alltransaction)
}

func Getalldataformpost(c *gin.Context) {
	var requestBody struct {
		From string `json:"from"`
		To   string `json:"to"`
	}
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	fromDate, err := time.Parse("2006-01-02", requestBody.From)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'from' date format"})
		return
	}
	toDate, err := time.Parse("2006-01-02", requestBody.To)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'to' date format"})
		return
	}
	allmovies := service.Getalldatabydate(fromDate, toDate)
	c.JSON(http.StatusOK, allmovies)
}

func Getsumformpost(c *gin.Context) {
	type SumResponse struct {
		TotalAmount interface{} `json:"totalAmount"`
	}
	var requestBody struct {
		From string `json:"from"`
		To   string `json:"to"`
	}
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	fromDate, err := time.Parse("2006-01-02", requestBody.From)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'from' date format"})
		return
	}
	toDate, err := time.Parse("2006-01-02", requestBody.To)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'to' date format"})
		return
	}
	sum, _ := service.Getsumbydate(fromDate, toDate)
	sumResponse := SumResponse{
		TotalAmount: sum,
	}
	c.JSON(http.StatusOK, sumResponse)
}

func Getdatabyid(c *gin.Context) {
	var id struct{
		ID string `json:"id" bson:"id"`
	}
	if err := c.BindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	

	data := service.Getdatabyid(id.ID)
	c.JSON(http.StatusOK, data)
}

func Depositcontroll(c *gin.Context) {
	var requestBody struct {
		Amount string `json:"amount"`
		FromId string `json:"fromid"`
		ToId   string `json:"toid"`
	}
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	amount, err := strconv.ParseFloat(requestBody.Amount, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'from' date format"})
		return
	}

	fmt.Println("con")
	str := service.Deposit(amount, requestBody.FromId, requestBody.ToId)
	c.JSON(http.StatusOK, str)
}

func CreateProfile(c *gin.Context) {
	var profile models.Profile
	if err := c.BindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	service.Insert(profile)
	c.JSON(http.StatusOK, profile)
}

func UpdateProfile(c *gin.Context) {
	id := c.Param("id")
	service.UpdateOne(id)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Profile %s updated", id)})
}

func Deleteprofile(c *gin.Context) {
	id := c.Param("id")
	service.DeleteOne(id)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Profile %s deleted", id)})
}

func Deleteallprofile(c *gin.Context) {
	count := service.DeleteAll()
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Deleted %d profiles", count)})
}
