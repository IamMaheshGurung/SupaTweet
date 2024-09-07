package database

type Tweet struct {
	ID      int    `json:"id"`
	Content string `json:"name"`
	UserID  uint   `json:"description"`
}

var tweets = []*Tweet{
	{
		ID:      1,
		Content: "I am really happy to announce my retirement",
		UserID:  1,
	},
}

func GetTweets() []*Tweet {
	return tweets
}

func PostTweets(tweet *Tweet) {
	tweets = append(tweets, tweet)
}

func setTweet(tweetmore []*Tweet) {
	tweets = tweetmore
}
