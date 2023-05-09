package refrigerator

type Food struct {
	Name   *string `json:"name"`
	Volume int     `json:"volume"`
	Weight int     `json:"weight"`
}
