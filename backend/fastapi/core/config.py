from pydantic_settings import BaseSettings

class Settings(BaseSettings):
    api_prefix: str = "/api"
    db_host: str = "localhost"
    db_port: int = 5432
    db_name: str = "blog"
    db_user: str = "postgres"
    db_password: str = "secret"

settings = Settings()