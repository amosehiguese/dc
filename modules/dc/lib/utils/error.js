"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.loadErrorHandlers = void 0;
var secrets_1 = require("./secrets");
var logger_1 = require("./logger");
function loadErrorHandlers(app) {
    // catches 404 errors and forward to error handler
    app.use(function (req, res, next) {
        var err = new Error("Not Found");
        err.status = 404;
        next(err);
    });
    app.use(function (err, req, res, next) {
        if (err.name === "ValidationError") {
            return res.status(422).json({
                errors: Object.keys(err.errors).reduce(function (errors, key) {
                    errors[key] = err.errors[key].message;
                    return errors;
                }, {})
            });
        }
        logger_1.default.error(err);
        res.status(err.status || 500);
        res.json({
            errors: {
                message: err.message,
                error: !secrets_1.IS_PRODUCTION ? err : {},
            }
        });
    });
}
exports.loadErrorHandlers = loadErrorHandlers;
