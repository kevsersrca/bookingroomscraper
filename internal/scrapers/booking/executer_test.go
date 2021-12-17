package booking

import (
	"testing"
)

func TestBookingRoomExtractor(t *testing.T) {
	type TestCase struct {
		url       string
		useragent string
		invalid   bool
	}
	//TODO: can be add new test cases
	cases := []TestCase{
		TestCase{
			"https://www.booking.com/searchresults.en-gb.html?label=gen173nr-1FCAEoggI46AdIM1gEaJUCiAEBmAEJuAEHyAEN2AEB6AEB-AELiAIBqAIDuAKf1vGNBsACAdICJDkzMzFmMzJlLTBhN2YtNDFlZi05ODA5LWQ3NTk4YzU0N2M2ONgCBuACAQ&lang=en-gb&sid=a5ecc01b37543d91e33174990c4f3658&sb=1&sb_lp=1&src=index&src_elem=sb&error_url=https%3A%2F%2Fwww.booking.com%2Findex.en-gb.html%3Flabel%3Dgen173nr-1FCAEoggI46AdIM1gEaJUCiAEBmAEJuAEHyAEN2AEB6AEB-AELiAIBqAIDuAKf1vGNBsACAdICJDkzMzFmMzJlLTBhN2YtNDFlZi05ODA5LWQ3NTk4YzU0N2M2ONgCBuACAQ%3Bsid%3Da5ecc01b37543d91e33174990c4f3658%3Bsb_price_type%3Dtotal%26%3B&ss=Lapland%2C+Finland&is_ski_area=&checkin_year=2021&checkin_month=12&checkin_monthday=29&checkout_year=2022&checkout_month=1&checkout_monthday=31&group_adults=2&group_children=0&no_rooms=1&b_h4u_keep_filters=&from_sf=1&ss_raw=lapland&ac_position=0&ac_langcode=en&ac_click_type=b&dest_id=2989&dest_type=region&place_id_lat=67.46279&place_id_lon=25.377645&search_pageview_id=3b4e4c0f25fa0091&search_selected=true&search_pageview_id=3b4e4c0f25fa0091&ac_suggestion_list_length=5&ac_suggestion_theme_list_length=0",
			"",
			false,
		},
		TestCase{
			"",
			"",
			true,
		},
		TestCase{
			"https://www.booking.com/searchresults.en-gb.html?label=gen173nr-1FCAEoggI46AdIM1gEaJUCiAEBmAEJuAEHyAEN2AEB6AEB-AELiAIBqAIDuAKf1vGNBsACAdICJDkzMzFmMzJlLTBhN2YtNDFlZi05ODA5LWQ3NTk4YzU0N2M2ONgCBuACAQ&lang=en-gb&sid=a5ecc01b37543d91e33174990c4f3658&sb=1&sb_lp=1&src=index&src_elem=sb&error_url=https%3A%2F%2Fwww.booking.com%2Findex.en-gb.html%3Flabel%3Dgen173nr-1FCAEoggI46AdIM1gEaJUCiAEBmAEJuAEHyAEN2AEB6AEB-AELiAIBqAIDuAKf1vGNBsACAdICJDkzMzFmMzJlLTBhN2YtNDFlZi05ODA5LWQ3NTk4YzU0N2M2ONgCBuACAQ%3Bsid%3Da5ecc01b37543d91e33174990c4f3658%3Bsb_price_type%3Dtotal%26%3B&ss=Lapland%2C+Finland&is_ski_area=&checkin_year=2021&checkin_month=12&checkin_monthday=29&checkout_year=2022&checkout_month=1&checkout_monthday=31&group_adults=2&group_children=0&no_rooms=1&b_h4u_keep_filters=&from_sf=1&ss_raw=lapland&ac_position=0&ac_langcode=en&ac_click_type=b&dest_id=2989&dest_type=region&place_id_lat=67.46279&place_id_lon=25.377645&search_pageview_id=3b4e4c0f25fa0091&search_selected=true&search_pageview_id=3b4e4c0f25fa0091&ac_suggestion_list_length=5&ac_suggestion_theme_list_length=0",
			"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36",
			false,
		},
	}

	for _, ts := range cases {
		executor := BookingScraperExecutor{
			bookingUrl: ts.url,
			userAgent:  ts.useragent,
		}
		rooms, err := executor.ExtractRooms()

		invalid := err != nil
		invalid = len(rooms) == 0

		if invalid != ts.invalid {
			t.Errorf("invalid case\n\twant: %+v\n\tgot : %+v\n\terr : %s", ts.invalid, invalid, err)
		}
		if invalid {
			return
		}
	}

}
