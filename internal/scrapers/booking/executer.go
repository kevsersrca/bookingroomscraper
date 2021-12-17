package booking

import (
	"fmt"

	"bookingroomscraper/internal/core"

	"github.com/gocolly/colly"
)

type BookingScraperExecutor struct {
	bookingUrl string
	userAgent  string
}

func NewBookingScraperExecutor(bookingUrl, userAgent string) *BookingScraperExecutor {
	return &BookingScraperExecutor{
		bookingUrl: bookingUrl,
		userAgent:  userAgent,
	}
}

func (s *BookingScraperExecutor) ExtractRooms() ([]*core.Room, error) {
	c := colly.NewCollector(colly.UserAgent(s.userAgent))

	records := make([]bookingRoom, 0)

	c.OnHTML("#search_results_table", func(table *colly.HTMLElement) {
		room := bookingRoom{}
		table.ForEach("div._fe1927d9e._0811a1b54", func(_ int, tableCol *colly.HTMLElement) {
			// Hotel Room Type
			room.RoomType = tableCol.ChildText("span._c5d12bf22")

			//Price
			room.Price = tableCol.ChildText("span.fde444d7ef._e885fdc12")

			// Hotel name
			room.HotelName = tableCol.ChildText("div.fde444d7ef._c445487e2")

			// Review
			room.Review = tableCol.ChildText("div._9c5f726ff.bd528f9ea6")

		})

		// Set the criteria
		if room.HotelName != "" || room.Price != "" {
			records = append(records, room)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Scraping...")
	})

	c.Visit(s.bookingUrl)

	return toCoreRooms(records), nil
}
