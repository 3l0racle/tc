/*
 *Entry point for K-Coin
*/

package main

import (
  "os"
  "github.com/3l0racle/kcoin/cli"
)

func main() {
	defer os.Exit(0)
  cmd := cli.CommandLine{}
	cmd.Run()

}
