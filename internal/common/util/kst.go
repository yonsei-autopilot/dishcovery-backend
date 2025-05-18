package util

import "time"

var kstLoc *time.Location

func InitializeKst() {
	kst, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		panic("failed to load kst")
	}
	kstLoc = kst
}

func GetKstNow() time.Time {
	return time.Now().In(kstLoc)
}
