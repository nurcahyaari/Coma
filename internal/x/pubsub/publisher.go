package pubsub

import (
	"io"
	"strings"
)

type publisher struct {
	message chan io.Reader
}

type publisherOptions struct {
	bufferCapacity int
}

func newPublisher(options publisherOptions) *publisher {
	pub := &publisher{
		message: make(chan io.Reader, options.bufferCapacity),
	}

	return pub
}

func (p *publisher) publish(message MessageHandler) error {
	data, err := message()
	if err != nil {
		return err
	}

	p.message <- data
	return nil
}

func (p *publisher) close() {
	close(p.message)
}

func (p *publisher) retrieveMessages() ([]string, error) {
	messages := []string{}
	for message := range p.message {
		buf := new(strings.Builder)
		_, err := io.Copy(buf, message)
		if err != nil {
			return nil, err
		}

		messages = append(messages, buf.String())
	}

	return messages, nil
}

func (p *publisher) capacity() int {
	return cap(p.message)
}

func (p *publisher) len() int {
	return len(p.message)
}
