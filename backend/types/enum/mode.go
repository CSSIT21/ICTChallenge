package enum

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Mode string

const (
	ModePreview Mode = "preview"
	ModeStarted Mode = "started"
	ModeEnded   Mode = "ended"
)

func (r *Mode) UnmarshalJSON(data []byte) error {
	v := new(string)
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}

	val := Mode(*v)
	if val != ModePreview && val != ModeStarted && val != ModeEnded {
		return errors.New(fmt.Sprintf("%s is not a valid Type value", val))
	}

	*r = val
	return nil
}
