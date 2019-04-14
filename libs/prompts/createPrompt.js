const inquirer = require('inquirer');

async function createPrompt (productModel, name) {
  let createPrompt = await inquirer.prompt([{
    type: 'input',
    name: 'name',
    message: 'Enter new strut product:',
    when () { return !name; }
  }]);
  name = name || createPrompt.name;
  await productModel.create(name);
};

module.exports = createPrompt;