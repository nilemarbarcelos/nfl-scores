package parser

import (
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/nilemarbarcelos/nfl-scores/model"
)

func Parse(season string, weekNumber string) model.Week {
	rootURL := "http://www.nfl.com/scores/" + season + "/REG" + weekNumber
	doc, err := goquery.NewDocument(rootURL)
	if err != nil {
		log.Fatal(err)
	}

	games := []model.Game{}
	var byes []string

	doc.Find(".bye a").Each(func(i int, s *goquery.Selection) {
		byes = append(byes, s.Text())
	})

	doc.Find(".new-score-box-wrapper").Each(func(i int, s *goquery.Selection) {
		date := s.Find(".new-score-box-heading .date").Text()
		network := s.Find(".new-score-box-heading .network").Text()
		home := s.Find(".team-wrapper .home-team .team-data .team-name").Text()
		homeScore := s.Find(".team-wrapper .home-team .team-data .total-score").Text()
		away := s.Find(".team-wrapper .away-team .team-data .team-name").Text()
		awayScore := s.Find(".team-wrapper .away-team .team-data .total-score").Text()
		time := s.Find(".game-center-area .time-left").Text()

		homeTeam := model.Team{
			Name:  home,
			Score: homeScore,
		}

		awayTeam := model.Team{
			Name:  away,
			Score: awayScore,
		}

		games = append(games, model.Game{
			Date:    date,
			Network: network,
			Home:    homeTeam,
			Away:    awayTeam,
			Time:    time,
		})

	})
	week := model.Week{games, byes}
	return week
}
