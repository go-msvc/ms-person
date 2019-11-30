module github.com/go-msvc/ms-person

go 1.12

require (
	github.com/go-msvc/crud v0.0.0-00010101000000-000000000000
	github.com/go-msvc/errors v0.0.0-20191116111408-1c2c4914594f
	github.com/go-msvc/item v0.0.0-20191130074344-9e1a81e4ec36
	github.com/go-msvc/store v0.0.0-20191130075239-33af5ac1e630
	github.com/pkg/errors v0.8.1
)

replace github.com/go-msvc/store => ../store

replace github.com/go-msvc/crud => ../crud
