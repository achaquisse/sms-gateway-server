.PHONY: build_list_pending build_schedule build_update_status clean

all: build_list_pending build_schedule build_update_status

build_list_pending:
	@echo "Building list_pending"
	env GOOS=linux GOARCH=amd64 go build -o .out/list_pending cmd/listpending/main.go
	cd .out && zip list_pending.zip list_pending && rm list_pending

build_schedule:
	@echo "Building schedule"
	env GOOS=linux GOARCH=amd64 go build -o .out/schedule cmd/schedule/main.go
	cd .out && zip schedule.zip schedule && rm schedule

build_update_status:
	@echo "Building update_status"
	env GOOS=linux GOARCH=amd64 go build -o .out/update_status cmd/updatestatus/main.go
	cd .out && zip update_status.zip update_status && rm update_status

tf_apply:
	terraform -chdir=.iac apply -auto-approve

clean:
	@echo "Cleaning up..."
	rm -rf .out
	rm -rf gen