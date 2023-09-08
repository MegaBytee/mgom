package mgom

type ErrorCode int

// tags
const (
	DEFAULT ErrorCode = iota + 100
	CONNECT
	INDEX_CREATE
	SAVE
	COUNT
	UPDATE
	DELETE
	GET_CURSOR
)

func handleErrors(tag ErrorCode, err error) int {
	if err != nil {
		return int(tag)
	}
	return 0
}
