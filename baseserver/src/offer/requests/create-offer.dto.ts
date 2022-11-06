import { ApiProperty } from '@nestjs/swagger';
import { IsString, IsOptional, IsNumber } from 'class-validator';

export class CreateOfferDto {
  @ApiProperty({ type: 'string' })
  @IsString()
  name: string;

  @ApiProperty({ type: 'number', required: false })
  @IsNumber()
  @IsOptional()
  template: number;

  @ApiProperty({ type: 'string', required: false })
  @IsString()
  @IsOptional()
  issuerId: string;

  @ApiProperty({ type: 'string', required: false })
  @IsString()
  @IsOptional()
  link: string;

  @ApiProperty({
    type: 'string',
    description: 'preClaims Json string',
    required: false,
  })
  @IsString()
  @IsOptional()
  preClaims: [];
}
