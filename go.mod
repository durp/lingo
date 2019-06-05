module github.com/chewxy/lingo

go 1.12

require (
	github.com/abiosoft/ishell v2.0.0+incompatible
	github.com/abiosoft/readline v0.0.0-20180607040430-155bce2042db // indirect
	github.com/awalterschulze/gographviz v0.0.0-20190522210029-fa59802746ab
	github.com/fatih/color v1.7.0 // indirect
	github.com/flynn-archive/go-shlex v0.0.0-20150515145356-3f9db97f8568 // indirect
	github.com/kljensen/snowball v0.6.0
	github.com/leesper/go_rng v0.0.0-20190531154944-a612b043e353 // indirect
	github.com/mattn/go-colorable v0.1.2 // indirect
	github.com/pkg/browser v0.0.0-20180916011732-0a3d74bf9ce4
	github.com/pkg/errors v0.8.1
	github.com/stretchr/testify v1.3.0
	github.com/xtgo/set v1.0.0
	golang.org/x/text v0.3.2
	gonum.org/v1/gonum v0.0.0-20190602094245-c4a599f7b712 // indirect
	gorgonia.org/gorgonia v0.9.2-0.20190605194449-da60c18a067a
	gorgonia.org/tensor v0.9.0-beta
)

replace gorgonia.org/gorgonia => github.com/durp/gorgonia v0.9.2-0.20190605194449-da60c18a067a
