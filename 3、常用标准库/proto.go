package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

func Encode(message string) ([]byte, error) { // Encode 将消息编码
	var length = int32(len(message)) // 读取消息的长度，转换成int32类型（占4个字节）
	var pkg = new(bytes.Buffer)

	err := binary.Write(pkg, binary.LittleEndian, length) // 写入消息头
	if err != nil {
		return nil, err
	}

	err = binary.Write(pkg, binary.LittleEndian, []byte(message)) // 写入消息实体
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

func Decode(reader *bufio.Reader) (string, error) { // Decode 解码消息
	// 读取消息的长度
	lengthByte, _ := reader.Peek(4) // 读取前4个字节的数据
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}

	if int32(reader.Buffered()) < length+4 { // Buffered返回缓冲中现有的可读取的字节数。
		return "", err
	}

	pack := make([]byte, int(4+length)) // 读取真正的消息数据
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
