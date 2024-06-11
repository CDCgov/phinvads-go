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
import requests


@asynccontextmanager
async def lifespan(_: FastAPI) -> AsyncIterator[None]:
    redis = aioredis.from_url("redis://localhost:6379")
    FastAPICache.init(RedisBackend(redis), prefix="fastapi-cache")
    yield


app = FastAPI(lifespan=lifespan)


@app.get("/")
async def index():
    return dict(status="OK")


# Get https://phinvads.cdc.gov/baseStu3/ValueSet and return response
@app.get("/phinvads/ValueSet")
@cache(namespace="phinvads", expire=3600)
def value_set(
    name: str = None,
    title: str = None,
    identifier: str = None,
    code: str = None,
    version: str = None,
    _getpages: str = None,
):
    url = "https://phinvads.cdc.gov/baseStu3/ValueSet"
    params = {
        "name": name,
        "title": title,
        "identifier": identifier,
        "code": code,
        "version": version,
        "_getpages": _getpages,
    }
    return get(url, params)


def get(url, params):
    try:
        response = requests.get(url, params=params)
        res_type = response.headers.get("content-type")
        if res_type == "application/json":
            return response.json()
        else:
            return response.content
    except Exception as e:
        return {"error": str(e)}
