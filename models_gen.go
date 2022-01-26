// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package user

import (
	"time"

	"github.com/amaru0601/go-rent/ent/property"
)

type PropertyInput struct {
	Class       property.Class `json:"class"`
	Owner       int            `json:"owner"`
	Address     string         `json:"address"`
	Description string         `json:"description"`
	City        string         `json:"city"`
}

type UserInput struct {
	Birthday  *time.Time `json:"birthday"`
	Names     string     `json:"names"`
	Lastnames string     `json:"lastnames"`
	Activate  *bool      `json:"activate"`
	Email     string     `json:"email"`
}
