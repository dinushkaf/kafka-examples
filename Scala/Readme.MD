# Introduction
This will demonstrate how to use the Kafka client with Scala. You will need to install [Scala](https://www.scala-lang.org/download/) first and [sbt](https://www.scala-sbt.org/0.13/docs/Installing-sbt-on-Linux.html) also.

## Steps
* Create 3 files named *build.sbt*, *producer.scala*, *consumer.scala* in your folder.
* Make the 3 files as I have given in the source code.
* Open a new terminal window in your folder and run sbt.`sbt`
* Run the Kafka server
* Then run compile and then run `> compile`, `> run`. You can select enter the relevent key to run consumer.
* Open another terminal window in your folder and run `sbt`, then run `> run`. Enter relevent key for producer.
* Now you will see producer publishing data and Consumer receiving it.

## *build.sbt* file
* This file contains the relevent configurations need to run your code. Once you have changed the content of this file, you need to re-open the sbt shell.
* Here I have given the resolver and library dependency for our apache kafka library.

## *Producer.scala* file
* This file contains the code for Kafka Producer.
* You can change the host and the port according to your need.
`props.put("bootstrap.servers", "localhost:9092")`

## *Consumer.scala* file
* This file contains the code for Kafka Consumer.
* You can change the host and the port according to your need.
`props.put("bootstrap.servers", "localhost:9092")`

## References
* https://index.scala-lang.org/cakesolutions/scala-kafka-client/scala-kafka-client-akka/2.0.0?target=_2.12
* https://www.scala-sbt.org/0.12.4/docs/Getting-Started/Running.html
* https://gist.github.com/fancellu/f78e11b1808db2727d76
