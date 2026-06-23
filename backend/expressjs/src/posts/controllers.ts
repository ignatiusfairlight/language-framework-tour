import { NextFunction, Request, Response, Router } from "express";
import postService from './services';
import { CreatePost, EditPost } from "./DTOs";

const router = Router();

router.get('/posts', async (req: Request, res: Response, next: NextFunction) => {
  const result = await postService.getAll();
  res.json(result);
});

router.get('/posts/:id', async (req: Request, res: Response, next: NextFunction) => {
  const result = await postService.getById(Number(req.params.id));
  res.json(result);
});

router.post('/posts', async (req: Request, res: Response, next: NextFunction) => {
  const result = await postService.createPost(req.body as CreatePost);
  res.json(result);
});

router.patch('/posts/:id', async (req: Request, res: Response, next: NextFunction) => {
  const result = await postService.editPost(
    Number(req.params.id),
    req.body as EditPost
  );
  res.json(result);
});

router.delete('/posts/:id', async (req: Request, res: Response, next: NextFunction) => {
  await postService.deletePost(Number(req.params.id));
  res.json({message: "Post deleted"});
});

export default router;