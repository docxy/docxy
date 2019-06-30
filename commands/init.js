const fs = require("fs");
const path = require("path");
const shell = require("shelljs");

exports.command = "init";

exports.describe = "Initializes an AwesomeDocs project";

exports.builder = () => {};

exports.handler = function (argv) {
  // Check whether `git` is installed
  if (!shell.which("git")) {
    shell.echo("Error: `git` is required to use AwesomeDocs");
    shell.exit(1);
  }

  // Clone AwesomeDocs from its repo
  if (shell.exec("git clone --depth 1 https://github.com/AwesomeDocs/AwesomeDocs .awesome").code !== 0) {
    shell.echo("Error: Unable to fetch AwesommeDocs");
    shell.exit(1);
  }

  // Move shell runtime inside `.awesome` directory.
  shell.cd(".awesome");

  // Remove .git directory
  shell.rm("-rf", ".git");

  // Install Dependencies
  if (shell.exec("yarn").code !== 0) {
    shell.echo("Error: Unable to install AwesomeDocs dependencies");
    shell.exit(1);
  }

  // Go back to main directory
  shell.cd("..");

  // Create symlinks for internal directories
  let contentDir = path.resolve("content");
  let staticDir = path.resolve("static");

  if (fs.existsSync(contentDir) && !fs.lstatSync(contentDir).isSymbolicLink()) {
    shell.rm("-rf", ".awesome/content");
    shell.ln("-sf", contentDir, ".awesome/content/");
  } else {
    shell.ln("-sf", ".awesome/content/", "content/");
  }
  if (fs.existsSync(staticDir) && !fs.lstatSync(staticDir).isSymbolicLink()) {
    shell.rm("-rf", ".awesome/static");
    shell.ln("-sf", staticDir, ".awesome/static/");
  } else {
    shell.ln("-sf", ".awesome/static/", "static/");
  }

  shell.echo("Success: Initialized AwesomeDocs");
};
