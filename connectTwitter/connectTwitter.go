package connectTwitter

import (
"github.com/dghubble/go-twitter/twitter"
"github.com/dghubble/oauth1"
"fmt"
)

type Credentials struct {
    ApiKey              string
    ApikeySecret        string
    AccessToken         string
    AccessTokenSecret   string
}

func GetClient(creds *Credentials) (*twitter.Client, error){
	api := oauth1.NewConfig(creds.ApiKey, creds.ApikeySecret)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)
	http_client := api.Client(oauth1.NoContext, token)
	client := twitter.NewClient(http_client)

	verifyParams := &twitter.AccountVerifyParams{
        SkipStatus:   twitter.Bool(true),
        IncludeEmail: twitter.Bool(true),
    }

	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
    if err != nil {
        fmt.Println("Error in GetClient function!!")
		return nil, err
    }else{
	fmt.Printf("Authentication Successfull.\nUser Name: %+v\nTwitter Handle: %+v", user.Name, user.ScreenName)
	return client, nil}
}

