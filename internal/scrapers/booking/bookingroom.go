package booking

import (
	"bookingroomscraper/internal/core"
	"bookingroomscraper/internal/scrapers"
)

type bookingRoom struct {
	HotelName string
	RoomType  string
	Review    string
	Price     string
}

func (b bookingRoom) toCoreRoom() *core.Room {
	return core.NewRoom(
		b.HotelName,
		"booking",
		b.RoomType,
		scrapers.ExtractFloatValue(b.Price),
	)
}

func toCoreRooms(rooms []bookingRoom) []*core.Room {
	result := make([]*core.Room, len(rooms))
	for pos, room := range rooms {
		result[pos] = room.toCoreRoom()
	}

	return result
}
