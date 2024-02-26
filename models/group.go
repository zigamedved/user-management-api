package models

type Group struct {
	Document `bson:",inline" json:",inline"`
	Name     string `bson:"name" json:"name" validate:"required,max=100"`
}
