package stack_exchange

type API_SitesResponse struct {
	Items          []SiteInfo `json:"items"`
	HasMore        bool       `json:"has_more"`
	QuotaMax       int        `json:"quota_max"`
	QuotaRemaining int        `json:"quota_remaining"`
}

type SiteInfo struct {
	Name    string `json:"name"`
	Slug    string `json:"api_site_parameter"`
	SiteURL string `json:"site_url"`
}
