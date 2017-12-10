package model

type Week struct {
	Games []Game   `json:"games"`
	Byes  []string `json:"byes"`
}

type Game struct {
	Date    string `json:"date"`
	Network string `json:"network"`
	Home    Team   `json:"home"`
	Away    Team   `json:"away"`
	Time    string `json:"time"`
}

type Team struct {
	Name   string `json:"name"`
	Score  string `json:"score"`
	Record string `json:"record"`
	Logo   string `json:"logo"`
}
