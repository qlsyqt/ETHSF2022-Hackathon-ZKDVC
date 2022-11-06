import { Injectable } from '@nestjs/common';
import { ConfigService } from '../config/config.service';
import { HttpService } from '../http/http.service';

@Injectable()
export class AuthService {
  constructor(
    private readonly http: HttpService, // 加载http
    private readonly configService: ConfigService,
  ) {}

  /**
   * requestAuth
   * 
   * @param offerId 
   * @returns 
   */
  async requestAuth(offerId: string) {
    const { code, result } = (
      await this.http.get(
        `${this.configService.get(
          'ISSUE_HOST',
        )}/api/v1/identity/auth?offerId=${offerId}`,
      )
    ).data; // 发送http请求

    if (code === 200) {
      return result;
    }
  }
}
