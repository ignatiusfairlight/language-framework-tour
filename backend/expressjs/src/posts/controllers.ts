import { NextFunction, Request, Response, Router } from "express";
import postService from './services';
import { CreatePost, EditPost } from "./DTOs";

const router = Router();

router.get('/posts', async (req: Request, res: Response, next: NextFunction) => {
  const result = await postService.getAll();
  res.json(result);
});

router.get('/posts/:id', async (req: Request, res: Response, next: NextFunction) => {
  try {
    const result = await postService.getById(Number(req.params.id));
    res.json(result);
  } catch (error) {
    next(error)
  }
});

router.post('/posts', (req: Request, res: Response, next: NextFunction) => {
  const result = postService.createPost(req.body() as CreatePost);
  res.json(result);
});

router.patch('/posts/:id', (req: Request, res: Response, next: NextFunction) => {
  try {
    const result = postService.editPost(
      Number(req.params.id),
      req.body() as EditPost
    );
    res.json(result);
  } catch (error) {
    next(error);
  };
});

router.delete('/posts/:id', (req: Request, res: Response, next: NextFunction) => {
  try {
    const result = postService.deletePost(Number(req.params.id));
    res.json(result);
  } catch (error) {
    next(error);
  };
});

export default router;