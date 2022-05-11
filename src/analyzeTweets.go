package src

import (
	"github.com/dghubble/go-twitter/twitter"
	"fmt"
	"strings"
	"encoding/json"
)

type SearchTweetContents struct {
	Statuses []struct{
		Created_At	string `json:"created_at"`
		Full_Text	string `json:"full_text"`
	} `json:"statuses"`
}

func GetTweets(client *twitter.Client, search_query string) {

	search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{Query: search_query, 
		Lang: "en", 
		Since: "2022-01-01",
		TweetMode: "extended",
		Filter: "safe",
		Count: 100,})

	if err != nil {
		fmt.Println("Error in GetTweets function!\n", err)
	}

	fmt.Printf("\n\nResponse: %+v\n\n", resp.Status)
	bytes, err := json.Marshal(search)
    if err != nil {
		fmt.Println("Error in GetTweets function!\n", err)
    }
	json_data := string(bytes)
	var output SearchTweetContents
	json.Unmarshal([]byte(json_data), &output)
	for _, value := range output.Statuses{
		if !strings.Contains(value.Full_Text, "RT"){
			fmt.Printf("\n\nTweeted Date: %v Tweet Content: %v", value.Created_At, value.Full_Text)
		}
	}

}

