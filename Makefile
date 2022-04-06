setup-synadia-%:
	NATS_IMAGE=synadia/nats-server:$* docker-compose up -d --force-recreate

setup-nats-%:
	NATS_IMAGE=nats:$*-alpine docker-compose up -d --force-recreate

teardown-environment:
	docker-compose down -v

run-performance-test:
	go run bench.go run constant -r 10000/s -d 60s -c 5000 jetstream
