import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { TypeOrmCrudService } from '@nestjsx/crud-typeorm';
import { Offer } from './entity/offer.entity';

@Injectable()
export class OfferService extends TypeOrmCrudService<Offer> {
  constructor(@InjectRepository(Offer) repo) {
    super(repo);
  }

  async getDid() {
    const result = await this.repo.query(
      "select * from t_did where username = 'account@knn3'",
    );

    if (result.length) {
      return result[0]?.did;
    }
  }
}
