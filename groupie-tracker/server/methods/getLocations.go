package methods

import "fmt"

func GetLocations() (map[string][]int8, error) {
	locationsMap := map[string][]int8{}
	var laocationStruct Locations
	for i := 1; i <= 52; i++ {
		url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%v", i)
		err := FetchParser(url, &laocationStruct)
		if err != nil {
			return nil, err
		}

		for _, loc := range laocationStruct.Locations {
			// user id (i) is for maping location with user
			locationsMap[loc] = append(locationsMap[loc], int8(i))
		}
	}
	
	return locationsMap, nil
}
