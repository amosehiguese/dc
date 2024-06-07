import { Application } from "express";
import * as bodyParser from "body-parser";
import { loadErrorHandlers } from "./utils/error";
import helmet from "helmet";
import { SESSION_SECRET } from "./utils/secrets";
import "./database";
import User from "./database/models/user";
import logger from "./utils/logger";

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
app.get("/", (req, res)=>{
  res.send("Hello world")
})
app.get("/api/v1", async (req, res)=>{

  const user = await User.findOrCreate({
    where: {email:"amosehiguese@gmail.com"},

  })

  user && logger.info("User created successfully")
  res.json(user)
});

loadErrorHandlers(app)

export default app;
