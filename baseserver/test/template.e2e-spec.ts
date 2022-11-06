import { Test, TestingModule } from '@nestjs/testing';
import { INestApplication } from '@nestjs/common';
import * as request from 'supertest';
import { AppModule } from './../src/app.module';

describe('TemplateController (e2e)', () => {
  let app: INestApplication;

  beforeEach(async () => {
    const moduleFixture: TestingModule = await Test.createTestingModule({
      imports: [AppModule],
    }).compile();

    app = moduleFixture.createNestApplication();
    await app.init();
  });

  afterAll(() => {});

  it('post /template', async () => {
    const response = await request(app.getHttpServer())
      .post('/template')
      .send({
        name: '1',
        dataCategory: '2',
        subCategory: '3',
        isExpirable: true,
        isAutoRevokable: true,
        classfications: JSON.stringify([
          {
            name: '33',
            description: '22',
            lowerBound: [0, 1], // 是否选择，是否包含 
            upperBound: [0, 1],
          },
        ]),
      });

    console.log('response', response);
    expect(response.status).toEqual(201);
    expect(response.body.result).toBeTruthy();
  });

  it('get id /template/:id', async () => {
    const response = await request(app.getHttpServer()).get('/template/' + 1);

    console.log('response', response.body);
    expect(response.status).toEqual(200);
    expect(response.body.result).toBeTruthy();

    expect(response.body.result.id).toEqual(1);
  });

  it('list /template', async () => {
    const response = await request(app.getHttpServer()).get('/template');

    console.log('response', response.body);
    expect(response.status).toEqual(200);
    expect(response.body.result.data).toBeTruthy();
  });
});
