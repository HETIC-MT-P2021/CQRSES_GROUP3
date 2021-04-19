package cqrs

import (
	"fmt"
	"reflect"
)

type CommandMessage interface {
	CommandType() string
	Payload() interface{}
}

type CommandBus struct {
	handlers map[string]CommandHandler
}

func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[string]CommandHandler),
	}
}

func (b *CommandBus) Dispatch(command CommandMessage) (interface{}, error) {
	if handler, ok := b.handlers[command.CommandType()]; ok {
		return handler.Handle(command)
	}
	return nil, fmt.Errorf("the command bus does not have a handler for commands of type: %s", command.CommandType())
}

func (b *CommandBus) RegisterHandler(handler CommandHandler, command interface{}) error {
	typeName := reflect.TypeOf(command).Name()
	if _, ok := b.handlers[typeName]; ok {
		return fmt.Errorf("duplicate command handler registration with command bus for command of type: %s", typeName)
	}

	b.handlers[typeName] = handler

	return nil
}

type CommandDescriptor struct {
	command interface{}
}

func NewCommandMessage(command interface{}) *CommandDescriptor {
	return &CommandDescriptor{
		command: command,
	}
}

func (c *CommandDescriptor) CommandType() string {
	return reflect.TypeOf(c.command).Name()
}

// Command returns the actual command payload of the message.
func (c *CommandDescriptor) Payload() interface{} {
	return c.command
}
