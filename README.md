#APP_TWITTER-SENTIMENT-ANALYZER

This application searches twitter to find out tweets related to specific products/places/restaurants/brands/celebrity/any keyword or specific users' timelines etc. This program will help to do sentiment analysis and find interesting, hilarious facts about a specific topic using tweets (likes/retweets etc.) posted on twitter.

In order to run this app, follow the below mentioned steps -

1. This app calls the Twitter API behind the scene, hence raise a request for a twitter developer account on the developer portal to get access to Twitter API -  https://developer.twitter.com
Once account is created successfully, generate your secret keys on the portal to authenticate against the API.
Following are the secrect keys to be generated -
API_KEY, API_KEY_SECRET, ACCESS_TOKEN, ACCESS_TOKEN_SECRET

2. Set the secret keys as environment variables.
export API_KEY=xxxx
export API_KEY_SECRET=xxxx
export ACCESS_TOKEN=xxxx
export ACCESS_TOKEN_SECRET=xxxx

3. Set GOPATH - 
cd APP_TWITTER-SENTIMENT-ANALYZER
export GOPATH=${PWD}/go

4. Run the app -
For help related to usage run - 
go run main.go -h

Examples -
go run main.go -s <any_search_string> //to do general twitter search
or
go run main.go -t <any_twitter_user_handle> //to explore tweets from specific user timeline
