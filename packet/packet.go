package packet

// Packet协议定义

/*
### packet header
1 byte: commandID

### submit packet

8字节 ID 字符串
任意字节 payload

### submit ack packet

8字节 ID 字符串
1字节 result
*/

const (
	CommandConn   = iota + 0x01 // 0x01, 连接请求包
	CommandSubmit               // 0x02, 消息请求包
)

const (
	CommandConnAck   = iota + 0x80 // 0x81, 连接请求的响应包
	CommandSubmitAck               //0x82, 消息请求的响应包
)
