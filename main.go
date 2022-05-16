package main

import (
	"github.com/dinesh-man/app_twitter-sentiment-analyzer/auth"
	"github.com/dinesh-man/app_twitter-sentiment-analyzer/src"
	"fmt"
	"os"
	"flag"
	"strings"
)

func main() {

	//Check environment variables
	env_vars := [4]string{"API_KEY", "API_KEY_SECRET", "ACCESS_TOKEN", "ACCESS_TOKEN_SECRET"}
	env_check := true
	for i:=0; i<len(env_vars); i++ {
		if os.Getenv(env_vars[i]) == "" {
			fmt.Println("Environment variable not found -> ", env_vars[i])
			env_check = false
		}
	}
	if !env_check {
		os.Exit(1)
	}

	//command-line inputs
	search_string := flag.String("s", "", "Use this optional flag to search twitter with a search string.\nInput search string in double quotes.")
	twitter_user := flag.String("t", "", "Use this optional flag to search timeline of a specific twitter user.\nInput twitter handle of the user in double quotes.")
	flag.Parse()
	if flag.NFlag() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println("\n**** Twitter Sentiment Analysis Program ****\n")
	//Authentication using Twitter API keys
	creds := auth.Credentials {
		ApiKey: os.Getenv("API_KEY"),
		ApikeySecret: os.Getenv("API_KEY_SECRET"),
		AccessToken: os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
	}

	client, err := auth.GetClient(&creds)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *search_string != ""{
		src.GetTweets(client, strings.Trim(*search_string, " "))
	} 
	if *twitter_user != "" {
		src.GetUserTweets(client, strings.Trim(*twitter_user, " "))
	}
	
}
