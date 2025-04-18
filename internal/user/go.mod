module github.com/Ellan98/ding-water-service/user

replace github.com/Ellan98/ding-water-service/common => ../common

go 1.24.1

require github.com/Ellan98/ding-water-service/common v0.0.0-00010101000000-000000000000

require (
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/rs/zerolog v1.34.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
)
