package json_utils

type Employee struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Designation string    `json:"designation,omitempty"`
	Addresses   []Address `json:"address,omitempty"`
}

type Address struct {
	DoorNumber int    `json:"doorNumber,omitempty"`
	Street     string `json:"street,omitempty"`
	City       string `json:"city,omitempty"`
	Country    string `json:"country,omitempty"`
}
