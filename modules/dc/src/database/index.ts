import { Sequelize } from "sequelize";
import logger from "../utils/logger";
import { DB } from "../utils/secrets";

const dbURI = `postgres://${DB.USER}:${DB.PASSWORD}@${DB.HOST}:${DB.PORT}/${DB.NAME}`

export const sequelize = new Sequelize(dbURI)

const initDB = async () => {
  try {
    await sequelize.authenticate();
    await sequelize.sync({force: false})
    logger.info("DB successfully connected.")

  } catch(error) {
    logger.error("An error occured while initializing database.")
    process.exit(1);
  }
}


export default initDB
