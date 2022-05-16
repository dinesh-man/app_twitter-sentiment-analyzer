package src

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/go-gota/gota/dataframe"
	"fmt"
	"encoding/json"
	"os"
	"time"
	"strings"
)

type SearchTweetContents struct {
	Statuses []struct{
		Tweet_Timestamp	string	`json:"created_at"`
		Tweet_Text		string	`json:"full_text"`
		Likes			int		`json:"favorite_count"`
	} `json:"statuses"`
}

func GetTweets(client *twitter.Client, search_query string) {

	fmt.Printf("\nQuery: %v\nSearching...\n", search_query)
	current_date := time.Now().Format("2006-01-02") //YYYY-MM-DD format
	search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{Query: search_query, 
		Lang: "en",
		Until: current_date,
		TweetMode: "extended",
		Filter: "safe",
		Count: 100,})
	if err != nil {
		fmt.Println("Error in GetTweets function!\n", err)
		os.Exit(1)
	}

	fmt.Printf("Twitter API Response: %+v\n", resp.Status)
	bytes, err := json.Marshal(search)
    if err != nil {
		fmt.Println("Error in GetTweets function!\n", err)
		os.Exit(1)
    }
	json_data := string(bytes)
	var output SearchTweetContents
	json.Unmarshal([]byte(json_data), &output)
	/* Uncomment for debugging
	for _, value := range output.Statuses{
		fmt.Printf("\n\nTweeted Date: %#v Tweet Content: %#v Likes: %d", value.Tweet_Timestamp, value.Tweet_Text, value.Likes)
	}*/
	if len(output.Statuses) > 0 {
		temp_df := dataframe.LoadStructs(output.Statuses)
		df := temp_df.Arrange(dataframe.RevSort("Likes")) //Sort by most liked tweets
		file_name := "./Twitter-Search-" + strings.Replace(search_query, " ", "_", -1)+ "-" + time.Now().Format("2006-01-02_15:04:05") + ".csv"
		fmt.Println("Writing search results to file: ", file_name)
		f, err := os.Create(file_name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		df.WriteCSV(f)
		fmt.Println("Done!  :)")
	} else {
		fmt.Println("No results found! :(")
	}
}
