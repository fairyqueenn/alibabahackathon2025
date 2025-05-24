from fastapi import FastAPI


from loguru import logger
from router.middleware import setup_middleware


import uvicorn

app = FastAPI()

# Initialize DB connection on startup
setup_middleware(app)
conn = None

@app.on_event("startup")
def startup():
    logger.info("initialized database connection")
    global conn
    conn = get_connection()
    if conn is None:
        logger.error("Failed to establish database connection")
        raise RuntimeError("Failed to establish database connection")


@app.on_event("shutdown")
def shutdown():
    global conn
    if conn:
        logger.info("closing database connection")
        conn.close()


app.include_router(router)

if __name__ == "__main__":
    import uvicorn
    uvicorn.run("main:app", host="0.0.0.0", port=8000, reload=True)