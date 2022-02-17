require (
	db v0.0.0
	github.com/google/renameio v0.1.0 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/kisielk/gotool v1.0.0 // indirect
	golang.org/x/tools/gopls v0.7.4 // indirect
	modtest v0.0.0
	routers v0.0.0
)

replace routers v0.0.0 => ./routers

replace db v0.0.0 => ./db

replace modtest v0.0.0 => ./modtest

module test

go 1.13
