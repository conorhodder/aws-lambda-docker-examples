start:
	docker compose watch

simulate:
	@curl "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"foo":"bar"}'
	@echo "\n" # no newline returned by default from AWS
	@curl "http://localhost:9001/2015-03-31/functions/function/invocations" -d '{"baz":"quux"}'