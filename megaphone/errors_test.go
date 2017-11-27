package megaphone_test

import (
	"errors"
	"testing"

	"github.com/redbubble/megaphone-client-golang/megaphone"
	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {

	t.Run("PublicationError", func(t *testing.T) {
		t.Run("NewPublicationError()", func(t *testing.T) {
			actualError := megaphone.NewPublicationError(errors.New("fluent#PostWithTime: message must be a map"), "invalid payload")
			assert.Equal(t, "The following event couldn't be published: error: fluent#PostWithTime: message must be a map, event: invalid payload", actualError.Error())
		})
	})

	t.Run("ConfigError", func(t *testing.T) {
		t.Run("NewConfigError()", func(t *testing.T) {
			actualError := megaphone.NewConfigError(errors.New("can't connect to host:port"))
			assert.Equal(t, "The configuration of the megaphone client is not valid: error: can't connect to host:port", actualError.Error())
		})

		t.Run("NewConfigErrorWithField()", func(t *testing.T) {
			actualError := megaphone.NewConfigErrorWithField(errors.New("invalid integer"), "port")
			assert.Equal(t, "The configuration of the megaphone client is not valid: field: port, error: invalid integer", actualError.Error())
		})
	})

	t.Run("PayloadError", func(t *testing.T) {
		t.Run("NewPayloadError()", func(t *testing.T) {
			actualError := megaphone.NewPayloadError(errors.New("JSON invalid format"), "{{}")
			assert.Equal(t, "Cannot publish the message, invalid payload: error: JSON invalid format, event: {{}", actualError.Error())
		})
	})

}
