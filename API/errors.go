package API
/*
import (
	"github.com/tildae/vkgo/updates"
	"github.com/tildae/vkgo/updates/scenes"
)*/

type Error struct{
	Code	int		`json:"error_code"`
	Message	string		`json:"error_msg"`
	Params	[]ErrorParams	`json:"request_params"`
}
type ErrorParams struct{
	Key	string	`json:"key"`
	Value	string	`json:"value"`
}
/*
type Err map[string]func(bot *Options, update updates.Update)
var Errs = Err{}

func(e Err) Add(typ string, err Err) {
	ntyp := scenes.Types[typ]
	if ntyp == "" { ntyp = typ }
	e[ntyp] = err
}*/