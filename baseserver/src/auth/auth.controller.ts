import { Body, Controller, Post } from '@nestjs/common';
import { AuthService } from './auth.service';
import { CreateAuthDto } from './requests';

@Controller('auth')
export class AuthController {
  constructor(public authService: AuthService) {}

  @Post('create')
  async create(@Body() dto: CreateAuthDto) {
    return this.authService.requestAuth(dto.offerId);
    /*
    return {
      id: '1',
      typ: 'application/iden3comm-plain-json',
      type: 'https://iden3-communication.io/authorization/1.0/request',
      thid: 'de5de3e1-a343-4ed3-a308-88320f65bb38',
      body: {
        callbackUrl: 'http://0.0.0.0:8080/api/v1/identity/callback',
        reason: 'auth login',
        message: '5577006791947779410',
        scope: [
          {
            id: 0,
            circuit_id: 'auth',
            rules: {
              challenge: '5577006791947779410',
            },
          },
          {
            id: 1,
            circuit_id: 'ecdsa',
            rules: {
              challenge: '8674665223082153551',
            },
          },
        ],
      },
      from: '1125GJqgw6YEsKFwj63GY87MMxPL9kwDKxPUiwMLNZ',
    };

    */
  }
}
