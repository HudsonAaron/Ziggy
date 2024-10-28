package gwebsocket

// websocket消息结构体
type GWSMsg struct {
	Field string      // 字段名
	Value interface{} // 字段值
}

// 监听消息广播通道
func (gws *GWServer) BroadcastMsg() {
	for {
		msg := <-gws.broadcast
		for client := range gws.clients {
			if client.conn != nil {
				err := client.conn.WriteJSON(msg)
				if err != nil {
					client.conn.Close()
					delete(gws.clients, client)
				}
			} else {
				delete(gws.clients, client)
			}
		}
	}
}
