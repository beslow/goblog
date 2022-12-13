package ga

type Menu struct {
	ID         int
	ParentID   int
	Type       int
	Order      int
	Title      string
	Icon       string
	Uri        string
	Header     string
	PluginName string
	Uuid       string
}

func (menu Menu) GetID() int {
	return menu.ID
}
