package cqrs

import (
	"reflect"
	"testing"
)

type testQuery struct {
	testField 	string
}

type testQueryHandler struct {}

func (t testQueryHandler) Handle(message QueryMessage) (interface{}, error, int) {
	return message.Payload(), nil, 0
}

func TestQueryBus_Dispatch(t *testing.T) {
	b := NewQueryBus()
	b.RegisterHandler(testQueryHandler{}, testQuery{})

	type args struct {
		query QueryMessage
	}
	tests := []struct {
		name    string
		b       *QueryBus
		args    args
		want    interface{}
		want1   int
		wantErr bool
	}{
		{
			name: "Test with handler",
			b: b,
			args: args{
				query: NewQueryMessage(testQuery{testField: "hello"}),
			},
			want: testQuery{testField: "hello"},
			want1: 0,
			wantErr: false,
		},
		{
			name: "Test without handler",
			b: b,
			args: args{
				query: NewQueryMessage("hello"),
			},
			want: nil,
			want1: 0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err, got1 := tt.b.Dispatch(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryBus.Dispatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryBus.Dispatch() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("QueryBus.Dispatch() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestQueryBus_RegisterHandler(t *testing.T) {
	b := NewQueryBus()
	type args struct {
		handler QueryHandler
		query   interface{}
	}
	tests := []struct {
		name    string
		b       *QueryBus
		args    args
		wantErr bool
	}{
		{
			name: "Test with queryHandler",
			b: b,
			args: args{
				handler: testQueryHandler{},
				query: testQuery{testField: "hello"},
			},
			wantErr: false,
		},
		{
			name: "Test with same queryHandler",
			b: b,
			args: args{
				handler: testQueryHandler{},
				query: testQuery{testField: "hello"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.b.RegisterHandler(tt.args.handler, tt.args.query); (err != nil) != tt.wantErr {
				t.Errorf("QueryBus.RegisterHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewQueryMessage(t *testing.T) {
	type args struct {
		query interface{}
	}
	tests := []struct {
		name string
		args args
		want *QueryDescriptor
	}{
		{
			name: "Test basic message",
			args: args{
				query: testQuery{testField: "Hello"},
			},
			want: &QueryDescriptor{query: testQuery{testField: "Hello"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQueryMessage(tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQueryMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryDescriptor_QueryType(t *testing.T) {
	tests := []struct {
		name string
		c    *QueryDescriptor
		want string
	}{
		{
			name: "Test basic testQuery type",
			c: NewQueryMessage(testQuery{testField: "Hello"}),
			want: "testQuery",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.QueryType(); got != tt.want {
				t.Errorf("QueryDescriptor.QueryType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryDescriptor_Payload(t *testing.T) {
	tests := []struct {
		name string
		c    *QueryDescriptor
		want interface{}
	}{
		{
			name: "Test basic hello message",
			c: NewQueryMessage(testQuery{testField: "Hello"}),
			want: testQuery{testField: "Hello"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Payload(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryDescriptor.Payload() = %v, want %v", got, tt.want)
			}
		})
	}
}
