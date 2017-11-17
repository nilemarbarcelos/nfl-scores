package parser

import (
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/nilemarbarcelos/nfl-scores/model"
)

func Parse(season string, week string) []model.Game {
	rootURL := "http://www.nfl.com/scores/" + season + "/REG" + week
	doc, err := goquery.NewDocument(rootURL)
	if err != nil {
		log.Fatal(err)
	}
	games := []model.Game{}

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
	return games
}

//   <div class="new-score-box-wrapper" id="scorebox-2017110900">
//   <div class="new-score-box-heading" id="yui_3_10_3_2_1510613392638_21456">
// 	  <p id="yui_3_10_3_2_1510613392638_21455">
// 		  <span class="date" title="Date Aired" id="yui_3_10_3_2_1510613392638_21454">Thu, Nov 9</span>
// 		  <span class="network" title="Aired on Network">nbc, NFLN</span>
// 		  <span class="quarters-total" title="Score Key by Quarter">
// 			  <span class="total-score">TOTAL</span>
// 			  <span class="quarters-score">
// 				  <span>1</span><span>2</span><span>3</span><span>4</span><span>OT</span>
// 			  </span>
// 		  </span>
// 	  </p>
//   </div>
//   <div class="new-score-box" id="yui_3_10_3_2_1510613392638_2271">
//   <div class="team-wrapper" id="yui_3_10_3_2_1510613392638_2270">
// 		  <div class="away-team" id="yui_3_10_3_2_1510613392638_2269">
// 			  <a href="/teams/profile?team=SEA" onclick="s_objectID=&quot;http://www.nfl.com/teams/profile?team=SEA_1&quot;;return this.s_oc?this.s_oc(e):true"><img src="http://i.nflcdn.com/static/site/7.5/img/logos/teams-matte-80x53/SEA.png" class="team-logo" alt=""></a>
// 			  <div class="team-data" id="yui_3_10_3_2_1510613392638_2268">
// 				  <div class="team-info" id="yui_3_10_3_2_1510613392638_2267">
// 					  <p class="team-record">

// 							  <a href="/teams/profile?team=SEA" onclick="s_objectID=&quot;http://www.nfl.com/teams/profile?team=SEA_2&quot;;return this.s_oc?this.s_oc(e):true">(6-3-0)</a>

// 					  </p>
// 					  <p class="team-name" id="yui_3_10_3_2_1510613392638_2266"><a href="/teams/profile?team=SEA" onclick="s_objectID=&quot;http://www.nfl.com/teams/profile?team=SEA_3&quot;;return this.s_oc?this.s_oc(e):true" id="yui_3_10_3_2_1510613392638_2265">Seahawks</a></p>
// 				  </div>
// 				  <p class="total-score">22</p>
// 				  <p class="quarters-score"><span class="first-qt">7</span><span class="second-qt">8</span><span class="third-qt">0</span><span class="fourth-qt">7</span><span class="ot-qt"></span></p>

// 			  </div>
// 		  </div>
// 	  </div>
// 	  <div class="team-wrapper">
// 		  <div class="home-team">
// 			  <a href="/teams/profile?team=ARI" onclick="s_objectID=&quot;http://www.nfl.com/teams/profile?team=ARI_1&quot;;return this.s_oc?this.s_oc(e):true"><img src="http://i.nflcdn.com/static/site/7.5/img/logos/teams-matte-80x53/ARI.png" class="team-logo" alt=""></a>
// 			  <div class="team-data">
// 				  <div class="team-info">
// 					  <p class="team-record">

// 							  <a href="/teams/profile?team=ARI" onclick="s_objectID=&quot;http://www.nfl.com/teams/profile?team=ARI_2&quot;;return this.s_oc?this.s_oc(e):true">(4-5-0)</a>

// 					  </p>
// 					  <p class="team-name"><a href="/teams/profile?team=ARI" onclick="s_objectID=&quot;http://www.nfl.com/teams/profile?team=ARI_3&quot;;return this.s_oc?this.s_oc(e):true">Cardinals</a></p>
// 				  </div>
// 				  <p class="total-score">16</p>
// 				  <p class="quarters-score"><span class="first-qt">0</span><span class="second-qt">7</span><span class="third-qt">3</span><span class="fourth-qt">6</span><span class="ot-qt"></span></p>
// 			  </div>
// 		  </div>
// 	  </div>
// 	  <div class="game-center-area">
// 		  <a href="/gamecenter/2017110900/2017/REG10/seahawks@cardinals" class="game-center-link" onclick="s_objectID=&quot;http://www.nfl.com/gamecenter/2017110900/2017/REG10/seahawks@cardinals_1&quot;;return this.s_oc?this.s_oc(e):true"><img src="http://i.nflcdn.com/static/site/7.5/img/scores/game-center-active.png" alt=""></a>
// 		  <p><span class="time-left">FINAL </span></p>
// 		  <div style="clear:right"></div>
// 		  <div class="big-plays"><span class="title">BIG PLAYS</span><span class="big-plays-count">11</span></div>
// 		  <div style="clear:both"></div>

// 			  <div class="highlight-link">
// 				  <a href="/gamecenter/2017110900/2017/REG10/seahawks@cardinals#menu=highlights|contentId:0ap3000000874863&amp;tab=analyze" onclick="s_objectID=&quot;http://www.nfl.com/gamecenter/2017110900/2017/REG10/seahawks@cardinals#menu=highlights|contentId:_1&quot;;return this.s_oc?this.s_oc(e):true">
// 					  <div class="highlight-button"><div class="highlight-arrow"></div></div>
// 					  <p>Seahawks vs. Cardinals highlights | Week 10</p>
// 				  </a>
// 			  </div>

// 	  </div>
// 	  <div class="products-social">
// 		  <p>
// 			  <span class="listen"><a href="/gamepass?icampaign=GP_Scores" onclick="s_objectID=&quot;http://www.nfl.com/gamepass?icampaign=GP_Scores_2&quot;;return this.s_oc?this.s_oc(e):true"><img src="http://i.nflcdn.com/static/site/7.5/img/scores/camera.png" alt=""> FULL GAME</a></span>
// 		  </p>
// 	  </div>
//   </div>
// </div>
