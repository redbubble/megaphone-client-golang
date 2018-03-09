package megaphone

import (
	"testing"

	"github.com/karlseguin/typed"
	"github.com/stretchr/testify/require"
)

func TestEvent(t *testing.T) {

	t.Run("toJSON()", func(t *testing.T) {
		required := require.New(t)
		origin := "my-awesome-service"
		topic := "work-updates"
		subtopic := "work-metadata-updated"
		schema := "https://github.com/redbubble/megaphone-event-type-registry/blob/master/streams/work-updates-schema-1.0.0.json"
		partitionKey := "1357924680"
		payload := []byte("{\"url\": \"https://www.redbubble.com/people/wytrab8/works/26039653-toadally-rad\"}")

		actualEvent, err := newEvent(topic, subtopic, schema, partitionKey, payload)
		actualEvent.Origin = origin
		required.NoError(err)
		json, err := actualEvent.toJSON()
		required.NoError(err)
		res, err := typed.JsonString(json)
		required.NoError(err)

		required.Equal(origin, res.String("origin"))
		required.Equal(topic, res.String("topic"))
		required.Equal(subtopic, res.String("subtopic"))
		required.Equal(schema, res.String("schema"))
		required.Equal(partitionKey, res.String("partitionKey"))
		required.Equal("https://www.redbubble.com/people/wytrab8/works/26039653-toadally-rad", res.Object("data").String("url"))
	})

	t.Run("streamName()", func(t *testing.T) {
		e := &event{
			Topic: "mega-updates",
		}
		deployEnv := "staging"

		require.Equal(t, "megaphone-streams-staging-mega-updates", e.streamName(deployEnv))
	})

}
