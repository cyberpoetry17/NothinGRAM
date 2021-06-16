package data

import (
	"encoding/json"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"io"
)

type Tag struct {
	ID      uuid.UUID `gorm:"primary_key,column:id"      json:"id"`
	TagName string    `gorm:"column:TagName"   json:"TagName"`
	Posts []Post    `gorm:"many2many:posts_tags"   json:"Posts"`
}


func (tag *Tag) BeforeCreate(scope *gorm.DB) error {
	tag.ID = uuid.New()
	return nil
}
//dekoduje usera u JSON format za slanje putem RESTa,ako ne-> vraca gresku
func (tag *Tag) FromJSON(response io.Reader) error {
	err := json.NewDecoder(response)
	return err.Decode(tag)
}

//prebacuje iz JSONa u GO
func (tag *Tag) ToJSON(write io.Writer) error {
	err := json.NewEncoder(write)
	return err.Encode(tag)
}
func (tag *Tag) ToJSONOne(write io.Writer) error {
	err := json.NewEncoder(write)
	return err.Encode(tag)
}