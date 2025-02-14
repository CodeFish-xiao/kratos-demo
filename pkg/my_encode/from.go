package my_encode

import (
	"github.com/go-kratos/kratos/v2/encoding"
)

var _ encoding.Codec = FormCodec{}

const (
	// Name is form codec name
	Name = "form-data"
)

type FormCodec struct {
}

func (f FormCodec) Marshal(v interface{}) ([]byte, error) {
	return nil, nil
}

func (f FormCodec) Unmarshal(data []byte, v interface{}) error {
	return nil
}

func (f FormCodec) Name() string {
	return Name
}
