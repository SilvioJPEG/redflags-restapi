package apiserver

type Player struct {
	ID   string `json:"uid"`
	Name string `json:"name"`
	Role *role  `json:"role"`
}

type role struct {
	ID       string `json:"uid"`
	RoleName string `json:"name"`
}
