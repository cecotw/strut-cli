#!/bin/node
const program = require('commander');
const process = require('process');
const colors = require('colors');
const { Product } = require('./libs/product');
const createPrompt = require('./libs/prompts/createPrompt');
const addApplicationPrompt = require('./libs/prompts/addApplicationPrompt');
const addProviderPrompt = require('./libs/prompts/addProviderPrompt');
const product = new Product();

console.log(colors.blue('Welcome to Strut!'));

async function main () {
  await product.init();

  program.version(require('./package.json').version);

  program
    .command('create [name]')
    .description('Create a new product')
    .action(async (name) => {
      await createPrompt(product, name);
      process.chdir(`./${product.name}`);
      await addApplicationPrompt(product);
    });

  program
    .command('add <type> [value]')
    .description('Add an <application|provider> to the product')
    .action(async (type, value) => {
      switch (type) {
        case 'application':
          await addApplicationPrompt(product);
          break;
        case 'provider':
          await addProviderPrompt(product, value);
          break;
        default:
          console.log(colors.red(`'${type}' is not a valid type, try --help for valid commands`));
          break;
      }
    });

  program.parse(process.argv);
}

main();
