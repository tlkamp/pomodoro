.PHONY:
build: clean
	@go build -o pomodoro-cli

.PHONY:
clean:
	@rm -f pomodoro-cli
