import { Entity, PrimaryGeneratedColumn, Column, CreateDateColumn, UpdateDateColumn } from "typeorm";

@Entity()
export class Product {
    @PrimaryGeneratedColumn()
    id: number;

    @Column()
    name: string;

    @Column({ type: "text", nullable: true })
    description: string;

    @Column("decimal", { precision: 10, scale: 2 })
    price: number;

    @Column()
    sku: string;

    @Column({ default: 0 })
    stockQuantity: number;

    @Column({ default: true })
    isActive: boolean;

    @Column({ nullable: true })
    category: string;

    @CreateDateColumn()
    createdAt: Date;

    @UpdateDateColumn()
    updatedAt: Date;
}