package httprequest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ReadMap(resp *http.Response) (r map[string]interface{}, err error) {
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		err = json.NewDecoder(resp.Body).Decode(&r)
		return r, err
	}

	return
}

func ReadStruct[T any](resp *http.Response, r *T) error {
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		err := json.NewDecoder(resp.Body).Decode(&r)
		return err
	}

	return fmt.Errorf("Error: %#v", resp)
}
