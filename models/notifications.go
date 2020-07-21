package models

type Event struct {
	Kind string                 `json:"kind"`
	Info map[string]interface{} `json:"info"`
}
