package stack_exchange

type Tag string

type ArticleInfo struct {
	Tags         []Tag    `json:"tags"`
	Owner        UserInfo `json:"owner"`
	ViewCount    int      `json:"view_count"`
	Score        int      `json:"score"`
	CreationDate uint64   `json:"creation_date"`
	ArticleId    uint64   `json:"article_id"`
	ArticleType  string   `json:"article_type"`
	Link         string   `json:"link"`
	Title        string   `json:"title"`
}
