import { MigrationInterface, QueryRunner, Table } from 'typeorm';

export class CreateBankAccountsTable1649489091075
  implements MigrationInterface
{
  public async up(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.createTable(
      new Table({
        name: 'bank_accounts',
        columns: [
          {
            name: 'id',
            type: 'uuid',
            isPrimary: true,
          },
          {
            name: 'account_number',
            type: 'varchar(255)', //Default is 255
          },
          {
            name: 'owner_name',
            type: 'varchar(255)', //Default is 255
          },
          {
            name: 'balance',
            type: 'double precision',
          },
          {
            name: 'created_at',
            type: 'timestamp',
            default: 'CURRENT_TIMESTAMP',
          },
        ],
      }),
    );
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.dropTable('bank_accounts');
  }
}
