package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

// Organization is the data type of organizations from data/Organizations_API_v1.json and IMPLEMENTS TreeNode
type Organization struct {
	ID               int    `json:"OrgID"`
	ImageURL         string `json:"OrgImageURL"`
	URLText          string `json:"OrgURLText"`
	URLAddress       string `json:"OrgURLAddress"`
	Type             string `json:"OrgType"`
	AbbrevName       string `json:"OrgAbbrevName"`
	Name             string `json:"OrgName"`
	JurisdictionType string `json:"OrgJurisdictionType"`
	ParentID         int    `json:"OrgParentID"`
	LastUpdatedDate  string `json:"LastUpdatedDate"`
}

// RecArea is the data type of recreation areas from data/RecAreas_API_v1.json and IMPLEMENTS TreeNode
type RecArea struct {
	OrgRecAreaID    string  `json:"OrgRecAreaID"`
	LastUpdatedDate string  `json:"LastUpdatedDate"`
	Email           string  `json:"RecAreaEmail"`
	ReservationURL  string  `json:"RecAreaReservationURL"`
	Longitude       float64 `json:"RecAreaLongitude"`
	ID              int     `json:"RecAreaID"`
	Phone           string  `json:"RecAreaPhone"`
	Description     string  `json:"RecAreaDescription"`
	MapURL          string  `json:"RecAreaMapURL"`
	Latitude        float64 `json:"RecAreaLatitude"`
	StayLimit       string  `json:"StayLimit"`
	FeeDescription  string  `json:"RecAreaFeeDescription"`
	Directions      string  `json:"RecAreaDirections"`
	Keywords        string  `json:"Keywords"`
	Name            string  `json:"RecAreaName"`
}

// Facility is the data type of facilities from data/Facilities_API_v1.json and IMPLEMENTS TreeNode
type Facility struct {
	Description       string  `json:"FacilityDescription"`
	Email             string  `json:"FacilityEmail"`
	Latitude          float64 `json:"FacilityLatitude"`
	UseFeeDescription string  `json:"FacilityUseFeeDescription"`
	LegacyID          float64 `json:"LegacyFacilityID"`
	OrgID             string  `json:"OrgFacilityID"`
	MapURL            string  `json:"FacilityMapURL"`
	Name              string  `json:"FacilityName"`
	LastUpdatedDate   string  `json:"LastUpdatedDate"`
	TypeDescription   string  `json:"FacilityTypeDescription"`
	ADAAccess         string  `json:"FacilityADAAccess"`
	Directions        string  `json:"FacilityDirections"`
	ID                int     `json:"FacilityID"`
	ReservationURL    string  `json:"FacilityReservationURL"`
	StayLimit         string  `json:"StayLimit"`
	Longitude         float64 `json:"FacilityLongitude"`
	Phone             string  `json:"FacilityPhone"`
	Keywords          string  `json:"Keywords"`
}

// TreeNode defines the methods for our n-ary tree of nodes
type TreeNode interface {
	GetID() string
	GetType() string
	GetParent() TreeNode
	SetParent(parent TreeNode)
	GetChildren() []TreeNode
	AddChild(c TreeNode)
	GetName()
}

func main() {
	data, err := ioutil.ReadDir("data")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range data {
		content, err := ioutil.ReadFile(fmt.Sprintf("data/%s", file.Name()))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s:\n%s\n\n\n\n\n", file.Name(), content)
	}

	// fmt.Printf("Organizations File: \n%s", organizations)
}
