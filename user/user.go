package user

import "testReflection/db"

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// 1) create the user obj & callBack func using reflection &
//    call the callback func with new doc as parameter
func GetUsers_CallBack_ReflectionFunc() []User {
	users := make([]User, 0, 5)
	db.QueryResults_CallBack_ReflectionFunc(func(user *User) {
		users = append(users, *user)
	})
	return users
}

// 2) pass the user object and callback function,
//    so that db pkg doesn't have to create user object & callback function using reflection
func GetUsers_CallBack_ObjNFunc() []User {
	users := make([]User, 0, 5)
	db.QueryResults_CallBack_ObjNFunc(&User{}, func(userI interface{}) {
		u := userI.(*User)
		users = append(users, *u)
	})
	return users
}

// 3) get db client as callback & read user
func GetUsers_CallBack_DBClient() []User {
	users := make([]User, 0, 5)
	db.QueryResults_CallBack_DBClient(func(client db.DbClient) error {
		user := User{}
		if err := client.ReadDocument(&user); err != nil {
			return err
		}
		users = append(users, user)
		return nil
	})

	return users
}

// 4) handle db client in user package
func GetUsers_ExposeClient() []User {
	users := make([]User, 0, 5)

	client := db.QueryResults_ExposeClient()
	for client.HasMore() {
		user := User{}
		client.ReadDocument(&user)
		users = append(users, user)
	}
	client.Close() // responsibility of caller to close the client

	return users
}
