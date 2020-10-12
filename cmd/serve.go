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

// the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the development server",
	Long: `Starts the development server and serves the Docxy site so
that you can view it in a browser.

Example:
  docxy serve
  docxy serve --port 1337
`,
	Run: func(cmd *cobra.Command, args []string) {
		// check whether `.docxy` directory exists
		if !IsDir(".docxy") {
			fmt.Println("error: this is not a docxy project - https://docxy.traction.one")
			os.Exit(126)
		}

		// change working directory to `.docxy`
		Cd(".docxy")

		// start the docxy development server
		ExecNode("run", "start", "--port", fmt.Sprintf("%s", cmd.Flag("port").Value))
	},
}

func init() {
	// initialize serve command
	rootCmd.AddCommand(serveCmd)

	// add flags
	serveCmd.PersistentFlags().String("port", "36299", "Port number for the development server")
}
