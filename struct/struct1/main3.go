package main

type User struct {
	Name string
	ID string
	Age int
}

type VIPUser struct {
	UserInfo User
	VIPLevel int
	Price int
}

func main() {
	user := User { "김미정", "mi", 31 }
	println(user.Name)

	vip := VIPUser { user, 1, 1}
	println(vip.UserInfo.Name)
	println(vip.UserInfo.ID)
}