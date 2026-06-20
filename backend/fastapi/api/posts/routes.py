from fastapi import APIRouter, Depends, Response, status, HTTPException
from sqlmodel.ext.asyncio.session import AsyncSession
from core.database import get_session, SessionDep
from . import services
from .schemas import PostRead, PostCreate, PostEdit

router = APIRouter()

@router.get("/", response_model=list[PostRead])
async def show_all(session: SessionDep):
    return await services.get_all_posts(session)

@router.get("/{id}", response_model=PostRead)
async def show_one(session: SessionDep, id: int):
    return await services.get_by_id(session, id)

@router.post("/", response_model=PostRead)
async def create(session: SessionDep, data: PostCreate):
    return await services.create_post(session, data)

@router.patch(
    "/{id}",
    response_model=PostRead,
    responses={404: {"description":"Post not found"}}
)
async def update(session: SessionDep, id: int, data: PostEdit):
    post = await services.update_post(session, id, data)
    if post is None:
        raise HTTPException(status_code=404, detail="Post not found")
    return post

@router.delete(
    "/{id}",
    status_code=status.HTTP_204_NO_CONTENT,
    responses={404: {"description":"Post not found"}}
)
async def delete(session: SessionDep, id: int):
    deleted = await services.delete_post(session, id)
    if not deleted:
        raise HTTPException(status_code=404, detail="Post not found")
    return Response(status_code=status.HTTP_204_NO_CONTENT)