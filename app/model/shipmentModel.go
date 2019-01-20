package model

type ShipmentModel struct{
	Name    string `json:"name"`
	Zipcode string `json:"zipcode"`
	Street string `json:"street"`
    Neighborhood string `json:"neighborhood"`
    StateShort string `json:"state_short"`
    City string `json:"city"`

}