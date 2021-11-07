build_image:
	docker build -t dtrondoli/my-finance-server .

run_app:
	docker-compose -f .devops/app.yml up -d

lint:
	golint ./...
	go fmt -n ./...

unit_test:
	go test ./...