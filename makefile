#	create new schema
#		make create schema=user
create:
	go run -mod=mod entgo.io/ent/cmd/ent new $(schema)  


#	shows your schemas 
describe:
	go run -mod=mod entgo.io/ent/cmd/ent describe ./ent/schema

#	starts your setup
run:
	go run .

#	generates your schemas again
generate:
	go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema