ent-generate:
	go run -mod=mod entgo.io/ent/cmd/ent generate --target ./ent ./internal/schema/ --feature sql/modifier --feature sql/execquery --feature sql/versioned-migration

wire:
	wire ./internal/wire

dev:
	env GO_ENV=DEV wails dev

build_windows:
	rm -rf ./resources/database.db && rm -rf ./resources/databaseV034.db && rm -rf ./resources/log.log && env OS=window go run ./cmd/copy/main.go && env GOOS="windows" GOARCH="amd64" CGO_ENABLED="1" CC="x86_64-w64-mingw32-gcc" wails build -clean -ldflags="-s -w" -skipbindings --nsis

build_mac:
	rm -rf ./resources/database.db && rm -rf ./resources/databaseV034.db && rm -rf ./resources/log.log && env CGO_ENABLED="1" wails build -clean && env OS=darwin go run ./cmd/copy/main.go && productbuild --component ./build/bin/AllRounderKim.app/ /Applications/ ./AllrounderKim.pkg