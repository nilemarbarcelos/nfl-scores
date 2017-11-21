[![Build Status](https://travis-ci.org/nilemarbarcelos/nfl-scores.svg?branch=master)](https://travis-ci.org/nilemarbarcelos/nfl-scores)
[![Go Report Card](https://goreportcard.com/badge/github.com/nilemarbarcelos/nfl-scores)](https://goreportcard.com/report/github.com/nilemarbarcelos/nfl-scores)
[![GoDoc](https://godoc.org/github.com/nilemarbarcelos/nfl-scores?status.svg)](https://godoc.org/github.com/nilemarbarcelos/nfl-scores)

### Request
* https://scores-nfl.herokuapp.com/nfl-scores/{season}/{week}

### Response
Request: https://scores-nfl.herokuapp.com/nfl-scores/2017/11
```json
{  
   "games":[  
      {  
         "date":"Mon, Nov 20",
         "network":"espn",
         "home":{  
            "name":"Seahawks",
            "score":"--"
         },
         "away":{  
            "name":"Falcons",
            "score":"--"
         },
         "time":"8:30 PM  ET"
      },
      {  
         "date":"Thu, Nov 16",
         "network":"nbc, NFLN",
         "home":{  
            "name":"Steelers",
            "score":"40"
         },
         "away":{  
            "name":"Titans",
            "score":"17"
         },
         "time":"FINAL"
      },
      {  
         "date":"Sun, Nov 19",
         "network":"fox",
         "home":{  
            "name":"Bears",
            "score":"24"
         },
         "away":{  
            "name":"Lions",
            "score":"27"
         },
         "time":"FINAL"
      },
      {  
         "date":"Sun, Nov 19",
         "network":"cbs",
         "home":{  
            "name":"Browns",
            "score":"7"
         },
         "away":{  
            "name":"Jaguars",
            "score":"19"
         },
         "time":"FINAL"
      },
      {  
         "date":"Sun, Nov 19",
         "network":"cbs",
         "home":{  
            "name":"Packers",
            "score":"0"
         },
         "away":{  
            "name":"Ravens",
            "score":"23"
         },
         "time":"FINAL"
      },
      {  
         "date":"Sun, Nov 19",
         "network":"fox",
         "home":{  
            "name":"Texans",
            "score":"31"
         },
         "away":{  
            "name":"Cardinals",
            "score":"21"
         },
         "time":"FINAL"
      },
      {  
         "date":"Sun, Nov 19",
         "network":"fox",
         "home":{  
            "name":"Dolphins",
            "score":"20"
         },
         "away":{  
            "name":"Buccaneers",
            "score":"30"
         },
         "time":"FINAL"
      },
      {  
         "date":"Sun, Nov 19",
         "network":"fox",
         "home":{  
            "name":"Vikings",
            "score":"24"
         },
         "away":{  
            "name":"Rams",
            "score":"7"
         },
         "time":"FINAL"
      },
      {  
         "date":"Sun, Nov 19",
         "network":"cbs",
         "home":{  
            "name":"Giants",
            "score":"12"
         },
         "away":{  
            "name":"Chiefs",
            "score":"9"
         },
         "time":"FINAL OT"
      },
      {  
         "date":"Sun, Nov 19",
         "network":"fox",
         "home":{  
            "name":"Saints",
            "score":"34"
         },
         "away":{  
            "name":"Redskins",
            "score":"31"
         },
         "time":"FINAL OT"
      },
      {  
         "date":"Sun, Nov 19",
         "network":"fox",
         "home":{  
            "name":"Chargers",
            "score":"54"
         },
         "away":{  
            "name":"Bills",
            "score":"24"
         },
         "time":"FINAL"
      },
      {  
         "date":"Sun, Nov 19",
         "network":"cbs",
         "home":{  
            "name":"Broncos",
            "score":"17"
         },
         "away":{  
            "name":"Bengals",
            "score":"20"
         },
         "time":"FINAL"
      },
      {  
         "date":"Sun, Nov 19",
         "network":"cbs",
         "home":{  
            "name":"Raiders",
            "score":"8"
         },
         "away":{  
            "name":"Patriots",
            "score":"33"
         },
         "time":"FINAL"
      },
      {  
         "date":"Sun, Nov 19",
         "network":"nbc",
         "home":{  
            "name":"Cowboys",
            "score":"9"
         },
         "away":{  
            "name":"Eagles",
            "score":"37"
         },
         "time":"FINAL"
      }
   ],
   "byes":[  
      "Jets",
      "Colts",
      "49ers",
      "Panthers"
   ]
}
```