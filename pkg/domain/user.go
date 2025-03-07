package domain



type User struct {
	ID         int    `json:"id" gorm:"primarykey"`
	Name       string `json:"name"`
	Email      string `json:"email" gorm:"unique"`
	Username   string `json:"username"`
	Phone int `json:"phone"`
	Password   string `json:"password"`
}