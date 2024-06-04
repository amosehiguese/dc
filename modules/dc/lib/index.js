"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var app_1 = require("./app");
var secrets_1 = require("./utils/secrets");
var logger_1 = require("./utils/logger");
app_1.default
    .listen(secrets_1.APP_PORT, function () {
    logger_1.default.info("server running on port: ".concat(secrets_1.APP_PORT));
})
    .on("server", function (e) { return logger_1.default.error(e); });
