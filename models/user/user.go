package user

type User struct {
	Id int `gorm:"primary_key"`
	Name string
	Age int
}