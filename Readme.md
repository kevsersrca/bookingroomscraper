# Best Room Price Scraper from Booking.com

This repo is a tutorial of Large Scale Crawling and Scraping Best Practices in Gophercon 2021. That finds the cheapest room in the given booking url.

## Usage

Copy .env.sample to .env
Set your creds.

and then,

```
go run main.go
```


## Tests
To run all tests
``` 
go test -race $(go list ./...) 
```

## Codefactor
[![CodeFactor](https://www.codefactor.io/repository/github/kevsersrca/bookingroomscraper/badge)](https://www.codefactor.io/repository/github/kevsersrca/bookingroomscraper)
