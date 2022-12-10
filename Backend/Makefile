run:
	docker-compose -f ./build/docker/docker-compose.yml up --build

swagger:
	swagger generate spec -o ./swagger.json

serve-swagger:
	swagger serve -F=swagger swagger.json