# Introduction
This will demonstrate how to use the Kafka client with Scala. You will need to install [Scala](https://www.scala-lang.org/download/) first and [sbt](https://www.scala-sbt.org/0.13/docs/Installing-sbt-on-Linux.html) also.

## Dependencies
Install *kafka-go* library. `go get github.com/segmentio/kafka-go`

## Steps
* Install the dependencies and download consumer and producer files.
* Change the host and port in files according to your need. `
* Open a terminal window in you folder and build consumer and producer.<br>`go build -o consumer Consumer.go`<br>`go build -o producer Producer.go`
* Run the Kafka Server.
* Run consumer in a seperate terminal `./consumer`
* Run producer in a seperate terminal `./producer`
* You can see the data comming to consumer terminal

## References
* https://github.com/segmentio/kafka-go