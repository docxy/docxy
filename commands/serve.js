const shell = require("shelljs");

exports.command = "serve [port]";

exports.describe = "Serves the documentation in browser so that you can preview it while working on it.";

exports.builder = (yargs) => {
  yargs.positional("port", {
      describe: "port to bind on",
      type: "number",
      default: 8000,
  });
};

exports.handler = function (argv) {
  // Move shell runtime inside `.awesome` directory.
  if (shell.cd(".awesome").code !== 0) {
    shell.echo("Hint: Run `awesomedocs init` to initialize AwesomeDocs");
    shell.exit(1);
  }

  // Start Gatsby development server
  shell.exec("yarn start --port " + argv.port);
};
