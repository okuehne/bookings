package models

import "github.com/okuehne/bookings/internal/forms"

// TemplateData holds data sent from
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]any
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.From
}
