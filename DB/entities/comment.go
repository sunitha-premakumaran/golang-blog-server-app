package entities

import "gorm.io/gorm"

type Reaction string

const (
	Unlike Reaction = "unlike"
	Like            = "like"
	Heart 			= "heart"
	Clap			= "clap"
)

type CommentReaction struct {
	Reaction Reaction
	MadeBy	User
}


func (j CommentReaction) Value() (driver.Value, error) {
	valueString, err := json.Marshal(j)
	return string(valueString), err
}

func (j *CommentReaction) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &j); err != nil {
		return err
	}
	return nil
}


type Comment struct {
	gorm.Model

	Content 	string

	CommentReactions   []CommentReaction   `sql:"type:jsonb"`
	
	Author  	User `gorm:"foreignKey:ID"`
}

