module teste

go 1.15

replace connection/apiConnection => ../apiConnection

replace term/front => ../front

require (
	connection/apiConnection v0.0.0-00010101000000-000000000000
	github.com/gizak/termui/v3 v3.1.0 // indirect
	term/front v0.0.0-00010101000000-000000000000
)
