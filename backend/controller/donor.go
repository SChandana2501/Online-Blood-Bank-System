package controller

import (
	"bloodbanksystem/db"
	"bloodbanksystem/model"
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func SaveDonor(c *gin.Context) {
	var data model.Donor
	err := c.BindJSON(&data)
	if err != nil {
		color.Red("Error in unmarshalling!! " + err.Error())
		c.JSON(http.StatusBadGateway, gin.H{"status": "error"})
		return
	}
	fmt.Println(data)
	if data.FirstName == "" || data.LastName == "" || data.Email == "" || data.Phone == "" || data.Address == "" || data.City == "" || data.State == "" || data.ZipCode == "" || data.BloodGroup == "" {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "required fields missing"})
		return
	}
	data.State = strings.ToLower(data.State)
	data.City = strings.ToLower(data.City)
	db, err := db.Dbconnect()
	if err != nil {
		color.Red("db connection error!!! " + err.Error())
		c.JSON(http.StatusBadGateway, gin.H{"status": "error"})
		return
	}
	defer db.Disconnect(c)
	data.DateAdded = time.Now().Format("02 Jan 06")
	collection := db.Database("bloodbank").Collection("donors")
	_, err = collection.InsertOne(context.TODO(), data)
	if err != nil {
		color.Red("insert donor error!!! " + err.Error())
		c.JSON(http.StatusBadGateway, gin.H{"status": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success","message":"Data Saved"})
}

func GetDonors(c *gin.Context) {
	filter := make(map[string]interface{})
	bloodgroup := c.Query("bg")

	city := c.Query("city")
	state := c.Query("state")
	zipCode := c.Query("zipCode")
	if bloodgroup != "" {
		filter["bloodGroup"] = bloodgroup
	}
	if city != "" {
		filter["city"] = strings.ToLower(city)
	}
	if state != "" {
		filter["state"] = strings.ToLower(state)
	}
	if zipCode != "" {
		filter["zipCode"] = zipCode
	}
	fmt.Println(filter)
	var donors []model.Donor
	db, err := db.Dbconnect()
	if err != nil {
		color.Red("db connection error!!! " + err.Error())
		c.JSON(http.StatusBadGateway, gin.H{"status": "error in connecting to DB"})
		return
	}
	defer db.Disconnect(c)
	collection := db.Database("bloodbank").Collection("donors")
	cursor, err := collection.Find(c, filter)
	if err != nil {
		color.Red("getting donors error!!! " + err.Error())
		c.JSON(http.StatusBadGateway, gin.H{"status": "error in getting donors"})
		return
	}
	for cursor.Next(c) {
		var donor model.Donor
		err := cursor.Decode(&donor)
		if err != nil {
			color.Red("decoding donors error!!! " + err.Error())
			c.JSON(http.StatusBadGateway, gin.H{"status": "error in decoding donors"})
			return
		}
		donors = append(donors, donor)
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "donors": donors})
}
