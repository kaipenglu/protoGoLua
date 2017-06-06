package codec

import (
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"log"
	"reflect"
)

type MessageHandler func(msgId uint32, msg interface{})

type MsgProtoInfo struct {
	msgType    reflect.Type
	msgHandler MessageHandler
}

type Codec struct {
	protoMap map[uint32]MsgProtoInfo
}

func NewCodec() *Codec {
	cdc := &Codec{}
	cdc.protoMap = make(map[uint32]MsgProtoInfo)
	return cdc
}

func (cdc *Codec) RegisterProto(msgId uint32, msg interface{}, handler MessageHandler) {
	var mpi MsgProtoInfo
	mpi.msgType = reflect.TypeOf(msg.(proto.Message))
	mpi.msgHandler = handler
	cdc.protoMap[msgId] = mpi
}

func (cdc *Codec) Decode(b []byte) {
	if len(b) < 4 {
		return
	}

	msgId := binary.LittleEndian.Uint32(b[:4])
	if mgi, ok := cdc.protoMap[msgId]; ok {
		msg := reflect.New(mgi.msgType.Elem()).Interface()
		err := proto.Unmarshal(b[4:], msg.(proto.Message))
		if err != nil {
			log.Fatal(err)
			return
		}
		mgi.msgHandler(msgId, msg)
	}
}

func (cdc *Codec) Encode(msg interface{}) (b []byte) {
	v := reflect.ValueOf(msg)
	cmd := reflect.Indirect(v).FieldByName("Cmd")
	if !cmd.IsValid() || cmd.Type().Kind() != reflect.Ptr {
		return
	}
	if cmd.Elem().Type().Kind() != reflect.Int32 {
		return
	}

	h := make([]byte, 4)
	msgId := uint32(cmd.Elem().Int())
	binary.LittleEndian.PutUint32(h, msgId)

	t, err := proto.Marshal(msg.(proto.Message))
	if err != nil {
		log.Fatal(err)
		return
	}
	b = append(h, t...)
	return
}
