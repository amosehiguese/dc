"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var winston_1 = require("winston");
var fs = require("fs");
var DailyRotateFile = require("winston-daily-rotate-file");
var secrets_1 = require("./secrets");
var dir = secrets_1.LOG_DIR;
if (!fs.existsSync(dir)) {
    fs.mkdirSync(dir);
}
var logLevel = secrets_1.ENVIRONMENT === "dev" ? "debug" : "warn";
var options = {
    file: {
        level: logLevel,
        filename: dir + "/%DATE%.log",
        datePattern: "YYYY-MM-DD",
        zippedArchive: true,
        timeStamp: true,
        handleExceptions: true,
        humanReadableUnhandledException: true,
        prettyPrint: true,
        json: true,
        maxSize: "20m",
        colorSize: true,
        maxFiles: "14d",
    },
};
exports.default = (0, winston_1.createLogger)({
    transports: [
        new winston_1.transports.Console({
            stderrLevels: ["info", "error"],
            format: winston_1.format.combine(winston_1.format.errors({ stack: true }), winston_1.format.prettyPrint()),
        }),
    ],
    exceptionHandlers: [new DailyRotateFile(options.file)],
    exitOnError: false,
});
