schema?=schema
gs:
	go run -mod=mod entgo.io/ent/cmd/ent new $(schema)  

describe:
	go run -mod=mod entgo.io/ent/cmd/ent describe ./ent/schema

run:
	go run .