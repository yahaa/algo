package executor

import (
	"math/rand"
	"time"
)

// GetRemoteData get mock data from other remote storage server
func GetRemoteData(mink, maxk int) []Row {

	rand.Seed(int64(time.Now().UnixNano()))

	totalRow := rand.Int() % (maxk - mink)

	res := make([]Row, 0)
	keySet := make(map[int]struct{})

	for totalRow > 0 {
		key := rand.Intn(maxk-mink+1) + mink
		keySet[key] = struct{}{}

		totalRow--
	}

	for k := range keySet {
		res = append(res, Row{Key: k})
	}

	return res
}
