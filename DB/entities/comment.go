package entities

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type JSON json.RawMessage

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}

type Comment struct {
	BaseModel
	Content          string
	CommentReactions JSON `sql:"type:jsonb"`
	PostedByUserID   sql.NullString
	BelongsToBlogID  sql.NullString
	PostedBy         User `gorm:"references:ID;foreignKey:PostedByUserID"`
	BelongsToBlog    Blog `gorm:"references:ID;foreignKey:BelongsToBlogID"`
}
