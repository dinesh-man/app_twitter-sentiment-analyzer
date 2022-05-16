package src

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"fmt"
	"os"
	"strings"
	"encoding/json"
	"time"
)

type UserTweetContents struct {
	Tweet_Timestamp	string	`json:"created_at"`
	Tweet_Text		string	`json:"full_text"`
	Likes			int		`json:"favorite_count"`
}

func GetUserTweets(client *twitter.Client, twitter_handle string, include_retweets bool) {

	fmt.Printf("\nTwitter Handle: %v\nFetching User Timeline...\n", twitter_handle)
	tweets, resp, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams {
		ScreenName: twitter_handle,
		TweetMode: "extended",
		Count: 300,
		})
	
	if err != nil {
		fmt.Println("Error in GetUserTweets function!\n", err)
		os.Exit(1)
	}

	fmt.Printf("Twitter API Response: %+v\n", resp.Status)
	bytes, err := json.Marshal(tweets)
    if err != nil {
        fmt.Println("Error in GetUserTweets function!\n", err)
		os.Exit(1)
    }
	json_data := string(bytes)
	var output []UserTweetContents
	json.Unmarshal([]byte(json_data), &output)
	/* Uncomment for debugging
	for i := range output {
		if !strings.Contains(output[i].Full_Text, "RT"){
			fmt.Printf("\n\nTweet_Timestamp: %v Tweet_Text: %v Likes: %v", output[i].Tweet_Timestamp, output[i].Tweet_Text, output[i].Likes)
		}
	}*/
	if len(output) > 0 {

		FilterReTweets := func (prefix string) func (el series.Element) bool {
			return func (el series.Element) bool {
				if strings.HasPrefix(el.Val().(string), prefix) {
					return false
				}
					return true
			}
		}

		df := dataframe.LoadStructs(output)
		if !include_retweets {
			df = df.Filter(dataframe.F{Colname: "Tweet_Text", Comparator: series.CompFunc, Comparando: FilterReTweets("RT")})
		}
		df = df.Arrange(dataframe.RevSort("Likes")) //Sort by most liked tweets
		file_name := "./Twitter-UserTimeline-" + strings.Replace(twitter_handle, " ", "_", -1)+ "-" + time.Now().Format("2006-01-02_15:04:05") + ".csv"
		fmt.Println("Writing user tweets to file: ", file_name)
		f, err := os.Create(file_name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		df.WriteCSV(f)
		fmt.Println("Done!  :)\n")
	} else {
		fmt.Println("No results found! :(")
	}
}
