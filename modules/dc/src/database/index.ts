import { Sequelize } from "sequelize";
import logger from "../utils/logger";
import { DB } from "../utils/secrets";

const dbURI = `postgres://${DB.USER}:${DB.PASSWORD}@${DB.HOST}:${DB.PORT}/${DB.NAME}`

const sequelizeConnection = new Sequelize(dbURI)
const initModels = (sequelize: Sequelize) => {

}

// const { _ } = initModels(sequelizeConnection)

export const initDB = async () => {
  try {
    await sequelizeConnection.authenticate();
    await sequelizeConnection.sync({force: false})
    logger.info("DB successfully connected.")
  } catch(error) {
    logger.error("An error occured while initializing database.")
    process.exit(1);
  }
}
