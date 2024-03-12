# Kafka v3

## Using docker-compose

```
docker-compose up -d

```

Docker image source: https://hub.docker.com/r/bitnami/kafka

### Partition

Producer can only write a topic to a partition leader (default behaviour)
Partition can have many replica, called: partition replication factor
All replica of partition that successfully sync with the partition leader called, in-sync replica (ISR)

Consumer can only consumer a topic from a partition leader (default behaviour)
Since kafka 2.4+, it is possible to configure consumer to consumer a topic from the closest replica

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