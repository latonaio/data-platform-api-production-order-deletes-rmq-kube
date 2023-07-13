package requests

type Header struct {
	ProductionOrder     int   `json:"ProductionOrder"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}
