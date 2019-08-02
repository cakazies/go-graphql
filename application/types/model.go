package types

type (
	Users struct {
		Id         string `json:"id,omitempty"`
		Name       string `json:"name"`
		Age        int    `json:"age"`
		Profession string `json:"profession"`
		Friendly   string `json:"friendly"`
	}
)
