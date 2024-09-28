package methods

type Artist struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Relations    string
}

type Relation struct {
	DatesLocations map[string][]string
}

type Locations struct {
	Locations []string
}

type Dates struct {
	Dates []string
}
