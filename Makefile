IMAGE_NAME := ascii-art-web
IMAGE_TAG := latest
CONTAINER_NAME := asd

build:
	@docker build . -t ${IMAGE_NAME}:${IMAGE_TAG}

run:
	@docker run -p 8080:8080 --name $(IMAGE_NAME)-container $(IMAGE_NAME):$(IMAGE_TAG)

bash:
	@docker run -itd --name ${CONTAINER_NAME} ${IMAGE_NAME}
	@docker exec -it ascii-art-web bash

local:
	@go run .