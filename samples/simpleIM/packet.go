package simpleIM

import (
	"encoding/json"

	"encoding/binary"

	"bytes"
)

type TPacket struct {
	CMD     int    `json:"cmd"`
	NK      string `json:"nk"`
	Content string `json:"content"`
}

func Encode(cmd int, nk, msg string) ([]byte, error) {
	var p TPacket
	p.CMD = cmd
	p.NK = nk
	p.Content = msg
	bs, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	buff := bytes.NewBuffer([]byte{})
	binary.Write(buff, binary.LittleEndian, uint32(len(bs)))
	buff.Write(bs)
	return buff.Bytes(), nil
}

func Decode(data []byte) (TPacket, error) {
	var ret TPacket
	err := json.Unmarshal(data, &ret)
	return ret, err
}
