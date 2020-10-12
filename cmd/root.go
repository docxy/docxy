/*
	Copyright © 2020 Sankarsan Kampa

	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at https://mozilla.org/MPL/2.0/.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// the base command
var rootCmd = &cobra.Command{
	Use:   "docxy",
	Short: "The Docxy CLI",
	Long: `This is a CLI for Docxy.

Docxy is a React based open-source documentation site
generator. Build beautiful, blazing fast documentation sites
for your projects with just markdown.

https://docxy.traction.one`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
