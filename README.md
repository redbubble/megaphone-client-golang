# Megaphone Golang client

[![GoDoc](https://godoc.org/github.com/redbubble/megaphone-client-golang/megaphone?status.svg)](https://godoc.org/github.com/redbubble/megaphone-client-golang/megaphone)
[![Go Report Card](https://goreportcard.com/badge/github.com/redbubble/megaphone-client-golang/megaphone)](https://goreportcard.com/report/github.com/redbubble/megaphone-client-golang/megaphone)
[![Build Status](https://travis-ci.org/redbubble/megaphone-client-golang.svg?branch=master)](https://travis-ci.org/redbubble/megaphone-client-golang)

Send events to [Megaphone][megaphone] from [Golang][golang] applications.

[megaphone]: https://github.com/redbubble/megaphone
[golang]: https://golang.org/

## Getting Started

Install the package:

```bash
# Using Go get:
go get github.com/redbubble/megaphone-client-golang/megaphone

# Using govendor:
govendor fetch github.com/redbubble/megaphone-client-golang/megaphone@v1
```

## Usage for Fluentd Client

In order to be as unobstrusive as possible, this client will append events to local files (e.g. `./work-updates.stream`) unless:

* the `MEGAPHONE_FLUENT_HOST` and `MEGAPHONE_FLUENT_PORT` environment variables are set.
* **or** the Fluentd host and port values are passed as arguments to the client's constructor

That behaviour ensures that unless you want to send events to the Megaphone [streams][stream], you do not need to [start Fluentd][megaphone-fluentd] at all.

[stream]: https://github.com/redbubble/com/megaphone#stream
[megaphone-fluentd]: https://github.com/redbubble/megaphone-fluentd-container

### Publishing events

1. Start Fluentd, the easiest way to do so is using a [`redbubble/megaphone-fluentd`][megaphone-fluentd] container

1. Create your event and publish it:

```golang
// Configure a Megaphone client for your awesome service
client, err := megaphone.NewFluentdClient("my-awesome-service", "localhost", 24224)

// Create an event
topic := "work-updates"
subtopic := "work-metadata-updated"
schema := "https://github.com/redbubble/megaphone-event-type-registry/blob/master/streams/work-updates-schema-1.0.0.json"
partitionKey := "1357924680" # the Work ID in this case
payload := "{ url: \"https://www.redbubble.com/people/wytrab8/works/26039653-toadally-rad\" }"

// Publish your event
err := client.Publish(topic, subtopic, schema, partitionKey, payload)
if err != nil {
  // handle the error
}
```

## Usage using Kinesis Publisher (synchronous)

This publisher writes to Kinesis directly, using an AWS Kinesis client.

### Publishing events

1. Create a new client 

```golang

publisher, err := megaphone.NewKinesisSynchronousPublisher(kinesisclient.Config{
  Origin: "my-awesome-service",
  DeployEnv: "test",
  HostedOnAWS: true,
})
```

2. Publish your event
```golang
topic := "work-updates"
subtopic := "work-metadata-updated"
schema := "https://github.com/redbubble/megaphone-event-type-registry/blob/master/streams/work-updates-schema-1.0.0.json"
partitionKey := "1357924680" # the Work ID in this case
payload := "{ url: \"https://www.redbubble.com/people/wytrab8/works/26039653-toadally-rad\" }"

err := publisher.Publish(topic, subtopic, schema, partitionKey, payload)
if err != nil {
  // handle the error
}

```

## Credits

[![](doc/redbubble.png)][redbubble]

This Megaphone Golang client is maintained and funded by [Redbubble][redbubble].

[redbubble]: https://www.redbubble.com

## License

    megaphone
    Copyright (C) 2017 Redbubble

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
