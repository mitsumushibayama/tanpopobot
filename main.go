package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/tweetbot/keys"
	"github.com/tweetbot/text"
)

func main() {
	api := keys.GetTwitterApi()
	text := text.ChooseTweet()

	//followerslist := api.GetFollowersIdsAll(nil)

	SearchText := "tanpopobot"
	v := url.Values{}
	v.Set("count", "30")

	x := url.Values{}

	timelinetweets, err := api.GetUserTimeline(v)
	if err != nil {
		panic(err)
	}

	for i, tweet := range timelinetweets {
		if strings.Contains(tweet.FullText, SearchText) {
			tweetid := timelinetweets[i].IdStr
			flag := false

			//2回以上同一のツイートにリプライをしない
			for j := range timelinetweets {
				if timelinetweets[j].InReplyToStatusIdStr == tweetid {
					flag = true
					continue
				}
			}

			if flag {
				continue
			}
			x.Add("in_reply_to_status_id", tweetid)

			reply, err := api.PostTweet(text, x)
			if err != nil {
				panic(err)
			}
			fmt.Println(reply.Text)

		}
	}
}
