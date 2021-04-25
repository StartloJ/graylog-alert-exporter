// Package utils provides utility function to help and reduce dumplicate code
package utils

import (
	"encoding/json"
	"log"

	"github.com/itchyny/gojq"
	"github.com/sirupsen/logrus"
)

// PrettyJSON return string in pretty format for struct
func PrettyJSON(s interface{}) string {
	pretty, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	return string(pretty)
}

// GetValueFromJSON return value of json in specific path
func GetValueFromJSON(path string, j interface{}) (interface{}, error) {
	logrus.Debugf("Query path: %s", path)
	logrus.Debugf("Query Json:\n%s", PrettyJSON(j))

	query, err := gojq.Parse(path)
	if err != nil {
		logrus.Fatal(err)
	}

	var input map[string]interface{}
	b, err := json.Marshal(j)
	if err != nil {
		logrus.Error(err)
	}
	err = json.Unmarshal(b, &input)
	if err != nil {
		logrus.Error(err)
	}

	var result interface{}
	iter := query.Run(input)
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}

		if err, ok := v.(error); ok {
			logrus.Fatal(err)
			return nil, err
		}

		result = v
		logrus.Debugf("Query value: %v", result)
	}

	return result, nil
}
