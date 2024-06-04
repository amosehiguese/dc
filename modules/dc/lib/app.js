"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var bodyParser = require("body-parser");
var error_1 = require("./utils/error");
var helmet_1 = require("helmet");
var secrets_1 = require("./utils/secrets");
require("./database");
require("./utils/passport");
var express = require("express");
var session = require("express-session");
var compression = require("compression");
var app = express();
app.use((0, helmet_1.default)());
app.use(compression());
app.use(bodyParser.json());
app.use(session({
    secret: secrets_1.SESSION_SECRET,
    cookie: {
        maxAge: 60000,
    },
    resave: false,
    saveUninitialized: false
}));
// app.use("/api/v1", mainRouter);
(0, error_1.loadErrorHandlers)(app);
exports.default = app;
