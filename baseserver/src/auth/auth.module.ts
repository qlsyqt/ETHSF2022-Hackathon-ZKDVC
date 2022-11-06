import { Module } from '@nestjs/common';
import { HttpModule } from '../http/http.module';
import { AuthController } from './auth.controller';
import { AuthService } from './auth.service';

@Module({
  imports: [
    HttpModule.register({ baseURL: '' }), // 正常http请求使用该模块
  ],
  controllers: [AuthController],
  providers: [AuthService],
})
export class AuthModule {}
