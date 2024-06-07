"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.DB = exports.SESSION_SECRET = exports.JWT_SECRET = exports.LOG_DIR = exports.APP_PORT = exports.IS_PRODUCTION = exports.ENVIRONMENT = void 0;
var dotenv = require("dotenv");
var _ = require("lodash");
var path = require("path");
dotenv.config({
    path: ".env",
});
exports.ENVIRONMENT = _.defaultTo(process.env.APP_ENV, "dev");
exports.IS_PRODUCTION = exports.ENVIRONMENT === "production";
exports.APP_PORT = _.defaultTo(parseInt(process.env.APP_PORT), 3000);
exports.LOG_DIR = _.defaultTo(process.env.LOG_DIR, path.resolve('logs'));
exports.JWT_SECRET = _.defaultTo(process.env.JWT_SECRET, "secret");
exports.SESSION_SECRET = _.defaultTo(process.env.SESSION_SECRET, "secret");
exports.DB = {
    USER: _.defaultTo(process.env.DB_USER, "dc"),
    PASSWORD: _.defaultTo(process.env.DB_PWD, "secret"),
    HOST: _.defaultTo(process.env.DB_HOST, "localhost"),
    NAME: _.defaultTo(process.env.DB_NAME, "dcdb"),
    PORT: _.defaultTo(parseInt(process.env.DB_PORT), 27017),
};
