import app from "./app";
import { APP_PORT } from "./utils/secrets";
import logger from "./utils/logger";
import { initDB } from "./database";

const http = require('http');
const httpServer = http.createServer(app);
initDB()
  .then(()=>{
    httpServer
      .listen(APP_PORT, () => {
        logger.info(`server running on port: ${APP_PORT}`);
      })
      .on("server", (e: Error)=> logger.error(e));
    }
  )
