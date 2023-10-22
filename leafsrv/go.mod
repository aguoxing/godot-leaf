module server

go 1.19

require (
	github.com/name5566/leaf v0.0.0-20221021105039-af71eb082cda
	google.golang.org/protobuf v1.31.0
)

replace github.com/name5566/leaf v0.0.0-20221021105039-af71eb082cda => ../leaf

require github.com/gorilla/websocket v1.5.0 // indirect
