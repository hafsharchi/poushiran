import { Router } from "express";
import { BookController } from "../controller/BookController";

const router = Router();
const bookController = new BookController();

// Get all books
router.get("/", bookController.all.bind(bookController));

// Get one book
router.get("/:id", bookController.one.bind(bookController));

// Create a new book
router.post("/", bookController.create.bind(bookController));

// Update a book
router.put("/:id", bookController.update.bind(bookController));

// Delete a book
router.delete("/:id", bookController.remove.bind(bookController));

export default router;