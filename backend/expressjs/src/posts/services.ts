import db from '../common/database';
import HttpException from '../common/http-exception.model';
import { CreatePost, EditPost } from './DTOs';

const getAll = async () => {
  return await db.any('SELECT * FROM posts');
};

const getById = async (id: number) => {
  try {
    return await db.one('SELECT * FROM posts WHERE id = $1', [id]);
  } catch (error) {
    throw new HttpException(404, "Post not found")
  }
};

const createPost = async (data: CreatePost) => {
  return await db.one('INSERT INTO posts (title, content) VALUES ($1, $2) RETURNING *', [data.title, data.content]);
};

const editPost = async (id: number, data: EditPost) => {
  const fields = [];
  const values = [];
  let i = 1;

  if (data.title !== undefined) { fields.push(`title = $${i++}`); values.push(data.title); }
  if (data.content !== undefined) { fields.push(`content = $${i++}`); values.push(data.content); }

  values.push(id)

  try {
    return await db.one(
      `UPDATE posts SET ${fields.join(', ')}, updated_at = NOW() WHERE id = $${i} RETURNING *`,
      values
    ); 
  } catch (error) {
    throw new HttpException(404, "Post not found");
  }
};

const deletePost = async (id: number) => {
  const result = await db.result('DELETE FROM posts WHERE id = $1', [id]);
  if (result.rowCount === 0) {
    throw new HttpException(404, "Post not found");
  }
};

export default { getAll, getById, createPost, editPost, deletePost };