# go_reflection_test

## go test ./user  -cpuprofile cpu.prof -memprofile mem.prof -bench='BenchmarkGetUser*' -run=^a -benchmem  -test.benchtime=10s
```
goos: darwin
goarch: amd64
pkg: testReflection/user
BenchmarkGetUsers_CallBackReflectionFunc-12    	 4426987	       263 ns/op	     432 B/op	       4 allocs/op
BenchmarkGetUsers_CallBack_ObjNFunc-12         	 6293131	       186 ns/op	     384 B/op	       2 allocs/op
BenchmarkGetUsers_CallBackClient-12            	 7858377	       147 ns/op	     320 B/op	       1 allocs/op
BenchmarkGetUsers_ExposeClient-12              	10273468	       106 ns/op	     320 B/op	       1 allocs/op
PASS
ok  	testReflection/user	5.775s
```

## go tool pprof cpu.prof