package core

type Room struct {
	Name  string
	Type  string
	Owner string
	Price float32
}

func NewRoom(name, owner, roomType string, price float32) *Room {
	return &Room{
		Name:  name,
		Type:  roomType,
		Owner: owner,
		Price: price,
	}
}

type ByPrice []*Room

func (a ByPrice) Len() int {
	return len(a)
}

func (a ByPrice) Less(i, j int) bool {
	return a[i].Price < a[j].Price
}

func (a ByPrice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
