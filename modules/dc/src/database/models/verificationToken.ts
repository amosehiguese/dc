import { Model, InferAttributes, InferCreationAttributes, CreationOptional, DataTypes, ForeignKey, UUIDV4} from "sequelize";
import { sequelize } from "../index"
import _User  from "./user";

class VerificationToken extends Model<InferAttributes<VerificationToken>, InferCreationAttributes<VerificationToken>> {
  declare id: CreationOptional<string>;
  declare userId: ForeignKey<_User["id"]>;
  declare token: string;
  declare expiresAt: Date;
  declare createdAt: CreationOptional<Date>;
  declare updatedAt: CreationOptional<Date>;
}

VerificationToken.init(
  {
    id: {
      type: DataTypes.UUID,
      defaultValue: UUIDV4,
      primaryKey: true,
    },
    token: {
      type: DataTypes.STRING,
      allowNull: false,
      unique: true,
    },
    expiresAt: {
      type: DataTypes.DATE,
      allowNull: false,
    },

    createdAt: DataTypes.DATE,
    updatedAt: DataTypes.DATE,
  },
  {
    tableName: "verification_tokens",
    sequelize
  }
)

export default VerificationToken;
