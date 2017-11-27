package megaphone

import "fmt"

type PublicationError struct {
	e     error
	event string
}

func (pe *PublicationError) Error() string {
	return fmt.Sprintf("The following event couldn't be published: error: %s, event: %v", pe.e, pe.event)
}

func NewPublicationError(e error, event string) *PublicationError {
	return &PublicationError{
		e:     e,
		event: event,
	}
}

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

func NewConfigErrorWithField(e error, field string) *ConfigError {
	return &ConfigError{
		e:     e,
		field: field,
	}
}

func NewConfigError(e error) *ConfigError {
	return &ConfigError{
		e: e,
	}
}

type PayloadError struct {
	e       error
	payload string
}

func (pe *PayloadError) Error() string {
	return fmt.Sprintf("Cannot publish the message, invalid payload: error: %s, event: %v", pe.e, pe.payload)
}

func NewPayloadError(e error, payload string) *PayloadError {
	return &PayloadError{
		e:       e,
		payload: payload,
	}
}
