package model


type ReportTotalInfo struct {
	TempoTotal 		float64
	RequestTotal 	int64
	Status200  		int64
	StatusNOK 	    []StatusNOK
}

type StatusNOK struct {
	StatusCode int
	Count int64
}