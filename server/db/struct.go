package db

type Tree struct {
	Id       int    `json:"Id"`
	Name     string `json:"Name"`
	Type     int    `json:"Type"` // 0 => file; -1 => recycled file; 1 => dir
	FileName string `json:"FileName"`
	FileSize int    `json:"FileSize"`
	Parent   int    `json:"Parent"` // 0 => root path
	Uptime   string `json:"Uptime"` //MicroUnix
}
