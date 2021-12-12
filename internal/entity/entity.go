//Package entity has the structs that represents the problem domain
package entity

type Person struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
