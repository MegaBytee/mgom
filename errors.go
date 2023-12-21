package mgom

import "fmt"

// tags
const (
	DEFAULT      = iota + 100 //100
	CONNECT                   //101
	INDEX_CREATE              //102
	SAVE                      //103
	COUNT                     //104
	UPDATE                    //105
	DELETE                    //106
	GET_CURSOR                //107
	PAGINATE                  //108
)

type Error struct {
	Code int
	Msg  string
}

func handleErrors(tag int, err error) Error {
	e := Error{
		Code: 0,
		Msg:  "Ok",
	}
	if err != nil {
		e.Code = tag
		e.Msg = err.Error()
		fmt.Println("err=", e.Msg)
	}
	return e
}
