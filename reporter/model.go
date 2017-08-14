package reporter

import "time"

var (
	countByDay     map[string]countBytes
	filesProcessed []string
	noMvnoIMSIS    []int
)

type timeSlice []time.Time

type countBytes struct {
	Download    int            `json:"download"`
	Upload      int            `json:"upload"`
	Sumbytes    int            `json:"sumbytes"`
	RatingGroup map[int]int    `json:"ratingGroup"`
	Mvno        map[string]int `json:"mvno"`
}

func (s timeSlice) Len() int {
	return len(s)
}

func (s timeSlice) Less(i, j int) bool {
	return s[i].Before(s[j])
}

func (s timeSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
