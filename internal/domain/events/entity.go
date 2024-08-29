package events

type IDer interface {
	ID() string
}

type EntityNamer interface {
	EntityName() string
}

type Entity struct {
	Id   string
	Name string
}

var _ interface {
	IDer
	EntityNamer
} = (*Entity)(nil)

func NewEntity(id, name string) Entity {
	return Entity{
		Id:   id,
		Name: name,
	}
}

func (e Entity) ID() string             { return e.Id }
func (e Entity) EntityName() string     { return e.Name }
func (e Entity) Equals(other IDer) bool { return e.Id == other.ID() }

func (e *Entity) setID(id string)     { e.Id = id }
func (e *Entity) setName(name string) { e.Name = name }
