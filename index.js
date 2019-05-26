#!/usr/bin/env node

const yargs = require("yargs")

const loadCommamnd = (command) => require("./commands/" + command);

yargs
    .command(loadCommamnd("init"))
    .command(loadCommamnd("build"))
    .command(loadCommamnd("serve"))
    .argv;
