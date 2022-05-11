package analyzeUserTweets

import (
	"github.com/dghubble/go-twitter/twitter"
	"fmt"
	"strings"
	"encoding/json"
)

type UserTweetContents struct {
	Created_At	string `json:"created_at"`
	Full_Text	string `json:"full_text"`
}

func GetUserTweets(client *twitter.Client, twitter_handle string) {

	tweets, resp, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams {
		ScreenName: twitter_handle,
		TweetMode: "extended",
		Count: 100,
		})
	
	if err != nil {
		fmt.Println("Error in GetUserTweets function!\n", err)
	}

	fmt.Printf("\n\nResponse: %+v\n\n", resp.Status)
	bytes, err := json.Marshal(tweets)
    if err != nil {
        fmt.Println("Error in GetUserTweets function!\n", err)
    }
	json_data := string(bytes)
	var output []UserTweetContents
	json.Unmarshal([]byte(json_data), &output)
	for i := range output {
		if !strings.Contains(output[i].Full_Text, "RT"){
			fmt.Printf("\n\nTweeted Date: %v Tweet Content: %v", output[i].Created_At, output[i].Full_Text)
		}
	}

}
