package sure

import "strings"

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func FindX(x int, objects []*BucketObject) []*BucketObject {
	// Key with latest object, no matter how deep inside, determines root level freshness
	topLevel := make([]string, 0)
	removeObjects := make([]*BucketObject, 0)

	for _, o := range objects {
		tl := strings.Split(o.Key, "/")[0]

		// Pass through until X latest top level folders skipped
		if len(topLevel) == x && !Contains(topLevel, tl) {
			removeObjects = append(removeObjects, o)
			continue
		}
		if !Contains(topLevel, tl) {
			topLevel = append(topLevel, tl)
		}
	}
	return removeObjects
}
