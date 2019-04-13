const fs = require('fs');
const path = require('path');
const colors = require('colors');

exports.ProductModel = class ProductModel {
  get productModel() {
    return {
      name: null,
      version: '0.0.0',
      applications: [],
      providers: []
    };
  };

  get applicationModel () {
    return {
      name: null,
      repository: {
        type: null,
        url: null
      },
      install: [], // install commands
      build: [], // build commands
      artifacts: [], // artifact paths,
      start: [] // start the application
    };
  };

  get providerModel () {
    return {
      name: null,
      applications: []
    };
  }

  get providerApplicationModel () {
    return {
      name: null,
      resources: []
    };
  }

  async init() {
    this.product = await this.loadProduct();
  }

  async loadProduct () {
    const productJson = 'product.json';
    return new Promise((resolve, reject) => {
      fs.stat(productJson, err => {
        if (err && err.code === 'ENOENT') {
          resolve(this.productModel);
        } else {
          fs.readFile(productJson, 'utf8', (err, data) => {
            if (err) throw err;
            resolve(JSON.parse(data));
          });
        }
      });
    });
  }

  updateProductFile (data, dir) {
    let productJsonPath = 'product.json';
    if (dir) {
      productJsonPath = path.resolve(dir, productJsonPath);
    }
    return new Promise((resolve, reject) => {
      fs.writeFile(
        productJsonPath,
        JSON.stringify(data || this.product, null, 2),
        (err) => {
          if (err) {
            reject(err);
          }
          resolve();
        }
      );
    });
  }

  async create (name) {
    console.log(colors.yellow(`Creating product: ${colors.gray(name)}`));
    this.name = name || 'myproduct';
    let dir = `./${name}`;
    if (!fs.existsSync(dir)) {
      fs.mkdirSync(dir);
    }
    await this.updateProductFile({ ...this.productModel, name }, dir);
    console.log(colors.green('DONE!'));
  }

  async addApplication (application) {
    console.log(colors.yellow(`Adding application to product: ${colors.gray(application.name)}`));
    if (!this.product.applications.some(a => a.name === application.name)) {
      this.product.applications.push(application);
      await this.updateProductFile();
      console.log(colors.green('DONE!'));
    } else {
      console.log(colors.red('A name with this application already exists.'));
    }
  }

  async addProvider (provider) {
    console.log(colors.yellow(`Adding provider to product: ${colors.gray(provider.name)}`));
    if (!this.product.providers.some(p => p.name === provider.name)) {
      this.product.providers.push(provider);
      await this.updateProductFile();
      console.log(colors.green('DONE!'));
    } else {
      console.log(colors.red('You already have this provider added'));
    }
  }

  async link (application, provider) {
    console.log(colors.yellow(`Linking ${colors.gray(application.name)} to ${colors.gray(provider.name)}`));
    if (provider.applications.some(a => a.name === application.name)) {
      console.log(colors.red('Application already linked to provider'));
    } else {
      this.product.providers
        .find(p => p.name === provider.name)
        .applications.push({ ...this.providerApplicationModel, name: application.name });
      await this.updateProductFile();
      console.log(colors.green('DONE!'));
    }
  }
};
