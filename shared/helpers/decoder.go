package helpers

import (
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
)

// Decode create a new hook in the mapstructure pkg for the decoder config.
// Especially used to map a string to a time.Time type.
func Decode(input interface{}, output interface{}) error {

	stringToDateTimeHook := func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if t == reflect.TypeOf(time.Time{}) && f == reflect.TypeOf("") {
			return time.Parse(time.RFC3339, data.(string))
		}
		return data, nil
	}

	config := mapstructure.DecoderConfig{
		DecodeHook: stringToDateTimeHook,
		Result:     &output,
	}

	decoder, err := mapstructure.NewDecoder(&config)
	if err != nil {
		return err
	}
	
	err = decoder.Decode(input)
	if err != nil {
		return err
	}
	return nil
}
