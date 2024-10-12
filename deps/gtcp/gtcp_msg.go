package gtcp

// 消息结构体
// | 4字节 长度 | 数据 |
type TcpMsg struct {
	// 消息头长度
	len uint32
	// 消息体
	data []byte
	// 消息体类型 大端序 | 小端序
	bigEndian bool
}

var (
	// 消息体类型 大端序 | 小端序
	BigEndian = true
	// 消息体最长长度
	MaxMsgLen = 65535
)

// 初始化tcp消息
func (tc *TcpConn) NewTcpMsg() *TcpMsg {
	return &TcpMsg{data: tc.msg, bigEndian: BigEndian}
}
