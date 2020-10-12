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

// the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds the Docxy website",
	Long: `Builds the Docxy website, ready for deployment to
production.

Example:
  docxy build
`,
	Run: func(cmd *cobra.Command, args []string) {
		// check whether `.docxy` directory exists
		if !IsDir(".docxy") {
			fmt.Println("error: this is not a docxy project - https://docxy.traction.one")
			os.Exit(126)
		}

		// change working directory to `.docxy`
		Cd(".docxy")

		// build the docxy site
		ExecNode("run", "build")

		// remove the `build` directory
		RmRf(Resolve("../build/"))

		// create symlink for the `build` directory
		Ln(Resolve("./public/"), Resolve("../build/"))
	},
}

func init() {
	// initialize build command
	rootCmd.AddCommand(buildCmd)
}
