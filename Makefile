run:
	docker build --tag list-ns:latest  .
	docker run --rm -v ~/.aws:/root/.aws --name list-ns list-ns:latest
