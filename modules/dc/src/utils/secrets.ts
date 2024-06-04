import * as dotenv from "dotenv";
import * as _ from "lodash";
import * as path from "path";

dotenv.config({
  path: ".env",
});

export const ENVIRONMENT = _.defaultTo(process.env.APP_ENV, "dev");
export const IS_PRODUCTION = ENVIRONMENT === "production";
export const APP_PORT = _.defaultTo(parseInt(process.env.APP_PORT as string), 3000);
export const LOG_DIR = _.defaultTo(process.env.LOG_DIR, path.resolve('logs'));
export const JWT_SECRET = _.defaultTo(process.env.JWT_SECRET, "secret");
export const SESSION_SECRET = _.defaultTo(process.env.SESSION_SECRET, "secret");
export const DB = {
  USER : _.defaultTo(process.env.DB_USER, "dc"),
  PASSWORD : _.defaultTo(process.env.DB_USER_PWD, "secret"),
  HOST : _.defaultTo(process.env.DB_HOST, "localhost"),
  NAME : _.defaultTo(process.env.DB_NAME, "dcdb"),
  PORT : _.defaultTo(parseInt(process.env.DB_PORT as string), 27017),
}



