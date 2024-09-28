package methods

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
// function that Fetches URLs
// get the data from the url to local
func FetchParser(url string, result interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request failed with status %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	
	defer resp.Body.Close()
	return json.Unmarshal(body, result)
}
