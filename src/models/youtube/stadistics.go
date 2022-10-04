package youtube

type YtResponse struct {
	Kind     string   `json:"kind"`
	Etag     string   `json:"etag"`
	Items    []Item   `json:"items"`
	PageInfo PageInfo `json:"pageInfo"`
}

type Item struct {
	Kind       string      `json:"kind"`
	Etag       string      `json:"etag"`
	ID         string      `json:"id"`
	Statistics *Statistics `json:"statistics"`
	Snippet    *Snippet    `json:"snippet"`
}

type Statistics struct {
	ViewCount     string `json:"viewCount"`
	LikeCount     string `json:"likeCount"`
	FavoriteCount string `json:"favoriteCount"`
	CommentCount  string `json:"commentCount"`
}

type PageInfo struct {
	TotalResults   int64 `json:"totalResults"`
	ResultsPerPage int64 `json:"resultsPerPage"`
}
