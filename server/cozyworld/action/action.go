package action

import (
	"encoding/json"
	"fmt"
)

type ActionType int64

const (
	UnknownActionType   ActionType = 0
	ConnectedActionType ActionType = 1
	MoveActionType      ActionType = 2
)

// Action wrap the underlying data blob with helper functions to access data.
type Action interface {
	Type() (ActionType, error)
	Data() map[string]any
}

// UnclassifiedAction implements the minimal interface of an action.
type UnclassifiedAction struct {
	Action

	data map[string]any
}

func (u *UnclassifiedAction) Type() (ActionType, error) {
	return getValue[ActionType](u, "type")
}

func (u *UnclassifiedAction) Data() map[string]any {
	return u.data
}

func JsonToAction(data []byte) Action {
	a := &UnclassifiedAction{data: make(map[string]any)}
	json.Unmarshal(data, &a.data)
	return a
}

func getValue[T any](a Action, path ...string) (T, error) {
	var (
		empT   T
		maybeT any
		ok     bool
	)

	for _, key := range path {
		maybeT, ok = a.Data()[key]
		if !ok {
			return empT, fmt.Errorf("Missing %q field from data.", path)
		}

	}

	t, ok := maybeT.(T)
	if !ok {
		return empT, fmt.Errorf("Unable to cast value of key %v.", path)
	}

	return t, nil

}
