package client

import (
	"bot/internal/client/models"
	"encoding/json"
	"errors"
	"github.com/codingsince1985/geo-golang"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func (a adapter) GetLocation(city string) (*geo.Location, error) {
	count := 0
	location := &geo.Location{}
	var err error
	for count < 10 {
		location, err = a.geocoder.Geocode(city)
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

func (a adapter) GetTimeByLocation(lat, lng float64) (string, error) {
	timeURL, err := url.Parse(a.config.TimeURL)
	if err != nil {
		a.logger.Error().Err(err).Msg("create url")
		return "", err
	}

	query := url.Values{
		"key":    []string{a.config.TimeToken},
		"format": []string{"json"},
		"by":     []string{"position"},
		"lat":    []string{strconv.FormatFloat(lat, 'G', -1, 64)},
		"lng":    []string{strconv.FormatFloat(lng, 'G', -1, 64)},
	}

	req, err := http.NewRequest(http.MethodGet, timeURL.String()+"?"+query.Encode(), nil)
	if err != nil {
		a.logger.Error().Err(err).Msg("create request")
		return "", err
	}

	resp, err := a.client.Do(req)
	if err != nil {
		a.logger.Error().Err(err).Msg("get time request")
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		a.logger.Error().Err(err).Msg("response read")
		return "", err
	}

	var timeResponse models.TimeResponse
	err = json.Unmarshal(body, &timeResponse)
	if err != nil {
		a.logger.Error().Err(err).Msg("response unmarshal")
		return "", err
	}

	if timeResponse.Timestamp == 0 {
		a.logger.Error().Err(err).Msg("response returned 0 as unix time")
		return "", errors.New("city doesn't exist")
	}

	return time.Unix(int64(timeResponse.Timestamp), 0).String(), nil
}
