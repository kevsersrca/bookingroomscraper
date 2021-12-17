package scrapers

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"bookingroomscraper/internal/core"
)

type ScraperExecutor interface {
	ExtractRooms() ([]*core.Room, error)
}

type Scraper struct {
	executor ScraperExecutor
}

func NewScraper(executor ScraperExecutor) *Scraper {
	return &Scraper{
		executor: executor,
	}
}

func (s Scraper) GetBestRoomPrice() (*core.Room, error) {
	rooms, err := s.executor.ExtractRooms()
	if err != nil {
		return nil, err
	}
	if len(rooms) == 0 {
		return nil, fmt.Errorf("no room found")
	}

	sort.Sort(core.ByPrice(rooms))
	return rooms[0], nil
}

func ExtractFloatValue(valueStr string) float32 {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	floatStr := re.FindString(valueStr)
	if len(floatStr) != 0 {
		// TODO improve numeric format conversion
		floatStr := strings.ReplaceAll(floatStr, ",", "")
		floatValue, _ := strconv.ParseFloat(floatStr, 32)
		return float32(floatValue)
	}

	return 0
}
