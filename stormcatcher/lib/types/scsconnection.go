package types

import (
	pb "stormcatcher/lib/proto"
)

type ScsServerConnection struct {
	ScsHostAddress string
	ScsId          string
	LiveFlag       bool
	Stream         *pb.Vnode_ScsPushServer
	Req            chan *pb.ScsPushMsg
	Cancel         chan bool
	RetryCount     uint
}
