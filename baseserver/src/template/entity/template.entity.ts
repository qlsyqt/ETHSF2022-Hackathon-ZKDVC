import { Offer } from '../../offer/entity/offer.entity';
import {
  Entity,
  PrimaryGeneratedColumn,
  Column,
  CreateDateColumn,
  Index,
  UpdateDateColumn,
  OneToMany,
} from 'typeorm';

@Entity()
export class Template {
  @PrimaryGeneratedColumn() id: number;

  @Column({ nullable: true })
  name: string;

  @Column({ nullable: true })
  dataCategory: string;

  @Column({ nullable: true })
  subCategory: string;

  // @Column({ nullable: true })
  // description: string;

  @Column({ nullable: true })
  isExpirable: boolean;

  @Column({ nullable: true })
  isAutoRevokable: boolean;

  // @Column({ nullable: true })
  // creatorUserId: string;

  @Column({ nullable: true })
  classfications: string;

  @OneToMany(() => Offer, (offer) => offer.template)
  offer: Offer[];

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

/*


  {
    
    lower:0,include:1,
    upper:1,include:1

    lower-number:11

    upper-number:233

    time 

    trigger : 0 1
}
string

[{},{}]
0 1  2 


json.string


6 
[{}} ,{}}]



offer 
  id
  clams [{expired：time}]


  */
