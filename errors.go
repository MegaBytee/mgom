package mgom

import "fmt"

type ErrorCode int

// tags
const (
	DEFAULT      ErrorCode = iota + 100 //100
	CONNECT                             //101
	INDEX_CREATE                        //102
	SAVE                                //103
	COUNT                               //104
	UPDATE                              //105
	DELETE                              //106
	GET_CURSOR                          //107
)

func handleErrors(tag ErrorCode, err error) int {
	if err != nil {
		fmt.Println("err=", err)
		switch tag {
		case SAVE:
			return 0
		default:
			return int(tag)
		}

	}
	return 0
}
