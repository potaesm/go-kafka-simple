# Go Kafka Simple

## Install and run Kafka
```bash
# Download Kafka
https://archive.apache.org/dist/kafka/3.4.0/kafka_2.13-3.4.0.tgz
cd kafka_2.13-3.4.0
# Run Zookeeper
bin/zookeeper-server-start.sh config/zookeeper.properties
# Run Kafka
bin/kafka-server-start.sh config/server.properties
```

## Run main
```bash
go mod download
go run main.go
```