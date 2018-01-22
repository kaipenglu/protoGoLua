package codec

import (
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"log"
	"reflect"
)

type MessageHandler func(msgId uint32, msg, client interface{})

type Codec struct {
	handlers []MessageHandler
	protoMap map[uint32]reflect.Type
}

func NewCodec() *Codec {
	cdc := &Codec{}
	cdc.protoMap = make(map[uint32]reflect.Type)
	return cdc
}

func (cdc *Codec) RegisterProto(msgId uint32, msg interface{}) {
	msgType := reflect.TypeOf(msg.(proto.Message))
	cdc.protoMap[msgId] = msgType
}

func (cdc *Codec) RegisterHandle(handler MessageHandler) {
	cdc.handlers = append(cdc.handlers, handler)
}

func (cdc *Codec) Decode(b []byte, client interface{}) {
	if len(b) < 4 {
		return
	}

	msgId := binary.LittleEndian.Uint32(b[:4])
	if msgType, ok := cdc.protoMap[msgId]; ok {
		msg := reflect.New(msgType.Elem()).Interface()
		err := proto.Unmarshal(b[4:], msg.(proto.Message))
		if err != nil {
			log.Fatal(err)
			return
		}

		for i := 0; i < len(cdc.handlers); i++ {
			cdc.handlers[i](msgId, msg, client)
		}
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

func (cdc *Codec) MsgIdToInterface(msgId uint32) interface{} {
	if msgType, ok := cdc.protoMap[msgId]; ok {
		msg := reflect.New(msgType.Elem()).Interface()
		return msg
	}
	return nil
}
