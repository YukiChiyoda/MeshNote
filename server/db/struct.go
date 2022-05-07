package db

type User struct {
	Id       int    `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
}

type Tree struct {
	Id       int    `json:"Id" db:"id"`
	Name     string `json:"Name" db:"name"`
	Type     int    `json:"Type" db:"type"` // 0 => file; -1 => recycled file; 1 => dir
	FileName string `json:"FileName" db:"filename"`
	FileSize int    `json:"FileSize" db:"filesize"`
	Parent   int    `json:"Parent" db:"parent"` // 0 => root path
	Uptime   string `json:"Uptime" db:"uptime"` // MicroUnix
}

type Log struct {
	Id     int    `json:"Id" db:"id"`
	User   string `json:"User" db:"user"`
	Event  string `json:"Event" db:"event"`
	Detail string `json:"Detail" db:"detail"`
}

const (
	Id_Undefine       = 0
	Type_File         = 0
	Type_Recyled_File = -1
	Type_Dir          = 1
	FileName_Dir      = "<dir>"
	FileSize_Empty    = 0
	Parent_Root       = 0
)
