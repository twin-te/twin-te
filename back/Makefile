gorm-gen:
	rm -rf ./db/gen && go run ./db/gorm_gen.go

buf-gen:
	rm -rf ./handler/api/rpcgen && buf generate --template ./buf.gen.yaml ../proto
