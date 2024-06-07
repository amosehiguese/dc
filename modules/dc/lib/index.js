"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var app_1 = require("./app");
var secrets_1 = require("./utils/secrets");
var logger_1 = require("./utils/logger");
var database_1 = require("./database");
var http = require('http');
var httpServer = http.createServer(app_1.default);
(0, database_1.initDB)()
    .then(function () {
    httpServer
        .listen(secrets_1.APP_PORT, function () {
        logger_1.default.info("server running on port: ".concat(secrets_1.APP_PORT));
    })
        .on("server", function (e) { return logger_1.default.error(e); });
});
