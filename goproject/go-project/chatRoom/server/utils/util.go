package utils

import (
	"chatRoom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (transfer *Transfer) ReadPkg() (msg message.Message, err error) {
	// 先读取长度
	_, err = transfer.Conn.Read(transfer.Buf[:4])
	fmt.Println("等待读取...")
	if err != nil {
		fmt.Printf("conn.Read(buf[:4]) err, err=%v", err)
		return
	}
	fmt.Printf("读到的buf=%v\n", transfer.Buf[:4])
	// 根据读到的长度读取数据
	pkgLength := binary.BigEndian.Uint32(transfer.Buf[:4])
	n, err := transfer.Conn.Read(transfer.Buf[:pkgLength])
	if uint32(n) != pkgLength || err != nil {
		fmt.Printf("conn.Read(buf[:pkgLength]) fail, err=%v", err)
		return
	}
	err = json.Unmarshal(transfer.Buf[:pkgLength], &msg)
	if err != nil {
		fmt.Printf("json.Unmarshal(buf[:pkgLength], &msg) fail, err=%v", err)
		return
	}
	return
}

func (transfer *Transfer) WritePkg(data []byte) (err error) {
	// 先发送一个长度
	pkgLength := uint32(len(data))
	binary.BigEndian.PutUint32(transfer.Buf[:4], pkgLength)
	n, err := transfer.Conn.Write(transfer.Buf[:4])
	if n != 4 || err != nil {
		fmt.Printf("conn.Write(buf) err, err=%v", err)
		return
	}
	// 发送数据
	n, err = transfer.Conn.Write(data)
	if n != int(pkgLength) || err != nil {
		fmt.Printf("conn.Write(data) err, err=%v", err)
		return
	}
	return
}
