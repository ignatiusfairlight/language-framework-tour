from datetime import datetime
from sqlmodel import SQLModel

class PostRead(SQLModel):
    id: int
    title: str
    content: str
    created_at: datetime
    updated_at: datetime