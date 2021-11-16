package key_value

import "time"

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
