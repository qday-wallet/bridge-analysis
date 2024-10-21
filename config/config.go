package config

type Log struct {
	Path  string `json:"Path" gorm:"column:Path"`
	Delay int    `json:"Delay" gorm:"column:Delay"`
}

type DB struct {
	User     string `json:"User" gorm:"column:User"`
	Port     int    `json:"Port" gorm:"column:Port"`
	DbName   string `json:"DbName" gorm:"column:DbName"`
	Addr     string `json:"Addr" gorm:"column:Addr"`
	Password string `json:"Password" gorm:"column:Password"`
}

type Config struct {
	Root string `json:"Root" gorm:"column:Root"`
	Log  *Log   `json:"Log" gorm:"column:Log"`
	DB   *DB    `json:"DB" gorm:"column:DB"`
	Port int    `json:"Port" gorm:"column:Port"`
}
