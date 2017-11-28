// Package logger provides a logger that the Megaphone client
// can use to publish to a local file or to Megaphone via Fluentd.
package logger

// Logger is a generic mecanism for megaphone.Client to post events.
type Logger interface {
	Post(topic string, message interface{}) error
}

// NewLogger returns a new Logger to be used by a megaphone.Client.
// Conventionally, a FileLogger is returned if no configuration for
// a FluentLogger is provided, allowing for easy local development
// in absence of a local Megaphone Fluentd container.
func NewLogger(host string, port int) (Logger, error) {
	if host != "" && port > 0 {
		fluentLogger, err := NewFluentLogger(host, port)
		if err != nil {
			return nil, err
		}
		return fluentLogger, nil
	}

	return NewFileLogger()
}
