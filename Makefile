dev:
	air .

gen:
	go generate ./ent

gqlgen:
	go run github.com/99designs/gqlgen