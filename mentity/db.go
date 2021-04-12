package mentity

import "fmt"

type Db struct {
	Host     string `yaml:"host" json:"host"`
	Port     uint32 `yaml:"port" json:"port"`
	User     string `yaml:"user" json:"user"`
	Password string `yaml:"password" json:"password"`
	Name     string `yaml:"name" json:"name"`
}

func (d Db) DSN() string {
	if d.Port == 0 {
		d.Port = 3306
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", d.User, d.Password, d.Host, d.Port, d.Name)
}
