package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Orgs is a slice of all organizations
var Orgs []Organization

func main() {
	output := ""

	Orgs = getOrgs()
	Orgs = append(Orgs, Organization{ID: 0, Name: "ROOT"})
	for _, org := range Orgs {
		output += fmt.Sprintf("OrgID: %d\n", org.ID)
	}

	recAreas := getRecAreas()
	for _, recArea := range recAreas {
		output += fmt.Sprintf("RecAreaID: %d\n", recArea.ID)
	}

	facilities := getFacilities()
	for _, facility := range facilities {
		output += fmt.Sprintf("FacilityID: %d\n", facility.ID)
	}

	err := ioutil.WriteFile("hello.txt", []byte(output), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// TreeNode defines the methods for our n-ary tree of nodes
type TreeNode interface {
	GetID() int
	GetType() string
	GetParent() TreeNode
	SetParent(parent TreeNode)
	GetChildren() []TreeNode
	AddChild(c TreeNode)
	GetName() string
}

// Organization is the data type of organizations from data/Organizations_API_v1.json and IMPLEMENTS TreeNode
type Organization struct {
	ID       int    `json:"OrgID"`
	Name     string `json:"OrgName"`
	ParentID int    `json:"OrgParentID"`
	Parent   TreeNode
	Children []TreeNode
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

// GetID returns an Organization's ID
func (o Organization) GetID() int {
	return o.ID
}

// GetType returns Organization when called on an Organization object
func (o Organization) GetType() string {
	return "Organization"
}

// GetParent returns an Organization's parent TreeNode
func (o Organization) GetParent() TreeNode {
	return o.Parent
}

// SetParent sets an Organization's parent TreeNode
func (o Organization) SetParent(parent TreeNode) {
	o.Parent = parent
}

// GetChildren returns an Organization's children
func (o Organization) GetChildren() []TreeNode {
	return o.Children
}

// AddChild adds a child TreeNode to an Organization
func (o Organization) AddChild(child TreeNode) {
	o.Children = append(o.Children, child)
}

// GetName returns an Organization's name
func (o Organization) GetName() string {
	return o.Name
}

// RecArea is the data type of recreation areas from data/RecAreas_API_v1.json and IMPLEMENTS TreeNode
type RecArea struct {
	ID       int    `json:"RecAreaID"`
	Name     string `json:"RecAreaName"`
	Parent   TreeNode
	Children []TreeNode
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

// GetID returns an RecArea's ID
func (r RecArea) GetID() int {
	return r.ID
}

// GetType returns RecArea when called on an RecArea object
func (r RecArea) GetType() string {
	return "RecArea"
}

// GetParent returns an RecArea's parent TreeNode
func (r RecArea) GetParent() TreeNode {
	return r.Parent
}

// SetParent sets an RecArea's parent TreeNode
func (r RecArea) SetParent(parent TreeNode) {
	r.Parent = parent
}

// GetChildren returns an RecArea's children
func (r RecArea) GetChildren() []TreeNode {
	return r.Children
}

// AddChild adds a child TreeNode to an RecArea
func (r RecArea) AddChild(child TreeNode) {
	r.Children = append(r.Children, child)
}

// GetName returns an RecArea's name
func (r RecArea) GetName() string {
	return r.Name
}

// Facility is the data type of facilities from data/Facilities_API_v1.json and IMPLEMENTS TreeNode
type Facility struct {
	Name   string `json:"FacilityName"`
	ID     int    `json:"FacilityID"`
	Parent TreeNode
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

// GetID returns an Facility's ID
func (f Facility) GetID() int {
	return f.ID
}

// GetType returns Facility when called on an Facility object
func (f Facility) GetType() string {
	return "Facility"
}

// GetParent returns an Facility's parent TreeNode
func (f Facility) GetParent() TreeNode {
	return f.Parent
}

// SetParent sets an Facility's parent TreeNode
func (f Facility) SetParent(parent TreeNode) {
	f.Parent = parent
}

// GetChildren returns an Facility's children
func (f Facility) GetChildren() []TreeNode {
	return nil
}

// AddChild adds a child TreeNode to an Facility
func (f Facility) AddChild(child TreeNode) {
	log.Fatal("Facilities cannot have children")
}

// GetName returns an Facility's name
func (f Facility) GetName() string {
	return f.Name
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
