package requests

type Item struct {
	ProductionOrder      int   `json:"ProductionOrder"`
	ProductionOrderItem  int   `json:"ProductionOrderItem"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}
