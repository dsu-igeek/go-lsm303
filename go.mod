module github.com/bskari/go-lsm303

go 1.15

require (
	periph.io/x/conn/v3 v3.6.10
	periph.io/x/host/v3 v3.7.2
	periph.io/x/periph v3.6.8+incompatible
)

replace periph.io/x/conn/v3 => ../../../periph.io/x/conn
