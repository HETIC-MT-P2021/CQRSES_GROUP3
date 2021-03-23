package cqrs

type CommandHandler interface {
	Handle(CommandMessage) (interface{}, error)
}

type QueryHandler interface {
	Handle(QueryMessage) (interface{}, error, int)
}
