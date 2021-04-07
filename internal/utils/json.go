// Package utils provides utility function to help and reduce dumplicate code
package utils

import (
	"encoding/json"
	"log"
)

// PrettyJSON return string in pretty format for struct
func PrettyJSON(s interface{}) string {
	pretty, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	return string(pretty)
}
