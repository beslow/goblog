package ga

type Permission struct {
	ID         int
	Name       string
	Slug       string
	HttpMethod string
	HttpPath   string
}

func (permission Permission) GetID() int {
	return permission.ID
}
