package src

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/jdkato/prose/v2"
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

func GetTweets(client *twitter.Client, search_query string, include_retweets bool) {

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
	
	if len(output.Statuses) > 0 {

		FilterReTweets := func (prefix string) func (el series.Element) bool {
			return func (el series.Element) bool {
				if strings.HasPrefix(el.Val().(string), prefix) {
					return false
				}
					return true
			}
		}

		df := dataframe.LoadStructs(output.Statuses)
		if !include_retweets {
			df = df.Filter(dataframe.F{Colname: "Tweet_Text", Comparator: series.CompFunc, Comparando: FilterReTweets("RT")})
		}
		df = df.Arrange(dataframe.RevSort("Likes")) //Sort by most liked tweets
		file_name := "./Twitter-Search_" + strings.Replace(search_query, " ", "_", -1)+ "-" + time.Now().Format("2006-01-02_15:04:05") + ".csv"
		fmt.Println("Writing search results to file: ", file_name)
		f, err := os.Create(file_name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		df.WriteCSV(f)
		fmt.Println("File created successfully! :)")
		fmt.Println("Analyzing tweets...")
		var tweet_keywords []string
		for _, value := range output.Statuses {
			//fmt.Printf("\n\nTweeted Date: %#v Tweet Content: %#v Likes: %d", value.Tweet_Timestamp, value.Tweet_Text, value.Likes)
			if include_retweets == false && strings.HasPrefix(value.Tweet_Text, "RT"){
				continue
			} else {
				doc, err := prose.NewDocument(strings.Replace(value.Tweet_Text, "\n", "", -1))
				if err != nil {
					fmt.Println(err)
				}

				for _, tok := range doc.Tokens() {
					if tok.Tag == "JJ" || tok.Tag == "VBG" || tok.Tag == "VB" || tok.Tag == "NN" || tok.Tag == "NNS" || tok.Label == "B-GPE" || tok.Label == "B-PERSON" {
						tweet_keywords = append(tweet_keywords, tok.Text)
					}
				}
			}
		}
		df_tk := dataframe.New(series.New(tweet_keywords, series.String, "Tweet_Keywords"))
		groups := df_tk.GroupBy("Tweet_Keywords")
		aggre := groups.Aggregation([]dataframe.AggregationType{dataframe.Aggregation_COUNT}, []string{"Tweet_Keywords"})
		aggre = aggre.Arrange(dataframe.RevSort("Tweet_Keywords_COUNT")) //Sort by most used words
		//fmt.Printf("\n\n%v\n\n",df)
		//fmt.Printf("\n\n%v\n\n",aggre)
		file_name = "./Twitter-Search-analysis_" + strings.Replace(search_query, " ", "_", -1)+ "-" + time.Now().Format("2006-01-02_15:04:05") + ".csv"
		fmt.Println("Writing analysis to file: ", file_name)
		f, err = os.Create(file_name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		aggre.WriteCSV(f)
		fmt.Println("File created successfully! :)")
		fmt.Println("Done!  :)")
	} else {
		fmt.Println("No results found! :(")
	}
}
