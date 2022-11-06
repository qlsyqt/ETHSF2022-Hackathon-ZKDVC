import { Test, TestingModule } from '@nestjs/testing';
import { INestApplication } from '@nestjs/common';
import * as request from 'supertest';
import { AppModule } from './../src/app.module';

describe('AuthController (e2e)', () => {
  let app: INestApplication;

  beforeEach(async () => {
    const moduleFixture: TestingModule = await Test.createTestingModule({
      imports: [AppModule],
    }).compile();

    app = moduleFixture.createNestApplication();
    await app.init();
  });

  afterAll(() => {});

  it('post /auth/create', async () => {
    const response = await request(app.getHttpServer())
      .post('/auth/create')
      .send({
        offerId: '1',
      });

    console.log('response', response);
    expect(response.status).toEqual(201);
    expect(response.body.result).toBeTruthy();
  });
});
