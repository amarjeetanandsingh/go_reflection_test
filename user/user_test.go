package user

import "testing"

// 1)
func BenchmarkGetUsers_CallBackReflectionFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		users := GetUsers_CallBack_ReflectionFunc()
		if len(users) > 0 {
		}
	}
}

//2)
func BenchmarkGetUsers_CallBack_ObjNFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		users := GetUsers_CallBack_ObjNFunc()
		if len(users) > 0 {
		}
	}
}

func BenchmarkGetUsers_CallBackClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		users := GetUsers_CallBack_DBClient()
		if len(users) > 0 {
		}
	}
}

func BenchmarkGetUsers_ExposeClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		users := GetUsers_ExposeClient()
		if len(users) > 0 {
		}
	}
}
