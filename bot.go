package main

import (
	"bufio"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"os"
)

const ConsumerKey = "CONSUMER_KEY"
const ConsumerKeySecret = "CONSUMER_SECRET"
const AccessToken = "ACCESS_TOKEN"
const AccessSecret = "ACCESS_SECRET"

func main() {
	config := oauth1.NewConfig(ConsumerKey, ConsumerKeySecret)
	token := oauth1.NewToken(AccessToken, AccessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)
	// Twitter client
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		IncludeEmail: twitter.Bool(true),
	}
	user, _, _ := client.Accounts.VerifyCredentials(verifyParams)
	fmt.Printf("User's Name:%+v\n", user.ScreenName)

	scanner := bufio.NewScanner(os.Stdin)
	var searchText string
	fmt.Print("Enter your search text: ")
	scanner.Scan()
	searchText = scanner.Text()

	var searchCount int

	fmt.Print("Enter how many tweets you want to retweet: ")
	_, err := fmt.Scanf("%d", &searchCount)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// Search tweets to retweet
	searchParams := &twitter.SearchTweetParams{
		Query: "%23" + searchText,
		Count: searchCount,
	}
	searchResult, _, _ := client.Search.Tweets(searchParams)

	if len(searchResult.Statuses) == 0 {
		fmt.Println("No Query Results! Exiting")
		os.Exit(0)
	}
	for _, tweet := range searchResult.Statuses {
		tweetID := tweet.ID
		client.Statuses.Retweet(tweetID, &twitter.StatusRetweetParams{})

		fmt.Printf("RETWEETED: %+v\n", tweet.Text)
	}

}
