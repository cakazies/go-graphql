package types

type (
	Products struct {
		IdPro    string `json:"ID_PRO"`
		ProName  string `json:"PRO_NAME"`
		QteStock int    `json:"QTE_IN_STOCK"`
		ProImg   string `json:"PRO_IMAGE"`
	}
	Users struct {
		Id         string `json:"id,omitempty"`
		Name       string `json:"name"`
		Age        int    `json:"age"`
		Profession string `json:"profession"`
		Friendly   string `json:"friendly"`
	}
)
