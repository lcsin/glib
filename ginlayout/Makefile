gorm:
	@cd cmd && cd gen && go run main.go --config ../../config/config.yaml
wire:
	@cd cmd && cd app && wire
run:
	@cd cmd && cd app && go build -o app.exe && ./app.exe --config ../../config/config.yaml