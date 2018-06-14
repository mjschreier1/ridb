package main

import (
	"encoding/json"
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
	OrgRecAreaID    float64 `json:"OrgRecAreaID"`
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

// RecAreaFacilityLink is the data type of data/RecAreaFacilities_API_v1.json and links facilities to their rec area
type RecAreaFacilityLink struct {
	RecAreaID  int
	FacilityID int
}

// OrgEntityLink is the data type of data/OrgEntities_API_v1.json and link orgs to rec areas and facilities
type OrgEntityLink struct {
	OrgID      int
	EntityID   int
	EntityType string
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
	orgs := getOrgs()
	for _, org := range orgs {
		fmt.Printf("OrgID: %d\n", org.ID)
	}

	recAreas := getRecAreas()
	for _, recArea := range recAreas {
		fmt.Printf("RecAreaID: %d\n", recArea.ID)
	}

	facilities := getFacilities()
	for _, facility := range facilities {
		fmt.Printf("FacilityID: %d\n", facility.ID)
	}
}

func getOrgs() []Organization {
	// Organizations matches the data being input to the program
	type organizationsDataFormat struct {
		Organizations []Organization `json:"RECDATA"`
	}
	var orgs organizationsDataFormat

	orgData, err := ioutil.ReadFile("data/Organizations_API_v1.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(orgData, &orgs)
	if err != nil {
		log.Fatal(err)
	}
	return orgs.Organizations
}

func getRecAreas() []RecArea {
	// RecAreas matches the data being input to the program
	type recAreasDataFormat struct {
		RecAreas []RecArea `json:"RECDATA"`
	}
	var recAreas recAreasDataFormat

	recAreaData, err := ioutil.ReadFile("data/RecAreas_API_v1.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(recAreaData, &recAreas)
	if err != nil {
		log.Fatal(err)
	}
	return recAreas.RecAreas
}

func getFacilities() []Facility {
	// Facilities matches the being input to the program
	type facilitiesDataFormat struct {
		Facilities []Facility `json:"RECDATA"`
	}
	var facilities facilitiesDataFormat

	facilitiesData, err := ioutil.ReadFile("data/Facilities_API_v1.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(facilitiesData, &facilities)
	if err != nil {
		log.Fatal(err)
	}
	return facilities.Facilities
}
