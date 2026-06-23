import express, { NextFunction, Request, Response } from 'express';
import HttpException from './src/common/http-exception.model';
import routes from './src/routes';

const app = express();

app.use(express.json());
app.use(routes);

app.use((err: Error | HttpException, req: Request, res: Response, next: NextFunction) => {
  if (err instanceof HttpException) {
    res.status(err.errorCode).json({message: err.message});
  } else if (err) {
    res.status(500).json({message: err.message});
  }
});

app.listen(8006, '0.0.0.0', () => {
    console.log('listening on 8006')
});