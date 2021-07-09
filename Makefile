test:
	go test ./...

regen:
	(cd entgql && go generate .)
	(cd entgql && go generate ./internal/...)
