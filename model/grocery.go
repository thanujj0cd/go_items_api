package model

//groceries model. This file likely contains the data model representing a grocery item.
// type Grocery struct {
// 	Name    string `json: "name"`
// 	Quality int    `json: "quantity"`
// }

type Grocery struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}
