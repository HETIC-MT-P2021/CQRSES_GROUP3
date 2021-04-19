package cqrs

import (
	"reflect"
	"testing"
)

type testCommand struct {
	testField 	string
}

type testCommandHandler struct {}
func (t testCommandHandler) Handle(message CommandMessage) (interface{}, error) {
	return message.Payload(), nil
}


func TestCommandBus_Dispatch(t *testing.T) {
	b := NewCommandBus()
	b.RegisterHandler(testCommandHandler{}, testCommand{})

	type args struct {
		command CommandMessage
	}
	tests := []struct {
		name    string
		b       *CommandBus
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Test with handler",
			b: b,
			args: args{
				command: NewCommandMessage(testCommand{testField: "hello world"}),
			},
			want: nil,
			wantErr: false,
		},
		{
			name: "Test without handler",
			b: b,
			args: args{
				command: NewCommandMessage("hello world"),
			},
			want: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.b.Dispatch(tt.args.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommandBus.Dispatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestCommandBus_RegisterHandler(t *testing.T) {
	b := NewCommandBus()
	type args struct {
		handler CommandHandler
		command interface{}
	}
	tests := []struct {
		name    string
		b       *CommandBus
		args    args
		wantErr bool
	}{
		{
			name: "Test with handler",
			b: b,
			args: args{
				command: testCommand{testField: "hello world"},
				handler: testCommandHandler{},
			},
			wantErr: false,
		},
		{
			name: "Test with existing handler",
			b: b,
			args: args{
				command: testCommand{testField: "hello world"},
				handler: testCommandHandler{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.b.RegisterHandler(tt.args.handler, tt.args.command); (err != nil) != tt.wantErr {
				t.Errorf("CommandBus.RegisterHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewCommandMessage(t *testing.T) {
	type args struct {
		command interface{}
	}
	tests := []struct {
		name string
		args args
		want *CommandDescriptor
	}{
		{
			name: "test basic string payload",
			args: args{
				command: "hello test",
			},
			want: &CommandDescriptor{command: "hello test"},
		},
		{
			name: "",
			args: args{
				command: 1,
			},
			want: &CommandDescriptor{command: 1},
		},
		{
			name: "",
			args: args{
				command: []string{"hello", "test"},
			},
			want: &CommandDescriptor{command: []string{"hello", "test"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCommandMessage(tt.args.command); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCommandMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommandDescriptor_CommandType(t *testing.T) {
	tests := []struct {
		name string
		c    *CommandDescriptor
		want string
	}{
		{
			name: "test with testCommand",
			c: NewCommandMessage(testCommand{testField: "test"}),
			want: "testCommand",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.CommandType(); got != tt.want {
				t.Errorf("CommandDescriptor.CommandType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommandDescriptor_Payload(t *testing.T) {
	tests := []struct {
		name string
		c    *CommandDescriptor
		want interface{}
	}{
		{
			name: "test with basic payload",
			c: NewCommandMessage(testCommand{testField: "test"}),
			want: testCommand{testField: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Payload(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommandDescriptor.Payload() = %v, want %v", got, tt.want)
			}
		})
	}
}
