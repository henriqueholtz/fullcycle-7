import {
  BeforeInsert,
  Column,
  CreateDateColumn,
  Entity,
  PrimaryGeneratedColumn,
} from 'typeorm';
import { V4 as uuidV4 } from 'uuid';

@Entity({
  name: 'bank_accounts',
})
export class BankAccount {
  @PrimaryGeneratedColumn('uuid')
  id: string;

  @Column()
  account_number: string;

  @Column()
  owner_name: string;

  @Column()
  balance: number;

  @CreateDateColumn()
  created_at: Date;

  @BeforeInsert()
  generateId() {
    if (!this.id) {
      this.id = uuidV4();
    }
  }

  @BeforeInsert()
  initBalance() {
    if (!this.balance) this.balance = 0;
  }
}
