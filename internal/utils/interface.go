package utils

import (
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"
)

// GetInterfaceValueAsString will return any value of interface as string type
func GetInterfaceValueAsString(val interface{}) string {
	switch t := val.(type) {
	case int:
		logrus.Debugf("%d == %T\n", t, t)
		return strconv.Itoa(t)
	case int8:
		logrus.Debugf("%d == %T\n", t, t)
		return strconv.Itoa(int(t))
	case int16:
		logrus.Debugf("%d == %T\n", t, t)
		return strconv.Itoa(int(t))
	case int32:
		logrus.Debugf("%d == %T\n", t, t)
		return strconv.Itoa(int(t))
	case int64:
		logrus.Debugf("%d == %T\n", t, t)
		return strconv.Itoa(int(t))
	case bool:
		logrus.Debugf("%t == %T\n", t, t)
		return strconv.FormatBool(t)
	case float32:
		logrus.Debugf("%g == %T\n", t, t)
		return fmt.Sprintf("%f", t)
	case float64:
		logrus.Debugf("%f == %T\n", t, t)
		return fmt.Sprintf("%f", t)
	case uint8:
		logrus.Debugf("%d == %T\n", t, t)
		return strconv.Itoa(int(t))
	case uint16:
		logrus.Debugf("%d == %T\n", t, t)
		return strconv.Itoa(int(t))
	case uint32:
		logrus.Debugf("%d == %T\n", t, t)
		return strconv.Itoa(int(t))
	case uint64:
		logrus.Debugf("%d == %T\n", t, t)
		return strconv.Itoa(int(t))
	case string:
		logrus.Debugf("%s == %T\n", t, t)
		return t
	default:
		logrus.Errorf("not found any type to convert for %T\n", t)
	}

	return ""
}
