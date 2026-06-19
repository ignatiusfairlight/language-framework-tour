from fastapi import APIRouter

router = APIRouter()

@router.get("/")
async def show_all():
    return {"message": "Hello world!"}

@router.get("/{id}")
async def show_one(id: int):
    return {"message": "Hello everynyan!"}

@router.post("/")
async def create():
    return {"message": "How are you?"}

@router.patch("/{id}")
async def update(id: int):
    return {"message": "Fine, thank you!"}

@router.delete("/{id}")
async def delete(id: int):
    return {"message": "Oh my gah"}