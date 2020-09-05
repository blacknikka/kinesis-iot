package entities

// Record is レコードデータ
type Record struct {
	Type    string `json:"type"`
	Log     string `json:"log"`
	Action1 bool   `json:"action_1"`
}
