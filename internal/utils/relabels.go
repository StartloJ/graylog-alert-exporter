// Package utils provides utility function to help and reduce dumplicate code
package utils

import "errors"

func RePriority(priority int) (string, error) {
	switch priority {
	case 0:
		return "resolved", nil
	case 1:
		return "minor", nil
	case 2:
		return "major", nil
	case 3:
		return "critical", nil
	}
	return "", errors.New("Cannot relabels priority to standard severity")
}
