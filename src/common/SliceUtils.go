package common

import ()

func SliceExistString(slice []string, str string) bool {
	for _, st := range slice {
		if Equals(st, str) {
			return true
		}
	}
	return false
}
