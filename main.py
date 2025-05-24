from fastapi import FastAPI
from api.api import router
from loguru import logger
from router.middleware import setup_middleware


import uvicorn

app = FastAPI()

# Initialize DB connection on startup
setup_middleware(app)


app.include_router(router)

if __name__ == "__main__":
    import uvicorn
    uvicorn.run("main:app", host="0.0.0.0", port=8000, reload=True)