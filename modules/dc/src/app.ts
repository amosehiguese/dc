import { Application } from "express";
import * as bodyParser from "body-parser";
import { loadErrorHandlers } from "./utils/error";
import helmet from "helmet";
import { SESSION_SECRET } from "./utils/secrets";
import "./database";
import "./utils/passport";

const express = require("express");
const session = require("express-session");
const compression = require("compression");

const app: Application = express();

app.use(helmet());
app.use(compression());
app.use(bodyParser.json());
app.use(session({
  secret: SESSION_SECRET,
  cookie: {
    maxAge: 60000,
  },
  resave: false,
  saveUninitialized: false
}));
// app.use("/api/v1", mainRouter);

loadErrorHandlers(app)

export default app;
