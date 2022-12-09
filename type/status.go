package _type

type RoleType string

var (
	Active  RoleType = "Active"
	Standby RoleType = "Standby"
)

type Status struct {
	role RoleType
}

func (s Status) Role() RoleType {
	return s.role
}

func (s *Status) SetRole(role RoleType) {
	s.role = role
}
