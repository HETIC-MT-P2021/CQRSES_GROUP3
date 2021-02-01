package domain

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/core/cqrs"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/domain/articles"
)

var Cb *cqrs.CommandBus
var Qb *cqrs.QueryBus

// InitBuses initializes the command and query buses
func InitBuses() {
	Cb = cqrs.NewCommandBus()
	Qb = cqrs.NewQueryBus()

	//> -------------------- ARTICLE COMMANDS ---------------------------
	_ = Cb.RegisterHandler(articles.NewArticleCommandHandler(), &articles.CreateArticleCommand{})
	// _ = Cb.RegisterHandler(articles.NewArticleCommandHandler(), articles.EditArticleCommand{})
	// _ = Cb.RegisterHandler(articles.NewArticleCommandHandler(), articles.DeleteArticleCommand{})
	//<-------------------- ARTICLE COMMANDS ---------------------------

	//> -------------------- ARTICLE QUERIES ---------------------------
	// _ = Qb.RegisterHandler(articles.NewArticleQueryHandler(), articles.GetArticleQuery{})
	//< -------------------- ARTICLE QUERIES ---------------------------
}