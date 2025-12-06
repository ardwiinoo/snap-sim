package model

type ValidateRequest struct {
	Amount        int64  `json:"amount"`
	Currency      string `json:"currency"`
	ExternalID    string `json:"external_id"`
	SourceAccount string `json:"source_account"`
	TargetAccount string `json:"target_account"`
	Description   string `json:"description"`
}
