from datetime import datetime
from sqlmodel import SQLModel

class PostRead(SQLModel):
    id: int
    title: str
    content: str
    created_at: datetime
    updated_at: datetime

class PostCreate(SQLModel):
    title: str
    content: str

class PostEdit(SQLModel):
    title: str | None = None
    content: str | None = None