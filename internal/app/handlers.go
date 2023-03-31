package app

import (
	"strconv"
	"time"
)

func (s service) GetTimeByCity(chatID int64, city string) (string, error) {
	date := time.Now()
	defer s.db.AddEntry(chatID, city, date)

	location, err := s.client.GetLocation(city)
	if err != nil {
		return "", err
	}

	timeByCity, err := s.client.GetTimeByLocation(location.Lat, location.Lng)
	if err != nil {
		return "", err
	}

	return timeByCity, nil
}

func (s service) GetStats(chatID int64) (string, error) {
	firstEntryDate, err := s.db.GetFirstEntryDate(chatID)
	if err != nil {
		return "", err
	}

	lastEntryDate, err := s.db.GetLastEntryDate(chatID)
	if err != nil {
		return "", err
	}

	entriesCount, err := s.db.GetEntriesCount(chatID)
	if err != nil {
		return "", err
	}

	return "First request: " + firstEntryDate.String() + "\n" +
		"Last request: " + lastEntryDate.String() + "\n" +
		"Total count: " + strconv.Itoa(int(entriesCount)), nil
}

func (s service) GetLocation(chatID int64, city string) (string, error) {
	date := time.Now()
	defer s.db.AddEntry(chatID, city, date)

	location, err := s.client.GetLocation(city)
	if err != nil {
		return "", err
	}

	return "Latitude:" + strconv.FormatFloat(location.Lat, 'G', -1, 64) +
		"\nLongtitude:" + strconv.FormatFloat(location.Lng, 'G', -1, 64), nil
}
