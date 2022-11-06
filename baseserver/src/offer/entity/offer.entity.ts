import {
  Entity,
  PrimaryGeneratedColumn,
  Column,
  CreateDateColumn,
  Index,
  UpdateDateColumn,
  JoinColumn,
  ManyToOne,
} from 'typeorm';

import { Template } from '../../template/entity/template.entity';

@Entity()
export class Offer {
  @PrimaryGeneratedColumn()
  id: number;

  @Column({ nullable: true })
  name: string;

  @ManyToOne(() => Template, (template) => template.offer)
  @JoinColumn()
  template: Template;

  @Column({ nullable: true })
  link: string;

  @Column({ nullable: true })
  preClaims: string;

  @Column({ nullable: true })
  issuerDid: string;

  @CreateDateColumn({
    name: 'createdAt',
    comment: 'createdAt',
    nullable: true,
    type: 'timestamp',
    default: () => 'CURRENT_TIMESTAMP(6)',
  })
  @Index()
  createdAt: Date;

  @UpdateDateColumn({
    name: 'updatedAt',
    comment: 'updatedAt',
    nullable: true,
    type: 'timestamp',
    default: () => 'CURRENT_TIMESTAMP(6)',
    onUpdate: 'CURRENT_TIMESTAMP(6)',
  })
  @Index()
  updatedAt: Date;
}
