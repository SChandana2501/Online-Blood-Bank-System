package model

import (
	_ "go.mongodb.org/mongo-driver/bson"
)

type Donor struct {
	FirstName                 string `bson:"firstName" json:"firstName"`
	LastName                  string `bson:"lastName" json:"lastName"`
	Age                       string `bson:"age" json:"age"`
	Email                     string `bson:"email" json:"email"`
	Phone                     string `bson:"phone" json:"phone"`
	Address                   string `bson:"address" json:"address"`
	City                      string `bson:"city" json:"city"`
	State                     string `bson:"state" json:"state"`
	ZipCode                   string `bson:"zipCode" json:"zipCode"`
	BloodGroup                string `bson:"bloodGroup" json:"bloodGroup"`
	Occupation                string `bson:"occupation" json:"occupation"`
	PreviousBloodDonation     string `bson:"previousBloodDonation" json:"previousBloodDonation"`
	PreviousBloodDonationDate string `bson:"previousBloodDonationDate" json:"previousBloodDonationDate"`
	Disease                   string `bson:"disease" json:"disease"`
	DiseaseName               string `bson:"diseaseName" json:"diseaseName"`
	Allergy                   string `bson:"allergy" json:"allergy"`
	AllergyName               string `bson:"allergyName" json:"allergyName"`
	DateAdded                 string `bson:"dateAdded" json:"dateAdded"`
}
