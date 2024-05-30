package dto

type (
	Default struct {
		Ok      bool   `json:"ok"`
		Message string `json:"message"`
		Data    any    `json:"data,omitempty"`
	}

	ErrorData struct {
		Error string `json:"error"`
	}
)
