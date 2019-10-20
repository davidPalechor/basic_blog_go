package japerk

import (
	"github.com/dghubble/sling"
)

type SentimentBody struct {
	Language string `url:"language"`
	Text     string `url:"text"`
}

type SentimentService struct {
	sling *sling.Sling
}

func newSentimentService(sling *sling.Sling) *SentimentService {
	return &SentimentService{
		sling: sling.Path("/"),
	}
}

func (s *SentimentService) SentimentAnalysis(text string) (*Sentiment, error) {
	var sentiment Sentiment
	body := &SentimentBody{
		Language: "english",
		Text:     text,
	}
	_, err := s.sling.New().Post("sentiment/").BodyForm(body).ReceiveSuccess(&sentiment)
	return &sentiment, err
}
