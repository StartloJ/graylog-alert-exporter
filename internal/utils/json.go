// Package utils provides utility function to help and reduce dumplicate code
package utils

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

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
func GetValueFromJSON(path string, j interface{}) interface{} {
	var obj map[string]interface{}
	b, err := json.Marshal(j)
	if err != nil {
		logrus.Panic(err)
	}
	json.Unmarshal(b, &obj)

	paths := strings.Split(path, ".")
	if len(paths) == 1 {
		return obj[path]
	}

	logrus.Debugf("Get value from json path %v \n%v\n", paths, PrettyJSON(obj))
	if _, err := strconv.Atoi(paths[1:][0]); err == nil {
		return getValueFromSlice(paths[1:], obj[paths[0]].([]interface{}))
	}
	return getValueFromMap(paths[1:], obj[paths[0]].(map[string]interface{}))
}

func getValueFromMap(paths []string, obj map[string]interface{}) interface{} {
	if len(paths) == 1 {
		return obj[paths[0]]
	}

	logrus.Debugf("Get value from json path %v \n%v\n", paths, PrettyJSON(obj))
	if _, err := strconv.Atoi(paths[1:][0]); err == nil {
		return getValueFromSlice(paths[1:], obj[paths[0]].([]interface{}))
	}
	return getValueFromMap(paths[1:], obj[paths[0]].(map[string]interface{}))
}

func getValueFromSlice(paths []string, obj []interface{}) interface{} {
	i, err := strconv.Atoi(paths[0])
	if err != nil {
		logrus.Fatal(err)
	}
	if len(paths) == 1 {

		return obj[i]
	}

	logrus.Debugf("Get value from json path %v \n%v\n", paths, PrettyJSON(obj))
	return getValueFromMap(paths[1:], obj[i].(map[string]interface{}))
}
