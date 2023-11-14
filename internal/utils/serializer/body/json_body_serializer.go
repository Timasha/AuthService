package body

import "encoding/json"

type JsonBodySerializer struct{}

func (j *JsonBodySerializer) Unmarshal(data []byte, serializableObject any) error {
	return json.Unmarshal(data, serializableObject)
}
func (j *JsonBodySerializer) Marshal(serializableObject any) ([]byte, error) {
	return json.Marshal(serializableObject)
}
