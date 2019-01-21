package Packet

type DataStruct struct {
	//header length(字节数==bytesOfLength) data
	Header        []byte
	BytesOfLength int //length所占字节数
}
