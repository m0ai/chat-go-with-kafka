show-kafka-info:
	/usr/local/bin/kafkacat -L -b localhost:9092

open-kafka-topic-ui:
	open http://localhost:8000/#/

open-zookeeper-navigator:
	open http://localhost:8004/editor/data

logs:
	docker-compose logs -f
