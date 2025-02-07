package model

import "time"


type ReportTotalInfo struct {
	TimeInitial 	time.Time
	TimeTotal 		time.Duration
	RequestTotal 	int64
	Status200  		int64
	StatusNOK 	    []StatusNOK
}

type StatusNOK struct {
	StatusCode int
	Count int64
}