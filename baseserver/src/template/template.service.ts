import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { TypeOrmCrudService } from '@nestjsx/crud-typeorm';
import { Template } from './entity/template.entity';

@Injectable()
export class TemplateService extends TypeOrmCrudService<Template> {
  constructor(@InjectRepository(Template) repo) {
    super(repo);
  }
}
