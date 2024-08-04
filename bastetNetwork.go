package wjbastetLibWSPackage

import (
	gatlingWSProtocol "github.com/gatlinglab/libGatlingWS/modProtocol"
)

type CWJBWSP_ParseData1 struct {
	RequestID uint16
	CMD1      byte
	CMD2      byte
	CMD3      byte
}

type CWJBWSP_Parser1 struct {
	parseData CWJBWSP_ParseData1
	sock      gatlingWSProtocol.IWJSocket
}

func WJBWSP_CreateParser1(socket gatlingWSProtocol.IWJSocket) *CWJBWSP_Parser1 {
	return &CWJBWSP_Parser1{}
}

func (pInst *CWJBWSP_Parser1) DataParse(data []byte) (*CWJBWSP_ParseData1, int) {
	datalen := len(data)
	if datalen < WJBP_LengthBasicData {
		return nil, -1
	}
	pInst.parseData.RequestID = uint16(data[0])<<8 | uint16(data[1])
	pInst.parseData.CMD1 = data[WJBP_OffsetCommand1]
	pInst.parseData.CMD2 = data[WJBP_OffsetCommand2]
	pInst.parseData.CMD3 = data[WJBP_OffsetCommand3]

	return &pInst.parseData, WJBP_LengthBasicData
	// the rest data is data[WJBP_LengthBasicData:]
}
func (pInst *CWJBWSP_Parser1) CommandSend(cmd1, cmd2, cmd3 byte, requestid uint16) error {
	_, err := pInst.DataSend(cmd1, cmd2, cmd3, requestid, nil)
	return err
}
func (pInst *CWJBWSP_Parser1) CommandSend2(cmd1, cmd2, cmd3 byte) error {
	_, err := pInst.DataSend(cmd1, cmd2, cmd3, 0, nil)
	return err
}

func (pInst *CWJBWSP_Parser1) DataSend(cmd1, cmd2, cmd3 byte, requestid uint16, data []byte) (int, error) {
	datalen := 0
	if data != nil {
		datalen = len(data)
	}
	dataSend := make([]byte, WJBP_LengthBasicData+datalen)
	if requestid != 0 {
		dataSend[0] = byte(requestid >> 8)
		dataSend[1] = byte(requestid)
	}
	dataSend[WJBP_OffsetCommand1] = cmd1
	dataSend[WJBP_OffsetCommand2] = cmd2
	dataSend[WJBP_OffsetCommand3] = cmd3

	if data != nil {
		copy(dataSend[WJBP_LengthBasicData:], data)
	}

	err := pInst.sock.WriteBinary(dataSend)

	return datalen, err
}
func (pInst *CWJBWSP_Parser1) DataSend2(cmd1, cmd2, cmd3 byte, data []byte) (int, error) {
	return pInst.DataSend(cmd1, cmd2, cmd3, 0, data)
}
