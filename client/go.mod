module csvclient

go 1.13

require (
	csvdata/pb v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.4.2
	golang.org/x/net v0.0.0-20190311183353-d8887717615a
	google.golang.org/grpc v1.30.0
	google.golang.org/protobuf v1.25.0
)

replace csvdata/pb => ../pb
