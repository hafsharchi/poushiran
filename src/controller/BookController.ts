import { Request, Response } from "express";
import { AppDataSource } from "../data-source";
import { Book } from "../entity/Book";

export class BookController {
    private bookRepository = AppDataSource.getRepository(Book);

    async all(req: Request, res: Response) {
        try {
            const books = await this.bookRepository.find();
            return res.json(books);
        } catch (error) {
            return res.status(500).json({ error: "Internal server error" });
        }
    }

    async one(req: Request, res: Response) {
        try {
            const id = parseInt(req.params.id);
            const book = await this.bookRepository.findOne({ where: { id } });

            if (!book) {
                return res.status(404).json({ error: "Book not found" });
            }
            return res.json(book);
        } catch (error) {
            return res.status(500).json({ error: "Internal server error" });
        }
    }

    async create(req: Request, res: Response) {
        try {
            const book = this.bookRepository.create(req.body);
            const results = await this.bookRepository.save(book);
            return res.json(results);
        } catch (error) {
            return res.status(500).json({ error: "Internal server error" });
        }
    }

    async update(req: Request, res: Response) {
        try {
            const id = parseInt(req.params.id);
            const book = await this.bookRepository.findOne({ where: { id } });

            if (!book) {
                return res.status(404).json({ error: "Book not found" });
            }

            this.bookRepository.merge(book, req.body);
            const results = await this.bookRepository.save(book);
            return res.json(results);
        } catch (error) {
            return res.status(500).json({ error: "Internal server error" });
        }
    }

    async remove(req: Request, res: Response) {
        try {
            const id = parseInt(req.params.id);
            const book = await this.bookRepository.findOne({ where: { id } });

            if (!book) {
                return res.status(404).json({ error: "Book not found" });
            }

            await this.bookRepository.remove(book);
            return res.json({ message: "Book has been removed" });
        } catch (error) {
            return res.status(500).json({ error: "Internal server error" });
        }
    }
}