package stack_exchange

type SiteInfo struct {
	Name    string `json:"name"`
	Slug    string `json:"api_site_parameter"`
	SiteURL string `json:"site_url"`
}
