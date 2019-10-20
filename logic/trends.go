package logic

import (
	"basic_blog_go/japerk"
	"basic_blog_go/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// TrendLogic defines a struct that will be associated to all operations
// related to Twitter trending topics
type TrendLogic struct {
	DB            utils.DBOrmer
	TwitterClient *twitter.Client
	JaperkClient  *japerk.Client
}

// NewTrendLogic returns an istance of UserLogic struct
func NewTrendLogic() *TrendLogic {
	return &TrendLogic{
		DB:            utils.NewDBWrapper(),
		TwitterClient: newTwitterClient(),
		JaperkClient:  newJaperkClient(),
	}
}

func newTwitterClient() *twitter.Client {
	config := oauth1.NewConfig(
		utils.GetAppConfig("twitter_api_key"),
		utils.GetAppConfig("twitter_api_secret"),
	)
	token := oauth1.NewToken(
		utils.GetAppConfig("twitter_access_token"),
		utils.GetAppConfig("twitter_access_token_secret"),
	)
	httpClient := config.Client(oauth1.NoContext, token)

	return twitter.NewClient(httpClient)
}

func newJaperkClient() *japerk.Client {
	return japerk.NewClient(&http.Client{})
}

// GetTrends returns a list of trending topics based on filters
func (t *TrendLogic) GetTrends(id int64, exclude string) ([]twitter.TrendsList, error) {
	trends, _, err := t.TwitterClient.Trends.Place(id, nil)

	if err != nil {
		return trends, fmt.Errorf("Error retrieving trends: (%v)", err)
	}

	return trends, err
}

// GetTweets returns a list of tweets based on search query
func (t *TrendLogic) GetTweets(q string) (*japerk.Sentiment, error) {
	tweets, _, err := t.TwitterClient.Search.Tweets(&twitter.SearchTweetParams{
		Query: q,
		Lang:  "en",
		Count: 100,
	})

	if err != nil {
		return nil, fmt.Errorf("Error retrieving tweets: [%v]", err)
	}

	statusText := []string{}
	for _, tweet := range tweets.Statuses {
		statusText = append(statusText, tweet.Text)
	}

	statuses := strings.Join(statusText, ".")

	sentiment, err := t.JaperkClient.Sentiment.SentimentAnalysis(statuses)

	return sentiment, err
}
