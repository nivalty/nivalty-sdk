package keys

type Signature struct {
	value []byte
}

func (signature *Signature) Bytes() []byte {
	return signature.value
}
