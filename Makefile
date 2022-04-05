setup-environment-%:
	NATS_VERSION=$* docker-compose up -d --force-recreate

teardown-environment:
	docker-compose down -v

run-performance-test:
	go run bench.go run constant -r 10000/s -d 60s -c 5000 jetstream
