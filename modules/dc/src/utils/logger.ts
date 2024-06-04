import { createLogger, format, transports } from "winston";
import * as fs from 'fs';
import DailyRotateFile = require("winston-daily-rotate-file");
import {ENVIRONMENT, LOG_DIR} from "./secrets";

let dir = LOG_DIR;

if (!fs.existsSync(dir)) {
  fs.mkdirSync(dir)
}

const logLevel = ENVIRONMENT === "dev" ? "debug" : "warn";

const options = {
  file: {
    level                           : logLevel,
    filename                        : dir + "/%DATE%.log",
    datePattern                     : "YYYY-MM-DD",
    zippedArchive                   : true,
    timeStamp                       : true,
    handleExceptions                : true,
    humanReadableUnhandledException : true,
    prettyPrint                     : true,
    json                            : true,
    maxSize                         : "20m",
    colorSize                       : true,
    maxFiles                        : "14d",
  },
};

export default createLogger({
  transports: [
    new transports.Console({
      stderrLevels: ["info", "error"],
      format      : format.combine(format.errors({stack: true}), format.prettyPrint()),
    }),
  ],
  exceptionHandlers: [new DailyRotateFile(options.file)],
  exitOnError      : false,
})

