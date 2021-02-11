package decoders

import (
	"sync"

	"github.com/gorilla/schema"
)

var (
	decoder *schema.Decoder
	once    sync.Once
)

func ParamsDecoder() *schema.Decoder {
	once.Do(func() {
		decoder = schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)
		decoder.SetAliasTag("param")
	})

	return decoder
}
