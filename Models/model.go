package models

// Todo struct is a model that represents a todo item
type Todo struct {
	ID uint            `json:"id"`
	Title string       `json:"title"`
	Description string `json:"description"`
}

// TableName sets the insert table name for this struct type
func (b *Todo) TableName() string {
	return "todos"
}
