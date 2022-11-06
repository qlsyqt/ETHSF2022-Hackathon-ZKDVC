import { Test, TestingModule } from '@nestjs/testing';
import { INestApplication } from '@nestjs/common';
import * as request from 'supertest';
import { AppModule } from './../src/app.module';

describe('OfferController (e2e)', () => {
  let app: INestApplication;

  beforeEach(async () => {
    const moduleFixture: TestingModule = await Test.createTestingModule({
      imports: [AppModule],
    }).compile();

    app = moduleFixture.createNestApplication();
    await app.init();
  });

  afterAll(() => {});

  it('post /offer', async () => {
    const response = await request(app.getHttpServer())
      .post('/offer')
      .send({
        name: '2',
        template: 1,
        link: '44',
        preClaims: JSON.stringify([
          {
            name: '33',
            datacategory: '2',
            subcategory: '2',
            lowerBound: [0, 1, 2, 1], // 是否选择，是否包含 , 具体数值 ,  是否triger
            upperBound: [0, 1, 2, 1],
            createDate: new Date(),
            expirationDate: new Date(), // 过期时间
          },
        ]),
      });

    console.log('response', response);
    expect(response.status).toEqual(201);
    expect(response.body.result).toBeTruthy();
  });

  it('get id /offer/:id', async () => {
    const response = await request(app.getHttpServer()).get('/offer/' + 1);

    console.log('response', response.body);
    expect(response.status).toEqual(200);
    expect(response.body.result).toBeTruthy();

    expect(response.body.result.id).toEqual(1);
  });

  it('list /offer', async () => {
    const response = await request(app.getHttpServer()).get('/offer');

    console.log('response', response.body);
    expect(response.status).toEqual(200);
    expect(response.body.result.data).toBeTruthy();
  });
});
