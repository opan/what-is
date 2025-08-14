# Kafka v3

## Using docker-compose

```
docker-compose up -d

# to access the kafka-ui, open the url below in browser
http://localhost:8080

# login to kafka broker

docker compose exec kafka-<number> sh

# or
docker compose exec kafka-<number> /bin/bash

# once inside, you can find the kafka CLI in /opt/bitnami/kafka/bin

cd /opt/bitnami/kafka/bin

# try one of the kafka command
./kafka-topics.sh --bootstrap-server localhost:9092 --list

```

official docker image from apache:
https://github.com/apache/kafka/blob/trunk/docker/examples/README.md


## Setup kafka cli/client

Docker image source: https://hub.docker.com/r/bitnami/kafka

### Topics

Kafka topic naming convention: https://cnr.sh/essays/how-paint-bike-shed-kafka-topic-naming-conventions

### Partition

Producer can only write a topic to a partition leader (default behaviour)
Partition can have many replica, called: partition replication factor
All replica of partition that successfully sync with the partition leader called, in-sync replica (ISR)

Consumer can only consumer a topic from a partition leader (default behaviour)
Since kafka 2.4+, it is possible to configure consumer to consumer a topic from the closest replica

Changing the number of partition in the middle is dangerous because its braking the ordering guarantee.

### Producer ACKs

producer can choose to receieve acks of data writes:
1. acks=0, no wait for ack (higher chance of data loss)
2. acks=1, wait for ack from leader (limited data loss)
3. acks=all, wait for ack from all partition leader and replicas (higher durability, but possibly increasing in latency)

### Kafka topic durability

Rules: for a replication factor of N, you can lose up to N-1 brokers and still safe.

e.g: replication factor of 3, can maintain the data from 2 brokers loss.

### Zookeeper (deprecated in kafka v3+)

- To manage kafka brokers.
- To help in leader election for partitions.
- To send notifications to any changes happen to kafka brokers.
- Operate in odd numbers, 1, 3, 5, 7

### Kafka producer

creating a topic

```
./kafka-topics.sh --bootstrap-server localhost:9092 --topic <topic-name> --partitions 3 --create
```

```
# --producer.config is optional
# standard command
./kafka-console-producer.sh --producer.config ../config/producer.properties --bootstrap-server localhost:9092 --topic first_topic

# with property, acks=all
./kafka-console-producer.sh --producer.config ../config/producer.properties --bootstrap-server localhost:9092 --topic first_topic --producer-property acks=all

# with more property to specify key message
./kafka-console-producer.sh --producer.config ../config/producer.properties --bootstrap-server localhost:9092 --topic first_topic --property parse.key=true --property key.separator=:


# specify producer-property with round-robin parititioner
./kafka-console-producer.sh --producer.config ../config/producer.properties --bootstrap-server localhost:9092 --topic first_topic --producer-property partitioner.class=org.apache.kafka.clients.producer.RoundRobinPartitioner
```

### Kafka consumer

```
# --consumer.config is optional
# default, consume only the latest messsage, only message published when the command is executed
./kafka-console-consumer.sh --consumer.config ../config/consumer.properties --bootstrap-server localhost:9092 --topic first_topic

# read all messages from beginning
./kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic first_topic --from-beginning

# specify --consumer-property, e.g: to set group.id
./kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic first_topic --from-beginning --consumer-property group.id=test-consumer-group

```

### Kafka consumer group

```
# list all consumer groups available
./kafka-consumer-groups.sh --bootstrap-server localhost:9092 --list

# describe a consumer group
./kafka-consumer-groups.sh --bootstrap-server localhost:9092 --describe --group <consumer-group-name>

# reset offset in a consumer group
# --dry-run is used to test the command before actually applying the command
./kafka-consumer-groups.sh --bootstrap-server localhost:9092 --group <consumer-group-name> --reset-offsets --to-earliest --topic <topic-name> --dry-run

# --execute is used to actually apply the command
./kafka-consumer-groups.sh --bootstrap-server localhost:9092 --group <consumer-group-name> --reset-offsets --to-earliest --topic <topic-name> --execute
```

### Kafka consumer offset reset behaviour

```
auto.offset.reset=latest # will read from the end of the log, or latest
auto.offset.reset=earliest # will read from the start of the log
auto.offset.reset=none # will throw an exception if no offset is found

```

default consumer offset retention:

- Kafka < 2.0: 1 day
- Kafka >= 2.0: 7 day

consumer offset retention config:
```
offset.retention.minutes
```

to replay data for consumer group:
- Take all consumers from a specific group down
- User `kafka-consumer-groups` command to set the target offset
- Restart consumers


### Consumer heartbeat thread

```
# specify how often the consumer groups send the heartbeats to consumer
# usually 1/3rd of `session.timeout.ms`
heartbeat.interval.ms # default 3ms


# specify how often to send heartbeat to broker 
# if no heartbeat send during this time period, then the consumer is considered dead
session.timeout.ms # default 45s for v3+, 10s for < v3

```

### Consumer poll thread

```
# define max amount between two poll calls before declaring that the consumer are dead
# to check issue with consumer, e.g: consumer is stuck
max.poll.interval.ms # default 5min 

# to control how many records to receive per poll request
max.poll.records # default 500

# control how much data to pull at least on each request
# to improve throughput and decrease request number with the cost of latency
fetch.min.bytes # default 1

# set max amount of time for kafka broker will block before answering the fetch request if there isn't sufficient data
# to immediately satisfy the requirement given by fetch.min.bytes
fetch.max.wait.ms # default 500

# max amount of data per partition the server will return
# the more partition, the more memory consumed
max.partition.fetch.bytes # default 1MB

# max data returned for each fetch request
fetch.max.bytes
```

### Consumer rack awareness

Must use kafka v2.4+

Broker config:
- config `rack.id` must be set to the data center ID, e.g: AZ ID in AWS
- `replica.selector.class` must be set to `org.apache.kafka.common.replica.RackAwareReplicaSelector`

Consumer config:
- set `rack.id` the same ID with the one set in the broker
