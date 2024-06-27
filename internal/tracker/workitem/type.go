package workitem

type Type struct {
	Id           string `json:"id"`
	AutoAttached bool   `json:"autoAttached"`
	Name         string `json:"name"`
}
