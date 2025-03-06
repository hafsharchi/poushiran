import "reflect-metadata";
import express from "express";
import cors from "cors";
import { AppDataSource } from "./data-source";
import bookRoutes from "./routes/book.routes";
import swaggerUi from "swagger-ui-express";

// Swagger documentation
const swaggerDocument = {
    openapi: "3.0.0",
    info: {
        title: "Book API",
        version: "1.0.0",
        description: "A simple Book API"
    },
    servers: [
        {
            url: "http://localhost:3000",
            description: "Development server"
        }
    ],
    paths: {
        "/api/books": {
            get: {
                summary: "Returns all books",
                responses: {
                    "200": {
                        description: "A list of books",
                        content: {
                            "application/json": {
                                schema: {
                                    type: "array",
                                    items: {
                                        $ref: "#/components/schemas/Book"
                                    }
                                }
                            }
                        }
                    }
                }
            },
            post: {
                summary: "Create a new book",
                requestBody: {
                    required: true,
                    content: {
                        "application/json": {
                            schema: {
                                $ref: "#/components/schemas/BookInput"
                            }
                        }
                    }
                },
                responses: {
                    "200": {
                        description: "Book created successfully"
                    }
                }
            }
        }
    },
    components: {
        schemas: {
            Book: {
                type: "object",
                properties: {
                    id: { type: "integer" },
                    title: { type: "string" },
                    author: { type: "string" },
                    description: { type: "string" },
                    isbn: { type: "string" },
                    price: { type: "number" },
                    createdAt: { type: "string", format: "date-time" },
                    updatedAt: { type: "string", format: "date-time" }
                }
            },
            BookInput: {
                type: "object",
                required: ["title", "author", "isbn"],
                properties: {
                    title: { type: "string" },
                    author: { type: "string" },
                    description: { type: "string" },
                    isbn: { type: "string" },
                    price: { type: "number" }
                }
            }
        }
    }
};

// Initialize express app
const app = express();

// Middleware
app.use(cors());
app.use(express.json());

// Routes
app.use("/api/books", bookRoutes);

// Swagger
app.use("/api-docs", swaggerUi.serve, swaggerUi.setup(swaggerDocument));

// Initialize TypeORM
AppDataSource.initialize()
    .then(() => {
        console.log("Data Source has been initialized!");
        
        // Start express server
        app.listen(3000, () => {
            console.log("Server is running on port 3000");
            console.log("Swagger documentation is available at http://localhost:3000/api-docs");
        });
    })
    .catch((error) => console.log("Error during Data Source initialization:", error));