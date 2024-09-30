package synchronization

type Host struct {
	Identifier string `json:"identifier"`
	Address    string `json:"address"`
	Port       int    `json:"port"`
}

