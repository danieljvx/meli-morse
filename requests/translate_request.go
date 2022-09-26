package requests

type TranslateRequest struct {
	Text string     `json:"text,omitempty" validate:"required"`
}
