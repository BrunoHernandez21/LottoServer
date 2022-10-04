package youtube

type Snippet struct {
	PublishedAt          string     `json:"publishedAt"`
	ChannelID            string     `json:"channelId"`
	Title                string     `json:"title"`
	Description          string     `json:"description"`
	Thumbnails           Thumbnails `json:"thumbnails"`
	ChannelTitle         string     `json:"channelTitle"`
	Tags                 []string   `json:"tags"`
	CategoryID           string     `json:"categoryId"`
	LiveBroadcastContent string     `json:"liveBroadcastContent"`
	Localized            Localized  `json:"localized"`
}

type Localized struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Thumbnails struct {
	Default  Default `json:"default"`
	Medium   Default `json:"medium"`
	High     Default `json:"high"`
	Standard Default `json:"standard"`
	Maxres   Default `json:"maxres"`
}

type Default struct {
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}
