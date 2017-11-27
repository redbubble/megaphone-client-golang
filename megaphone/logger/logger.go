package logger

type Logger interface {
	Post(topic string, message interface{}) error
}

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
