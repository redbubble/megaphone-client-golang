package megaphone

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {

	getConfig := func() Config {
		return Config{
			Origin: "my-awesome-service",
			Host:   os.Getenv("MEGAPHONE_FLUENT_HOST"),
			Port:   24224,
		}
	}

	t.Run("NewClient()", func(t *testing.T) {
		t.Run("Host defaults to MEGAPHONE_FLUENT_HOST", func(t *testing.T) {
			config := getConfig()
			config.Host = ""
			expectedHost := "localhost"
			os.Setenv("MEGAPHONE_FLUENT_HOST", expectedHost)
			client, err := NewClient(config)
			require.Nil(t, err)
			assert.Equal(t, expectedHost, client.config.Host)
			os.Unsetenv("MEGAPHONE_FLUENT_HOST")
		})

		t.Run("It fails when the port set by the user is not an valid number", func(t *testing.T) {
			config := getConfig()
			os.Setenv("MEGAPHONE_FLUENT_PORT", "not a valid port")
			_, err := NewClient(config)
			_, ok := err.(*ConfigError)
			assert.Equal(t, true, ok)
			os.Unsetenv("MEGAPHONE_FLUENT_PORT")
		})
	})

	t.Run("Publish()", func(t *testing.T) {

		type EventFields struct {
			Origin       string
			Topic        string
			Subtopic     string
			Schema       string
			PartitionKey string
			Payload      []byte
		}

		GetTestEventFields := func() EventFields {
			return EventFields{
				Origin:       "my-awesome-service",
				Topic:        "work-updates",
				Subtopic:     "work-metadata-updated",
				Schema:       "https://github.com/redbubble/megaphone-event-type-registry/blob/master/streams/work-updates-schema-1.0.0.json",
				PartitionKey: "1357924680",
				Payload:      []byte("{\"url\": \"https://www.redbubble.com/people/wytrab8/works/26039653-toadally-rad\"}"),
			}
		}

		t.Run("It publishes a message through the fluent logger", func(t *testing.T) {
			config := getConfig()
			client, err := NewClient(config)
			require.Nil(t, err)

			eventFields := GetTestEventFields()
			err = client.Publish(eventFields.Origin, eventFields.Topic, eventFields.Subtopic, eventFields.Schema, eventFields.PartitionKey, eventFields.Payload)
			require.Nil(t, err)
		})

		t.Run("It publishes a message through the file logger", func(t *testing.T) {
			config := getConfig()
			config.Port = 0
			client, err := NewClient(config)
			require.Nil(t, err)

			eventFields := GetTestEventFields()
			err = client.Publish(eventFields.Origin, eventFields.Topic, eventFields.Subtopic, eventFields.Schema, eventFields.PartitionKey, eventFields.Payload)
			require.Nil(t, err)
		})

		t.Run("It returns a new payload error", func(t *testing.T) {
			config := getConfig()
			client, err := NewClient(config)
			require.Nil(t, err)

			eventFields := GetTestEventFields()
			eventFields.Payload = []byte("{\"url\" \"https://www.redbubble.com/people/wytrab8/works/26039653-toadally-rad\"}")
			err = client.Publish(eventFields.Origin, eventFields.Topic, eventFields.Subtopic, eventFields.Schema, eventFields.PartitionKey, eventFields.Payload)
			_, ok := err.(*PayloadError)
			assert.Equal(t, true, ok)
		})
	})

}