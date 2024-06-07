// import { Model, InferAttributes, InferCreationAttributes, CreationOptional, DataTypes, ForeignKey, UUIDV4} from "sequelize";
// import { sequelize } from "../index"
// import _User  from "./user";

// export class Doctor extends Model<InferAttributes<Doctor>, InferCreationAttributes<Doctor>> {
//   declare id: CreationOptional<string>;
//   declare userId: ForeignKey<_User["id"]>;
//   declare specialization: string;
//   declare availability: string;
//   declare createdAt: CreationOptional<Date>;
//   declare updatedAt: CreationOptional<Date>;
// }

// Doctor.init(
//   {
//     id: {
//       type: DataTypes.UUID,
//       defaultValue: UUIDV4,
//       primaryKey: true,
//     },
//     specialization: {
//       type: DataTypes.STRING,
//       allowNull: false,
//     },
//     availability: {
//       type: DataTypes.STRING,
//     },

//     createdAt: DataTypes.DATE,
//     updatedAt: DataTypes.DATE,
//   },
//   {
//     tableName: "doctors",
//     sequelize
//   }
// )
