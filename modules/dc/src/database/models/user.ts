import { Model, InferAttributes, InferCreationAttributes, CreationOptional, DataTypes, UUIDV4} from "sequelize";
import { sequelize }  from "../index"

class User extends Model<InferAttributes<User>, InferCreationAttributes<User>> {
  declare id: CreationOptional<string>;
  declare email: string;
  declare password: string | null;
  declare firstName: string;
  declare lastName: string | null;
  declare role: string;
  declare verified: boolean;
  declare createdAt: CreationOptional<Date>;
  declare updatedAt: CreationOptional<Date>;
}

User.init(
  {
    id: {
      type: DataTypes.UUID,
      defaultValue: UUIDV4,
      primaryKey: true,
    },
    email: {
      type: DataTypes.STRING(128),
      allowNull: false,
      unique: true,
    },
    password: {
      type: DataTypes.STRING(128),
    },
    firstName: {
      type: DataTypes.STRING,
    },
    lastName: {
      type: DataTypes.STRING,
    },
    role: {
      type: DataTypes.ENUM("doctor", "user"),
      defaultValue: "user",
    },
    verified: {
      type: DataTypes.BOOLEAN,
      defaultValue: false,
    },

    createdAt: DataTypes.DATE,
    updatedAt: DataTypes.DATE,
  },
  {
    tableName: "users",
    sequelize
  },
)

export default User;
