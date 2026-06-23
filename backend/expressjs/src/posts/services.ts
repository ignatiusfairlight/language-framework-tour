import db from '../common/database';
import { CreatePost, EditPost } from './DTOs';

const getAll = async () => {
  return await db.any('SELECT * FROM posts');
};

const getById = async (id: number) => {
  return await db.one('SELECT * FROM posts WHERE id = $1', [id]);
};

const createPost = async (data: CreatePost) => {
  return await db.one('INSERT INTO posts (title, content) VALUES ($1, $2) RETURNING *', [data.title, data.content]);
};

const editPost = async (id: number, data: EditPost) => {
  return await db.one('UPDATE posts SET title = $1, content = $2, updated = NOW() WHERE id = $3 RETURNING *', [data.title, data.content, id]);
};

const deletePost = async (id: number) => {
  return await db.none('DELETE FROM posts WHERE id = $1', [id]);
};

export default { getAll, getById, createPost, editPost, deletePost };