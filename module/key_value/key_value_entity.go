package key_value

import (
	"encoding/json"
	"fmt"
	"time"
)

//Entity KeyValue entity
type Entity struct {
	Key   string `bson:"key" json:"key"`
	Value string `bson:"value" json:"value"`

	CreatedAt time.Time `bson:"created_at, omitempty"`
	UpdatedAt time.Time `bson:"updated_at, omitempty"`
}

//BeforeCreate Modifies KeyValue before create
func (e *Entity) BeforeCreate() *Entity {
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	return e
}

//BeforeUpdate Modifies KeyValue before update
func (e *Entity) BeforeUpdate() *Entity {
	e.UpdatedAt = time.Now()
	return e
}

//DecodeJson decodes string to KeyValue Entity
func DecodeJson(d string) *Entity {
	var keyValue Entity
	json.Unmarshal([]byte(d), &keyValue)
	return &keyValue
}

//EncodeJson encodes KeyValue Entity to String
func EncodeJson(e Entity) string {
	data, err := json.Marshal(e)
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(data[:])
}
