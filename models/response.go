package models

type (
	Response struct {
		Data        interface{} `json:"responseData"`
		RC          string      `json:"responseCode"`
		Description string      `json:"responseDescription"`
		Error       error       `json:"error,omitempty"`
	}
)
