package megaphone

import "fmt"

// PublicationError is returned when an event couldn't be published.
type PublicationError struct {
	e     error
	event string
}

func (pe *PublicationError) Error() string {
	return fmt.Sprintf("The following event couldn't be published: error: %s, event: %v", pe.e, pe.event)
}

// NewPublicationError return a PublicationError for a given event.
func NewPublicationError(e error, event string) *PublicationError {
	return &PublicationError{
		e:     e,
		event: event,
	}
}

// ConfigError is returned when the configuration provided for a Client is invalid.
type ConfigError struct {
	e     error
	field string
}

func (ce *ConfigError) Error() string {
	if ce.field == "" {
		return fmt.Sprintf("The configuration of the megaphone client is not valid: error: %s", ce.e)
	}
	return fmt.Sprintf("The configuration of the megaphone client is not valid: field: %s, error: %s", ce.field, ce.e)
}

// NewConfigErrorWithField returns a ConfigError with details about the invalid Config field.
func NewConfigErrorWithField(e error, field string) *ConfigError {
	return &ConfigError{
		e:     e,
		field: field,
	}
}

// NewConfigError returns a general ConfigError.
func NewConfigError(e error) *ConfigError {
	return &ConfigError{
		e: e,
	}
}

// PayloadError is returned when the event payload to be published is invalid.
type PayloadError struct {
	e       error
	payload string
}

func (pe *PayloadError) Error() string {
	return fmt.Sprintf("Cannot publish the message, invalid payload: error: %s, event: %v", pe.e, pe.payload)
}

// NewPayloadError returns a PayloadError for a given event payload.
func NewPayloadError(e error, payload string) *PayloadError {
	return &PayloadError{
		e:       e,
		payload: payload,
	}
}
