import pgPromise from 'pg-promise';

const pgp = pgPromise();
const db = pgp('postgres://postgres:secret@db:5432/blog');

export default db;