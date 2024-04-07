package parser

import (
	"github.com/MrClean-code/testD/pkg/model"
	"net/http"
)

type SiteParser interface {
	ParseData() ([]model.Deal, error)
}

type ParserList struct {
	SiteParser
}

func NewParserList(r *http.Request) *ParserList {
	return &ParserList{
		SiteParser: NewParserDealList(r),
	}
}
