from fastapi import FastAPI
from fastapi.exceptions import RequestValidationError
from starlette.exceptions import HTTPException
from starlette.middleware.cors import CORSMiddleware

app = FastAPI()

@app.get("/posts")
async def read_root():
    return {"message": "Hello world!"}

@app.get("/posts/{id}")
async def read_item(id: int):
    return {"id": id, "message": "Are you there?"}