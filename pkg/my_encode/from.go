package my_encode

import "github.com/go-kratos/kratos/v2/encoding"

var _ encoding.Codec = FormCodec{}

const (
	// Name is form codec name
	Name = "form-data"
	// Null value string
	nullStr = "null"
)

type FormCodec struct {
}

func (f FormCodec) Marshal(v interface{}) ([]byte, error) {
	return nil, nil
}

func (f FormCodec) Unmarshal(data []byte, v interface{}) error {
	// 解析
}

func (f FormCodec) Name() string {
	return Name
}
