#!/usr/bin/env node

const yargs = require("yargs")

const loadCommamnd = (command) => require("./commands/" + command);

yargs
    .argv;
