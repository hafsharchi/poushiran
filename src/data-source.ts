import "reflect-metadata";
import { DataSource } from "typeorm";
import { Book } from "./entity/Book";

export const AppDataSource = new DataSource({
    type: "sqlite",
    database: "database.sqlite",
    synchronize: false,
    logging: true,
    entities: [Book],
    migrations: ["src/migration/**/*.ts"],
    subscribers: [],
});