package pjlink

// Note that both human readable and raw requests use this struct
type PJRequest struct {
	Address   string `json:"address"`
	Port      string `json:"port"`
	Class     string `json:"class"`
	Password  string `json:"password"`
	Command   string `json:"command"`
	Parameter string `json:"parameter"`
}

type PJResponse struct {
	Class    string   `json:"class"`
	Command  string   `json:"command"`
	Response []string `json:"response"`
}
