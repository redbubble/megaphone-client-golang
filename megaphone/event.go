package megaphone

import (
	"encoding/json"
)

type event struct {
	Origin       string                 `json:"origin"`
	Topic        string                 `json:"topic"`
	Subtopic     string                 `json:"subtopic"`
	Schema       string                 `json:"schema"`
	PartitionKey string                 `json:"partitionKey"`
	Payload      map[string]interface{} `json:"data"`
}

func newEvent(topic, subtopic, schema, partitionKey string, payload []byte) (*event, error) {
	var mapPayload map[string]interface{}
	err := json.Unmarshal([]byte(payload), &mapPayload)
	if err != nil {
		return &event{}, err
	}

	return &event{
		Topic:        topic,
		Subtopic:     subtopic,
		Schema:       schema,
		PartitionKey: partitionKey,
		Payload:      mapPayload,
	}, nil
}

func (e *event) toJSON() (string, error) {
	bytes, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
