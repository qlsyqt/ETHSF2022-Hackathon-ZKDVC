import { Controller } from '@nestjs/common';
import { Crud } from '@nestjsx/crud';
import { Template } from './entity/template.entity';
import { TemplateService } from './template.service';
import { dto } from './requests';

@Controller('template')
@Crud({
  model: {
    type: Template,
  },
  dto,
  query: {
    alwaysPaginate: true,
  },
})
export class TemplateController {
  constructor(public service: TemplateService) {}
}
