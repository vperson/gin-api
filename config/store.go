package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	MySQLDbTp      = "mysql"
	SQLiteDbTp     = "sqlite"
	PostGreSQLDbTp = "postgres"
)

type DBConfig struct {
	Enable     bool   `yaml:"enable"`
	Tp         string `yaml:"tp"` // mysql
	Debug      bool   `yaml:"debug"`
	SqliteFile string `yaml:"sqliteFile"` // sqlite 属性
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Database   string `yaml:"database"`
	Extend     string `yaml:"extend"`    // mysql 属性
	SslMode    string `yaml:"sslMode"`   // pg 属性
	TimeZone   string `yaml:"timeZone"`  // pg 属性
	Charset    string `yaml:"charset"`   // mysql 属性
	Collation  string `yaml:"collation"` // mysql 属性
}

func (d *DBConfig) SetDefault() {
	d.Enable = false
	d.Tp = MySQLDbTp
	d.Database = "gin_api"
	d.Host = "127.0.0.1"
	d.Port = "3306"
	d.Debug = false
	d.Username = "root"
	d.Extend = "charset=utf8mb4&parseTime=True&loc=Local"
	d.SslMode = "disable"
	d.TimeZone = "Asia/Shanghai"
	d.Charset = "utf8mb4"
	d.Collation = "utf8mb4_unicode_ci"
}

func (d *DBConfig) CreateGormDialector() (gorm.Dialector, error) {
	switch d.Tp {
	case MySQLDbTp:
		return d.mysql(), nil
	case SQLiteDbTp:
		return d.sqlite(), nil
	case PostGreSQLDbTp:
		return d.postgres(), nil
	}
	return nil, fmt.Errorf("db type not found")
}

func (d *DBConfig) mysql() gorm.Dialector {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		d.Username,
		d.Password,
		d.Host,
		d.Port,
		d.Database,
		d.Extend)

	return mysql.Open(dsn)
}

func (d *DBConfig) sqlite() gorm.Dialector {
	return sqlite.Open(d.SqliteFile)
}

func (d *DBConfig) postgres() gorm.Dialector {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		d.Host,
		d.Username,
		d.Password,
		d.Database,
		d.Port,
		d.SslMode,
		d.TimeZone)

	return postgres.Open(dsn)
}
