server:
	docker run --rm -it -p 8080:8080 $(shell docker build -q -f server.Dockerfile .)

client:
	docker run --rm -it --net=host $(shell docker build -q -f client.Dockerfile .)
