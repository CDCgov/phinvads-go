from collections.abc import AsyncIterator
from contextlib import asynccontextmanager
import time

from fastapi import FastAPI
from starlette.requests import Request
from starlette.responses import Response

from fastapi_cache import FastAPICache
from fastapi_cache.backends.redis import RedisBackend
from fastapi_cache.decorator import cache

from redis import asyncio as aioredis


@asynccontextmanager
async def lifespan(_: FastAPI) -> AsyncIterator[None]:
    redis = aioredis.from_url("redis://redis:6379")
    FastAPICache.init(RedisBackend(redis), prefix="fastapi-cache")
    yield


app = FastAPI(lifespan=lifespan)


@cache()
async def get_cache():
    return 1


@app.get("/")
async def index():
    return dict(status="OK")


@app.get("/blocking")
@cache(namespace="test", expire=10)
def blocking():
    time.sleep(2)
    return {"ret": 42}

# Get https://phinvads.cdc.gov/baseStu3/CodeSystem/ and return a whole code system result
@app.get("/phinvads/CodeSystem")
@cache(namespace="phinvads", expire=3600)
def get_code_systems(code: str = None):
    url = "https://phinvads.cdc.gov/baseStu3/CodeSystem"
    params = {
        "code": code,
    }
    response = requests.get(url, params=params)
    res_type = response.headers.get("content-type")
    if res_type == "application/json":
        return response.json()
    else: 
        return response.content
