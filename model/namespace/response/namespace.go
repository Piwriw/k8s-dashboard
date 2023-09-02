package response

type Namespace struct {
	Name            string `json:"name"`
	CreateTimestamp int64  `json:"createTimestamp"`
	Status          string `json:"status"`
}
