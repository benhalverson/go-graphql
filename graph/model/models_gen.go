// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Dog struct {
	ID     string `json:"_id" bson:"_id"`
	Name   string `json:"name"`
	IsGood bool   `json:"isGood"`
}

type NewDog struct {
	Name   string `json:"name"`
	IsGood bool   `json:"isGood"`
}
