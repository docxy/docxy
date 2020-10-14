/*
	Copyright © 2020 Sankarsan Kampa

	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at https://mozilla.org/MPL/2.0/.
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a Docxy project",
	Long: `Initializes a Docxy project in a directory.

Example:
  docxy init
`,
	Run: func(cmd *cobra.Command, args []string) {
		// remove the `.docxy` directory
		RmRf(".docxy")

		// create the `contents` directory, if it doesn't exist
		if !IsDir("./contents") {
			MkDir("./contents")
		}

		// create the `static` directory, if it doesn't exist
		if !IsDir("./static") {
			MkDir("./static")
		}

		// get docxy
		Exec("git", "clone", "--single-branch", "--branch", "main", "--depth", "1", "https://github.com/docxy/docgen", ".docxy")

		// change working directory to `.docxy`
		Cd(".docxy")

		// remove the `.git` directory
		RmRf(".git")

		// create symlinks for the `contents` & `static` directories
		Ln(Resolve("../contents"), Resolve("./contents"))
		Ln(Resolve("../static"), Resolve("./static"))

		// install dependencies
		ExecNode("install")
	},
}

func init() {
	// initialize init command
	rootCmd.AddCommand(initCmd)
}
