package tool

import (
	"fmt"
)

func GenerateState() string {
	stateBytes := RandString(7)
	state := fmt.Sprintf("%x", []byte(stateBytes))
	return state
}
