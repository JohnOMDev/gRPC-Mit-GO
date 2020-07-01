module server

go 1.13

require (
	csvdata/pb v0.0.0-00010101000000-000000000000
	github.com/fsnotify/fsnotify v1.4.9
	google.golang.org/grpc v1.30.0
	google.golang.org/protobuf v1.25.0 // indirect
)

replace csvdata/pb => ../pb
