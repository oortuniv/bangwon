package response

import _type "bangwon/type"

type Status struct {
	Role string `json:"role"`
}

func From(state _type.Status) Status {
	return Status{
		Role: string(state.Role()),
	}
}

func (s Status) To() _type.Status {
	status := _type.Status{}
	status.SetRole(_type.RoleType(s.Role))
	return status
}
