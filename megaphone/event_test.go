package megaphone

import (
	"testing"

	"github.com/karlseguin/typed"
	"github.com/stretchr/testify/assert"
)

func TestEvent(t *testing.T) {

	t.Run("toJSON()", func(t *testing.T) {
		origin := "my-awesome-service"
		topic := "work-updates"
		subtopic := "work-metadata-updated"
		schema := "https://github.com/redbubble/megaphone-event-type-registry/blob/master/streams/work-updates-schema-1.0.0.json"
		partitionKey := "1357924680"
		payload := []byte("{\"url\": \"https://www.redbubble.com/people/wytrab8/works/26039653-toadally-rad\"}")

		actualEvent, err := newEvent(topic, subtopic, schema, partitionKey, payload)
		actualEvent.Origin = origin
		assert.Nil(t, err)
		json, err := actualEvent.toJSON()
		assert.Nil(t, err)
		res, err := typed.JsonString(json)
		assert.Nil(t, err)

		assert.Equal(t, origin, res.String("origin"))
		assert.Equal(t, topic, res.String("topic"))
		assert.Equal(t, subtopic, res.String("subtopic"))
		assert.Equal(t, schema, res.String("schema"))
		assert.Equal(t, partitionKey, res.String("partitionKey"))
		assert.Equal(t, "https://www.redbubble.com/people/wytrab8/works/26039653-toadally-rad", res.Object("data").String("url"))
	})

}
