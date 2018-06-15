package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	output := ""

	orgs := getOrgs()
	orgs = append(orgs, Organization{ID: 0, Name: "ROOT"})
	for _, org := range orgs {
		if org.ID != 0 {
			parentNode, err := GetOrgByID(orgs, org.ParentID)
			checkError(err)
			org.SetParent(parentNode)
		}
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

	orgEntityLinks := getOrgEntityLinks()
	for _, orgEntityLink := range orgEntityLinks {
		parentNode, err := GetOrgByID(orgs, orgEntityLink.OrgID)
		checkError(err)
		if orgEntityLink.EntityType == "RecArea" {
			childNode, err := GetRecAreaByID(recAreas, orgEntityLink.EntityID)
			checkError(err)
			childNode.SetParent(parentNode)
		} else if orgEntityLink.EntityType == "Facility" {
			childNode, err := GetFacilityByID(facilities, orgEntityLink.EntityID)
			checkError(err)
			childNode.SetParent(parentNode)
		}
	}

	err := ioutil.WriteFile("hello.txt", []byte(output), 0644)
	checkError(err)
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

// GetNodeByID takes a slice of nodes (i.e. orgs) and returns the node that matches the provided id
// func GetNodeByID(nodes []TreeNode, id int) TreeNode {
// 	for _, node := range nodes {
// 		if node.GetID() == id {
// 			return node
// 		}
// 	}
// 	return nil
// }

// Organization is the data type of organizations from data/Organizations_API_v1.json and IMPLEMENTS TreeNode
type Organization struct {
	ID       int    `json:"OrgID"`
	Name     string `json:"OrgName"`
	ParentID int    `json:"OrgParentID"`
	Parent   TreeNode
	Children []TreeNode
}

func getOrgs() []Organization {
	type organizationsDataFormat struct {
		Organizations []Organization `json:"RECDATA"`
	}
	var orgs organizationsDataFormat

	orgData, err := ioutil.ReadFile("data/Organizations_API_v1.json")
	checkError(err)
	err = json.Unmarshal(orgData, &orgs)
	checkError(err)
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
	parent.AddChild(o)
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

// GetOrgByID finds an org by ID
func GetOrgByID(orgs []Organization, id int) (Organization, error) {
	for _, org := range orgs {
		if org.ID == id {
			return org, nil
		}
	}
	return Organization{}, errors.New("GetOrgByID: ID out of range")
}

// RecArea is the data type of recreation areas from data/RecAreas_API_v1.json and IMPLEMENTS TreeNode
type RecArea struct {
	ID       int    `json:"RecAreaID"`
	Name     string `json:"RecAreaName"`
	Parent   TreeNode
	Children []TreeNode
}

func getRecAreas() []RecArea {
	type recAreasDataFormat struct {
		RecAreas []RecArea `json:"RECDATA"`
	}
	var recAreas recAreasDataFormat

	recAreaData, err := ioutil.ReadFile("data/RecAreas_API_v1.json")
	checkError(err)
	err = json.Unmarshal(recAreaData, &recAreas)
	checkError(err)
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
	parent.AddChild(r)
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

// GetRecAreaByID finds an recArea by ID
func GetRecAreaByID(recAreas []RecArea, id int) (RecArea, error) {
	for _, recArea := range recAreas {
		if recArea.ID == id {
			return recArea, nil
		}
	}
	return RecArea{}, errors.New("GetRecAreaByID: ID out of range")
}

// Facility is the data type of facilities from data/Facilities_API_v1.json and IMPLEMENTS TreeNode
type Facility struct {
	Name   string `json:"FacilityName"`
	ID     int    `json:"FacilityID"`
	Parent TreeNode
}

func getFacilities() []Facility {
	type facilitiesDataFormat struct {
		Facilities []Facility `json:"RECDATA"`
	}
	var facilities facilitiesDataFormat

	facilitiesData, err := ioutil.ReadFile("data/Facilities_API_v1.json")
	checkError(err)
	err = json.Unmarshal(facilitiesData, &facilities)
	checkError(err)
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
	parent.AddChild(f)
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

// GetFacilityByID finds an facility by ID
func GetFacilityByID(facilities []Facility, id int) (Facility, error) {
	for _, facility := range facilities {
		if facility.ID == id {
			return facility, nil
		}
	}
	return Facility{}, errors.New("GetFacilityByID: ID out of range")
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

func getOrgEntityLinks() []OrgEntityLink {
	type orgEntityLinksDataFormat struct {
		OrgEntityLinks []OrgEntityLink `json:"RECDATA"`
	}
	var orgEntityLinks orgEntityLinksDataFormat

	orgEntityLinksData, err := ioutil.ReadFile("data/Organizations_API_v1.json")
	checkError(err)
	err = json.Unmarshal(orgEntityLinksData, &orgEntityLinks)
	checkError(err)
	return orgEntityLinks.OrgEntityLinks
}

func checkError(err error) {
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
