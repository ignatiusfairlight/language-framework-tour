from sqlalchemy import Column, DateTime, func
from sqlmodel import SQLModel, Field
from datetime import datetime

class Post(SQLModel, table=True):
    __tablename__ = "posts"
    id: int | None = Field(default=None, primary_key=True)
    title: str
    content: str
    created_at: datetime | None = Field(
        default=None,
        sa_column=Column(DateTime, server_default=func.now())
    )
    updated_at: datetime | None = Field(
        default=None,
        sa_column=Column(DateTime, server_default=func.now(), onupdate=func.now())
    )