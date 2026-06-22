import { NextFunction, Request, Response, Router } from "express";

const router = Router();

router.get('/posts', (req: Request, res: Response, next: NextFunction) => {
  const result = { "message": "Hello world!" }
  res.json(result);
});

router.get('/posts/:id', (req: Request, res: Response, next: NextFunction) => {
  const result = { "message": "Hello everynyan!" }
  res.json(result);
});

router.post('/posts', (req: Request, res: Response, next: NextFunction) => {
  const result = { "message": "How are you?" }
  res.json(result);
});

router.patch('/posts/:id', (req: Request, res: Response, next: NextFunction) => {
  const result = { "message": "Fine, thank you!" }
  res.json(result);
});

router.delete('/posts/:id', (req: Request, res: Response, next: NextFunction) => {
  const result = { "message": "Oh my gah!" }
  res.json(result);
});

export default router;