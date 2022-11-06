import { ApiProperty } from '@nestjs/swagger';
import { IsString, IsOptional, IsBoolean } from 'class-validator';

export class CreateTemplateDto {
  @ApiProperty({ type: 'string' })
  @IsString()
  name: string;

  @ApiProperty({ type: 'string', required: false })
  @IsString()
  @IsOptional()
  dataCategory: string;

  @ApiProperty({ type: 'string', required: false })
  @IsString()
  @IsOptional()
  subCategory: string;

  @ApiProperty({ type: 'string', required: false })
  @IsString()
  @IsOptional()
  description: string;

  @ApiProperty({ type: 'boolean', required: false })
  @IsBoolean()
  @IsOptional()
  isExpirable: boolean;

  @ApiProperty({ type: 'boolean', required: false })
  @IsBoolean()
  @IsOptional()
  isAutoRevokable: boolean;

  @ApiProperty({
    type: 'string',
    description: 'Classfication Json string',
    required: false,
  })
  @IsString()
  @IsOptional()
  classfications: string;
}
