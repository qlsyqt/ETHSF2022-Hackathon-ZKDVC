package response

type FetchBadgeVO struct {
	DataCategory  string `json:"dataCategory"`
	Badge         string `json:"badge"`
	BadgeContract string `json:"badgeContract"`
}
