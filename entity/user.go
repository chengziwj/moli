package entity

type User struct {
	Uid      int    `db:"uid";json:"uid"`
	Name     string `db:"name";json:"name"`
	Password string `db:"password";json:"password"`
	Salt     string `db:"salt";json:"salt"`
}
