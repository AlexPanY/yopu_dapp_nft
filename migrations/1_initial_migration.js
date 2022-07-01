const Migrations = artifacts.require("Migrations");
const YopuToken = artifacts.require("YopuToken");

module.exports = function (deployer) {
  deployer.deploy(Migrations);
  deployer.deploy(YopuToken);
};
