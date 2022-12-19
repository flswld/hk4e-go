package net

import (
	"bytes"
	"encoding/binary"

	"hk4e/pkg/endec"
	"hk4e/pkg/logger"
)

/*
										原神KCP协议(带*为xor加密数据)
0			1			2					4											8(字节)
+---------------------------------------------------------------------------------------+
|											conv										|
+---------------------------------------------------------------------------------------+
|	cmd		|	frg		|		wnd			|					ts						|
+---------------------------------------------------------------------------------------+
|						sn					|					una						|
+---------------------------------------------------------------------------------------+
|						len					|		0X4567*		|		cmdId*			|
+---------------------------------------------------------------------------------------+
|		headLen*		|				payloadLen*				|		head*			|
+---------------------------------------------------------------------------------------+
|								payload*						|		0X89AB*			|
+---------------------------------------------------------------------------------------+
*/

type KcpMsg struct {
	ConvId    uint64
	CmdId     uint16
	HeadData  []byte
	ProtoData []byte
}

func (k *KcpConnectManager) decodeBinToPayload(data []byte, convId uint64, kcpMsgList *[]*KcpMsg, xorKey []byte) {
	// xor解密
	endec.Xor(data, xorKey)
	k.decodeLoop(data, convId, kcpMsgList)
}

func (k *KcpConnectManager) decodeLoop(data []byte, convId uint64, kcpMsgList *[]*KcpMsg) {
	// 长度太短
	if len(data) < 12 {
		logger.Debug("packet len less 12 byte")
		return
	}
	// 头部幻数错误
	if data[0] != 0x45 || data[1] != 0x67 {
		logger.Error("packet head magic 0x4567 error")
		return
	}
	// 协议号
	cmdIdByteSlice := make([]byte, 8)
	cmdIdByteSlice[6] = data[2]
	cmdIdByteSlice[7] = data[3]
	cmdIdBuffer := bytes.NewBuffer(cmdIdByteSlice)
	var cmdId int64
	err := binary.Read(cmdIdBuffer, binary.BigEndian, &cmdId)
	if err != nil {
		logger.Error("packet cmd id parse fail: %v", err)
		return
	}
	// 头部长度
	headLenByteSlice := make([]byte, 8)
	headLenByteSlice[6] = data[4]
	headLenByteSlice[7] = data[5]
	headLenBuffer := bytes.NewBuffer(headLenByteSlice)
	var headLen int64
	err = binary.Read(headLenBuffer, binary.BigEndian, &headLen)
	if err != nil {
		logger.Error("packet head len parse fail: %v", err)
		return
	}
	// proto长度
	protoLenByteSlice := make([]byte, 8)
	protoLenByteSlice[4] = data[6]
	protoLenByteSlice[5] = data[7]
	protoLenByteSlice[6] = data[8]
	protoLenByteSlice[7] = data[9]
	protoLenBuffer := bytes.NewBuffer(protoLenByteSlice)
	var protoLen int64
	err = binary.Read(protoLenBuffer, binary.BigEndian, &protoLen)
	if err != nil {
		logger.Error("packet proto len parse fail: %v", err)
		return
	}
	// 检查最小长度
	if len(data) < int(headLen+protoLen)+12 {
		logger.Error("packet len error")
		return
	}
	// 尾部幻数错误
	if data[headLen+protoLen+10] != 0x89 || data[headLen+protoLen+11] != 0xAB {
		logger.Error("packet tail magic 0x89AB error")
		return
	}
	// 判断是否有不止一个包
	haveMoreData := false
	if len(data) > int(headLen+protoLen)+12 {
		haveMoreData = true
	}
	// 头部数据
	headData := data[10 : 10+headLen]
	// proto数据
	protoData := data[10+headLen : 10+headLen+protoLen]
	// 返回数据
	kcpMsg := new(KcpMsg)
	kcpMsg.ConvId = convId
	kcpMsg.CmdId = uint16(cmdId)
	// kcpMsg.HeadData = make([]byte, len(headData))
	// copy(kcpMsg.HeadData, headData)
	// kcpMsg.ProtoData = make([]byte, len(protoData))
	// copy(kcpMsg.ProtoData, protoData)
	kcpMsg.HeadData = headData
	kcpMsg.ProtoData = protoData
	*kcpMsgList = append(*kcpMsgList, kcpMsg)
	// 递归解析
	if haveMoreData {
		k.decodeLoop(data[int(headLen+protoLen)+12:], convId, kcpMsgList)
	}
}

func (k *KcpConnectManager) encodePayloadToBin(kcpMsg *KcpMsg, xorKey []byte) (bin []byte) {
	if kcpMsg.HeadData == nil {
		kcpMsg.HeadData = make([]byte, 0)
	}
	if kcpMsg.ProtoData == nil {
		kcpMsg.ProtoData = make([]byte, 0)
	}
	bin = make([]byte, len(kcpMsg.HeadData)+len(kcpMsg.ProtoData)+12)
	// 头部幻数
	bin[0] = 0x45
	bin[1] = 0x67
	// 协议号
	cmdIdBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(cmdIdBuffer, binary.BigEndian, kcpMsg.CmdId)
	if err != nil {
		logger.Error("cmd id encode err: %v", err)
		return nil
	}
	bin[2] = (cmdIdBuffer.Bytes())[0]
	bin[3] = (cmdIdBuffer.Bytes())[1]
	// 头部长度
	headLenBuffer := bytes.NewBuffer([]byte{})
	err = binary.Write(headLenBuffer, binary.BigEndian, uint16(len(kcpMsg.HeadData)))
	if err != nil {
		logger.Error("head len encode err: %v", err)
		return nil
	}
	bin[4] = (headLenBuffer.Bytes())[0]
	bin[5] = (headLenBuffer.Bytes())[1]
	// proto长度
	protoLenBuffer := bytes.NewBuffer([]byte{})
	err = binary.Write(protoLenBuffer, binary.BigEndian, uint32(len(kcpMsg.ProtoData)))
	if err != nil {
		logger.Error("proto len encode err: %v", err)
		return nil
	}
	bin[6] = (protoLenBuffer.Bytes())[0]
	bin[7] = (protoLenBuffer.Bytes())[1]
	bin[8] = (protoLenBuffer.Bytes())[2]
	bin[9] = (protoLenBuffer.Bytes())[3]
	// 头部数据
	copy(bin[10:], kcpMsg.HeadData)
	// proto数据
	copy(bin[10+len(kcpMsg.HeadData):], kcpMsg.ProtoData)
	// 尾部幻数
	bin[len(bin)-2] = 0x89
	bin[len(bin)-1] = 0xAB
	// xor加密
	endec.Xor(bin, xorKey)
	return bin
}
