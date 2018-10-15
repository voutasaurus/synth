all:
	@./cmd/pub/dockerbuild.sh
	@./cmd/sub/dockerbuild.sh
	@docker-compose up
