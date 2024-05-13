package util

import (
	"github.com/ProtoconNet/mitum-currency/v3/cmds"
	"github.com/ProtoconNet/mitum2/util/encoder"
	jsonenc "github.com/ProtoconNet/mitum2/util/encoder/json"
)

var (
	Encoders *encoder.Encoders
	enc      encoder.Encoder
)

func init() {
	enc = jsonenc.NewEncoder()
	Encoders = encoder.NewEncoders(enc, enc)
	if err := cmds.LoadHinters(Encoders); err != nil {
		panic(err)
	}
}
