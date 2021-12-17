package main

import (
	"bookingroomscraper/internal/scrapers"
	"bookingroomscraper/internal/scrapers/booking"
	"bookingroomscraper/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load(".env")
}

func main() {
	bookingUrl := os.Getenv("BOOKING_URL")
	if bookingUrl == "" {
		log.Println("Booking url not found!")
		return
	}

	userAgent := os.Getenv("USER_AGENT")
	if userAgent == "" {
		log.Println("Default user agent using")
		userAgent = utils.DefaultUserAgent
	}

	scraper := scrapers.NewScraper(booking.NewBookingScraperExecutor(bookingUrl, userAgent))
	room, err := scraper.GetBestRoomPrice()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Cheapest price is %.2f,  %s and  %s. scraped from %s", room.Price, room.Name, room.Type, room.Owner)
}
