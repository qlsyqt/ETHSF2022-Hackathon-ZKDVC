import { Controller } from '@nestjs/common';
import {
  Crud,
  CrudController,
  CrudRequest,
  Override,
  ParsedBody,
  ParsedRequest,
} from '@nestjsx/crud';
import { Offer } from './entity/offer.entity';
import { OfferService } from './offer.service';
import { dto } from './requests';

@Controller('offer')
@Crud({
  model: {
    type: Offer,
  },
  dto,
  query: {
    alwaysPaginate: true,
    join: {
      template: {
        eager: true,
      },
    },
  },
})
export class OfferController implements CrudController<Offer> {
  constructor(public service: OfferService) {}

  get base(): CrudController<Offer> {
    return this;
  }

  @Override()
  async createOne(@ParsedRequest() req: CrudRequest, @ParsedBody() dto: Offer) {
    dto.issuerDid = await this.service.getDid();
    return this.base.createOneBase(req, dto);
  }
}
