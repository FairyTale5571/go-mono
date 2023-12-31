package chast

type (
	Client struct {
		Inn  string `json:"inn"`
		Name Name   `json:"name"`
	}
	Name struct {
		Last   string `json:"last"`
		First  string `json:"first"`
		Middle string `json:"middle"`
	}
)
