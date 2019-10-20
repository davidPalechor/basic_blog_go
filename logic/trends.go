package logic

import (
	"basic_blog_go/utils"
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// TrendLogic defines a struct that will be associated to all operations
// related to Twitter trending topics
type TrendLogic struct {
	DB utils.DBOrmer
}

// NewTrendLogic returns an istance of UserLogic struct
func NewTrendLogic() *TrendLogic {
	return &TrendLogic{
		DB: utils.NewDBWrapper(),
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

// GetTrends returns a list of trending topics based on filters
func (t *TrendLogic) GetTrends(id int64, exclude string) ([]twitter.TrendsList, error) {
	client := newTwitterClient()
	trends, _, err := client.Trends.Place(id, nil)

	if err != nil {
		return trends, fmt.Errorf("Error retrieving trends: (%v)", err)
	}

	return trends, err
}
