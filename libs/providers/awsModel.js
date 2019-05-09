const { readFile } = require('../utils');
const path = require('path');
const CloudFormation = require('aws-sdk/clients/cloudformation');
const { ProviderModel } = require('./baseProviderModel');
const cloudformation = new CloudFormation();

class AwsModel extends ProviderModel {
  async init() {
    this.providerName = this.application.providers.AWS.name || 'AWS';
    this.infrastructure = this.application.providers.AWS.infrastructure;
    this.infrastructureFiles = await Promise.all(this.infrastructure.map(
      resource => {
        return readFile(path.join(this.application.path, resource.path));
      }));
    this.infrastructureData = this.infrastructure.map((resource, i) => {
      return { ...resource, fileData: this.infrastructureFiles[i] };
    });
    console.log(this.infrastructureData);
  }
};

module.exports = {
  AwsModel
};
