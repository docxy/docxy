const shell = require("shelljs");

exports.command = "build";

exports.describe = "Builds the documentation website";

exports.builder = () => {};

exports.handler = function (argv) {
  // Move shell runtime inside `.awesome` directory.
  if (shell.cd(".awesome").code !== 0) {
    shell.echo("Hint: Run `awesomedocs init` to initialize AwesomeDocs");
    shell.exit(1);
  }

  // Build the documentation webisite
  shell.exec("gatsby build");

  // Go back to main directory
  shell.cd("..");

  // Create symlink for build directory
  shell.ln("-sf", ".awesome/public/", "build/");
};
