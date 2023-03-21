package app

type TimeResponse struct {
	Status           string `json:"status"`
	Message          string `json:"message"`
	CountryCode      string `json:"countryCode"`
	CountryName      string `json:"countryName"`
	RegionName       string `json:"regionName"`
	CityName         string `json:"cityName"`
	ZoneName         string `json:"zoneName"`
	Abbreviation     string `json:"abbreviation"`
	GmtOffset        int    `json:"gmtOffset"`
	Dst              string `json:"dst"`
	ZoneStart        int    `json:"zoneStart"`
	ZoneEnd          int    `json:"zoneEnd"`
	NextAbbreviation string `json:"nextAbbreviation"`
	Timestamp        int    `json:"timestamp"`
	Formatted        string `json:"formatted"`
}
