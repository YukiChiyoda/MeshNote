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

type Diary struct {
	Id     int    `json:"Id" db:"id"`
	User   int `json:"User" db:"user"`
	Year  int `json:"Year" db:"year"`
	Month  int `json:"Month" db:"month"`
	Day  int `json:"Day" db:"day"`
	Count  int `json:"Count" db:"count"`
}

const (
	ID_UNDEFINE       = 0
	TYPE_FILE         = 0
	TYPE_RECYLED_FILE = -1
	TYPE_DIR          = 1
	FILENAME_DIR      = "<dir>"
	FILESIZE_EMPTY    = 0
	PARENT_ROOT       = 0
	DIARY_EMPTY       = 0

	SPECIAL_INVALID_ID = -888
	SPECIAL_ERROR_ID   = -999
	SPECIAL_NULL_NAME  = "null"

	DEV_USER_ID = 0
)
