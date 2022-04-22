/*
Copyright © 2022 BORA KASMER <bora@borakasmer.com>

*/
//https://towardsdatascience.com/how-to-create-a-cli-in-golang-with-cobra-d729641c7177

//go install github.com/spf13/cobra-cli@latest
//cd my-calc
//Users/borakasmer/go/bin/cobra-cli init
//Users/borakasmer/go/bin/cobra-cli add add => We added "add" function
// Create "addInt()" in add.go
//*go install exchange-cli
//go env -w GO111MODULE=on NOT NEEEDD

//**** If there is no Go => Install Go https://go.dev/dl/
//go version (Check Version => 1.18 :) )
//go env (Check GOPATH)

//*** change go.mod "module exchange" =>"module github.com/borakasmer/exchange-cli"
//*** change main.go "import exchange/cmd" => "import "github.com/borakasmer/exchange-cli/cmd"
//*** change get.go "import exchange-cli/parser" => "import github.com/borakasmer/exchange-cli/parser"

//**** => go install github.com/borakasmer/exchange-cli@latest
//**** FOR GO.1.18 => MOVE "exchange-cli" and "cobra-cli" from "/Users/borakasmer/go/bin" to => "/usr/local/go/bin"

package main

import "github.com/borakasmer/exchange-cli/cmd"

func main() {
	cmd.Execute()
}
