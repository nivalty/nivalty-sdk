package msgs

type MessageType struct {
	t string
}

func (mt *MessageType) String() string {
	return mt.t
}

type Message interface {
	Bytes() []byte
	Type() MessageType
}
