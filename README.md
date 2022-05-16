#APP_TWITTER-SENTIMENT-ANALYZER

This application calls Twitter API to find out tweets related to specific products/places/restaurants/brands/celebrity/any search string or specific users' timelines etc. This program will help to do sentiment analysis and find interesting, hilarious facts about a specific topic using tweets posted on twitter (likes, retweets etc.). The search results are saved in a csv file, which can be later analyzed by the user using excel.

In order to run this app, follow the below mentioned steps -

This app calls the Twitter API behind the scene, hence raise a request for a twitter developer account on the developer portal to get access to Twitter API - https://developer.twitter.com Once account is created successfully, generate your secret keys on the portal to authenticate against the API. Following are the secret keys to be generated - API_KEY, API_KEY_SECRET, ACCESS_TOKEN, ACCESS_TOKEN_SECRET

Set the secret keys as environment variables.

`export API_KEY=xxxx`
`export API_KEY_SECRET=xxxx` 
`export ACCESS_TOKEN=xxxx` 
`export ACCESS_TOKEN_SECRET=xxxx`

Set GOPATH -
`cd APP_TWITTER-SENTIMENT-ANALYZER`
`export GOPATH=${PWD}/go`

Run the app -

`go run main.go -s <any_search_string>` //to do general twitter search 
`go run main.go -t <any_twitter_user_handle>` //to explore tweets from specific user timeline
`go run main.go -h` //to see usage help

By-default retweets are not included in the search result. To include those as well use the `-rt` option to set it to true.
For example -
`go run main.go -t <any_twitter_user_handle> -rt`

The search results are saved in a csv file under the project root directory.
