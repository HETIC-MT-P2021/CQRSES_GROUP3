package helpers

import (
	"reflect"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

type decodeTest struct {
	Date time.Time
}

func TestDecode(t *testing.T) {
	type args struct {
		input  map[string]interface{}
		output decodeTest
	}

	dateTestFailure := map[string]interface{}{"date": "24O11996"}

	dateTestSuccess := map[string]interface{}{"date": "24/01/1996"}

	tests := []struct {
		name    string
		args    args
	}{
		{
			name: "Test with string success",
			args: args{
				input: dateTestSuccess,
				output: decodeTest{},
			},
		},		
		{
			name: "Test with string failure",
			args: args{
				input: dateTestFailure,
				output: decodeTest{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Decode(tt.args.input, &tt.args.output); (err != nil){
				logrus.Error(tt.args.output)
				if reflect.TypeOf(tt.args.output.Date).String() != reflect.TypeOf(time.Now()).String() {
					t.Errorf("Decode() error = %v", err)
				}
			}
		})
	}
}
