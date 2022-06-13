package model

type User struct {
	Id       int    `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role"`
}

type Users []User

func (u User) InitData() Users {
	listUsers := Users{}
	bob := User{
		Id:       1,
		Username: "bob",
		Password: "123456",
		Role:     "user",
	}
	alice := User{
		Id:       2,
		Username: "alice",
		Password: "123456",
		Role:     "admin",
	}
	listUsers = append(listUsers, bob, alice)
	return listUsers
}

func (u Users) FindUser(user User) *User {
	for _, item := range u {
		if item.Username == user.Username && item.Password == user.Password {
			return &item
		}
	}
	return nil
}
