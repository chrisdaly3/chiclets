package database

type Venue struct {
	ID         uint   `gorm:"primary_key;not null"`
	Name       string `json:"name"`
	City       string `json:"city"`
	TimeZoneID uint
	TimeZone   TimeZone `json:"timeZone"`
}

type TimeZone struct {
	ID     string `json:"id"`
	Offset int    `json:"offset"`
	TZ     string `json:"tz"`
}

type Division struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	NameShort    string `json:"nameShort"`
	Link         string `json:"link"`
	Abbreviation string `json:"abbreviation"`
}

type Conference struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

type Franchise struct {
	ID       uint   `json:"franchiseId"`
	TeamName string `json:"teamName"`
	Link     string `json:"link"`
}

type Team struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	Abbreviation    string `json:"abbreviation"`
	TeamName        string `json:"teamName"`
	LocationName    string `json:"locationName"`
	FirstYearOfPlay string `json:"firstYearOfPlay"`
	DivisionID      uint
	Division        Division `json:"division"`
	VenueID         uint
	Venue           Venue `json:"venue"`
	ConferenceID    uint
	Conference      Conference `json:"conference"`
	FranchiseID     uint
	Franchise       Franchise `json:"franchise"`
	ShortName       string    `json:"shortName"`
	OfficialSiteURL string    `json:"officialSiteUrl"`
	Active          bool      `json:"active"`
}

type NHLResponse struct {
	Copyright string `json:"copyright"`
	Teams     []Team `json:"teams"`
}
