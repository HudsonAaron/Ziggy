package gtcp

import (
	"encoding/binary"
	"fmt"
)

// 解析数据包
func (tm *TcpMsg) Parse() ([]byte, error) {
	msgLen, err := tm.ParseMsgLen()
	if err != nil {
		return nil, err
	}
	if msgLen == 0 {
		return tm.data, nil
	}
	// 返回剩余的消息
	leftMsg := tm.data[4+msgLen:]
	tm.data = tm.data[4 : 4+msgLen]
	return leftMsg, nil
}

// 解析消息长度
func (tm *TcpMsg) ParseMsgLen() (uint32, error) {
	// 判断是否是大端序
	if tm.bigEndian {
		// 取出4字节，解析成uint32，为消息长度
		tm.len = binary.BigEndian.Uint32(tm.data[0:4])
		if tm.len > uint32(MaxMsgLen) {
			return 0, fmt.Errorf("message length exceeds maximum value")
		}
		// 判断剩余的消息长度是否足够
		if len(tm.data) < int(tm.len) {
			return 0, nil // 可能数据包收取不全
		}
		return tm.len, nil
	} else {
		tm.len = binary.LittleEndian.Uint32(tm.data[0:4])
		if tm.len > uint32(MaxMsgLen) {
			return 0, fmt.Errorf("message length exceeds maximum value")
		}
		if len(tm.data) < int(tm.len) {
			return 0, nil // 可能数据包收取不全
		}
		return tm.len, nil
	}
}
