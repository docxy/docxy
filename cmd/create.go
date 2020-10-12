/*
	Copyright © 2020 Sankarsan Kampa

	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at https://mozilla.org/MPL/2.0/.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a Docxy project",
	Long: `Creates a Docxy project in the specified directory.

Example:
  docxy create my-awesome-docs
`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectDir := Resolve(args[0])

		fmt.Println(projectDir)

		// get docxy starter
		Exec("git", "clone", "--single-branch", "--branch", "main", "--depth", "1", "https://github.com/docxy/starter", args[0])

		// change working directory to the project
		Cd(projectDir)

		// remove the `.git` directory
		RmRf(".git")

		// initialize a git repository
		Exec("git", "init")

		// initialize a git repository
		Exec("git", "add", ".")

		// initial commit
		Exec("git", "commit", "-m", `initialized Docxy site

Ref: https://docxy.traction.one
`)
	},
}

func init() {
	// initialize create command
	rootCmd.AddCommand(createCmd)
}
