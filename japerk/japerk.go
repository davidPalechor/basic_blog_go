package japerk

import (
	"net/http"

	"basic_blog_go/utils"

	"github.com/dghubble/sling"
)

const japerkAPI = "https://japerk-text-processing.p.rapidapi.com"

// Client is a japerk client to make japerk API request
type Client struct {
	sling     *sling.Sling
	Sentiment *SentimentService
}

func NewClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base(japerkAPI)
	base.Set("x-rapidapi-host", "japerk-text-processing.p.rapidapi.com")
	base.Set("x-rapidapi-key", utils.GetAppConfig("japerk_api_key"))
	return &Client{
		sling:     base,
		Sentiment: newSentimentService(base.New()),
	}
}
