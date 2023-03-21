package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/codingsince1985/geo-golang"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func (s service) GetTimeByCity(chatID int64, city string) (string, error) {
	date := time.Now()
	defer s.db.AddEntry(chatID, city, date)

	location, err := s.getLocation(city)
	if err != nil {
		return "", err
	}

	timeByCity, err := s.getTimeByLocation(location.Lat, location.Lng)
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

	location, err := s.getLocation(city)
	if err != nil {
		return "", err
	}

	return "Latitude:" + strconv.FormatFloat(location.Lat, 'G', -1, 64) +
		"\nLongtitude:" + strconv.FormatFloat(location.Lng, 'G', -1, 64), nil
}

func (s service) getLocation(city string) (*geo.Location, error) {
	count := 0
	location := &geo.Location{}
	var err error
	for count < 10 {
		location, err = s.geocoder.Geocode(city)
		if err == nil {
			break
		}
		count++
	}
	if err != nil {
		return nil, err
	}

	return location, err
}

func (s service) getTimeByLocation(lat, lng float64) (string, error) {
	timeURL, err := url.Parse("http://api.timezonedb.com/v2.1/get-time-zone")
	if err != nil {
		s.logger.Error().Err(err).Msg("create url")
		return "", err
	}

	query := url.Values{
		"key":    []string{s.config.TimeToken},
		"format": []string{"json"},
		"by":     []string{"position"},
		"lat":    []string{strconv.FormatFloat(lat, 'G', -1, 64)},
		"lng":    []string{strconv.FormatFloat(lng, 'G', -1, 64)},
	}

	req, err := http.NewRequest(http.MethodGet, timeURL.String()+"?"+query.Encode(), nil)
	if err != nil {
		s.logger.Error().Err(err).Msg("create request")
		return "", err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		s.logger.Error().Err(err).Msg("get time request")
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s.logger.Error().Err(err).Msg("response read")
		return "", err
	}
	fmt.Println(string(body))

	var timeResponse TimeResponse
	err = json.Unmarshal(body, &timeResponse)
	if err != nil {
		s.logger.Error().Err(err).Msg("response unmarshal")
		return "", err
	}

	fmt.Println(timeResponse)

	if timeResponse.Timestamp == 0 {
		s.logger.Error().Err(err).Msg("response returned 0 as unix time")
		return "", errors.New("city doesn't exist")
	}

	return time.Unix(int64(timeResponse.Timestamp), 0).String(), nil
}
