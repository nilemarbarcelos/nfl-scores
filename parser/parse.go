package parser

import (
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/nilemarbarcelos/nfl-scores/model"
)

func extractImages(s *goquery.Selection) string {
	str, exists := s.Attr("src")
	if exists {
		return str
	}
	return ""
}

func extractRecords(s string) string {
	var rgx = regexp.MustCompile(`\((.*?)\)`)
	rs := rgx.FindStringSubmatch(s)
	return rs[1]
}

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
		homeLogo := extractImages(s.Find(".team-wrapper .home-team .team-logo"))
		homeRecord := extractRecords(s.Find(".team-wrapper .home-team .team-data .team-info .team-record").Text())

		away := s.Find(".team-wrapper .away-team .team-data .team-name").Text()
		awayScore := s.Find(".team-wrapper .away-team .team-data .total-score").Text()
		awayLogo := extractImages(s.Find(".team-wrapper .away-team .team-logo"))
		awayRecord := extractRecords(s.Find(".team-wrapper .away-team .team-data .team-info .team-record").Text())

		time := strings.TrimSpace(s.Find(".game-center-area .time-left").Text())

		homeTeam := model.Team{
			Name:   home,
			Score:  homeScore,
			Record: homeRecord,
			Logo:   homeLogo,
		}

		awayTeam := model.Team{
			Name:   away,
			Score:  awayScore,
			Record: awayRecord,
			Logo:   awayLogo,
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

// <div class="new-score-box-wrapper">
// <div class="new-score-box-heading">
// 	<p>
// 		<span class="date" title="Date Airing">Sun, Dec 10</span>

// 		<span class="network" title="Airing on Network Channel">fox</span>
// 		<span class="tickets-link"><a href="/tickets/tampa-bay-buccaneers?icampaign=Scores_Tickets">GET TICKETS</a></span>
// 	</p>
// </div>
// <div class="new-score-box">
// 	<div class="team-wrapper">
// 		<div class="away-team">
// 			<a href="/teams/profile?team=DET"><img src="http://i.nflcdn.com/static/site/7.5/img/logos/teams-matte-80x53/DET.png" class="team-logo" alt=""></a>
// 			<div class="team-data">
// 				<div class="team-info">
// 					<p class="team-record">(6-6-0)</p>
// 					<p class="team-name"><a href="/teams/profile?team=DET">Lions</a></p>
// 				</div>
// 				<p class="total-score">--</p>
// 			</div>
// 		</div>
// 	</div>
// 	<div class="team-wrapper">
// 		<div class="home-team">
// 			<a href="/teams/profile?team=TB"><img src="http://i.nflcdn.com/static/site/7.5/img/logos/teams-matte-80x53/TB.png" class="team-logo" alt=""></a>
// 			<div class="team-data">
// 				<div class="team-info">
// 					<p class="team-record">(4-8-0)</p>
// 					<p class="team-name"><a href="/teams/profile?team=TB">Buccaneers</a></p>
// 				</div>
// 				<p class="total-score">--</p>
// 			</div>
// 		</div>
// 	</div>
// 	<div class="game-center-area">

// 		<a href="/gamecenter/2017121007/2017/REG14/lions@buccaneers" class="game-center-link"><img src="http://i.nflcdn.com/static/site/7.5/img/scores/game-center-regular.png" alt="Game Center"></a>
// 		<p><span class="time-left">1:00 PM  ET</span></p>
// 		<div class="comments" style="clear:both"><a href="/gamecenter/2017121007/2017/REG14/lions@buccaneers#tab=discuss"><fb:comments-count href="http://www.nfl.com/gamecenter/2017121007/2017/REG14/lions@buccaneers"></fb:comments-count></a></div>

// 	</div>
// 	<div class="products-social">
// 		<p>
// 			<span class="listen"><a href="/gamepass?icampaign=GP_Scores"><img src="http://i.nflcdn.com/static/site/7.5/img/scores/camera.png" alt=""> WATCH LIVE</a></span>
// 		</p>
// 	</div>
// </div>
// </div>
