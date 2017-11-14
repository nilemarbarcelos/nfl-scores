package model

type Game struct {
	Date 	string	`json:"date"`
	Network string 	`json:"network"`
	Home 	string 	`json:"home"`
	HomeScore 	string 	`json:"homeScore"`
	Away 	string 	`json:"away"`
	AwayScore 	string 	`json:"awayScore"`
}