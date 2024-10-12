package gproto

import "google.golang.org/protobuf/proto"

// 解析协议
func ParseProto(buf []byte) error {
	protoData := &ProtoData{}
	err := proto.Unmarshal(buf, protoData)
	if err != nil {
		return err
	}
	for _, ptr := range protoData.Comdata {
		protoID := (*ptr).ProtoId
		data := (*ptr).Data
		ProtoRouter(protoID, data)
	}
	return nil
}
