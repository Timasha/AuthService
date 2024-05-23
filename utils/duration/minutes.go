package duration

import (
	"encoding/json"
	"time"
)

type Minutes struct {
	time.Duration
}

func (d *Minutes) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Minutes())
}

func (d *Minutes) UnmarshalJSON(b []byte) error {
	var seconds int64
	if err := json.Unmarshal(b, &seconds); err != nil {
		return err
	}

	d.Duration = time.Duration(seconds * int64(time.Minute))

	return nil
}
