package stack_exchange

// type APITypes interface {
// 	UserInfo
// 	SiteInfo
// 	ArticleInfo
// }

type APIResponse struct {
	Items          interface{} `json:"items"`
	HasMore        bool        `json:"has_more"`
	QuotaMax       int         `json:"quota_max"`
	QuotaRemaining int         `json:"quota_remaining"`
}
