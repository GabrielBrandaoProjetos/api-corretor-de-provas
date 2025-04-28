package domain

import (
	"time"

	"github.com/google/uuid"
)

type Unit string

const (
	UnitHead   = "head"
	UnitBranch = "branch"
)

type School struct {
	ID        string
	Name      string
	Register  string
	Unit      Unit
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewSchool(name, register, unit, address string) *School {
	return &School{
		ID:        uuid.New().String(),
		Name:      name,
		Register:  register,
		Unit:      Unit(unit),
		Address:   address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
