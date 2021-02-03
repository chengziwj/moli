package entity

import "fmt"

type Db struct {
	Host     string `yaml:"host";json:"host"`
	Port     int    `yaml:"port";json:"port"`
	User     string `yaml:"user";json:"user"`
	Password string `yaml:"password";json:"password"`
	Name     string `yaml:"name";json:"name"`
}

func (d Db) DSN() string {
	port := d.Port
	if port == 0 {
		port = 3306
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", d.User, d.Password, d.Host, port, d.Name)
}
