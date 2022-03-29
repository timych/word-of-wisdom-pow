build:
	docker build . -f Dockerfile.server --tag word-of-wisdom-server
	docker build . -f Dockerfile.client --tag word-of-wisdom-client
run-server:
	docker network create wow-net || true
	docker run --rm --net wow-net --name wow-server word-of-wisdom-server
run-client:
	docker run --rm --net wow-net word-of-wisdom-client
