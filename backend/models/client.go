package models

type Client struct {
	Id              uint   `json:"id"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	BusinessType    uint   `json:"businessType"`
	ResidentialType uint   `json:"residentialType"`
}
