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
	//pages := api.GetFollowersIdsAll(nil)

	SearchText := "tanpopobot"
	v := url.Values{}
	v.Set("count", "20")

	x := url.Values{}

	timelinetweets, err := api.GetUserTimeline(v)
	if err != nil {
		panic(err)
	}

	for i, tweet := range timelinetweets {
		if strings.Contains(tweet.FullText, SearchText) {
			tweet_id := timelinetweets[i].IdStr
			status := timelinetweets[i]
			x.Add("in_reply_to_status_id", tweet_id)

			reply, err := api.PostTweet(text, x)
			if err != nil {
				panic(err)
			}
			fmt.Println(status)
			fmt.Println(reply.Text)
		}
	}

	//time.Sleep(time.Hour * 24 * 365)
}
