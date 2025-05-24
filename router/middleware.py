from fastapi import Request
from fastapi.middleware.cors import CORSMiddleware
from loguru import logger

def setup_middleware(app):
    # Middleware untuk CORS
    app.add_middleware(
        CORSMiddleware,
        allow_origins=["*"],  # Ganti dengan domain spesifik jika perlu
        allow_credentials=True,
        allow_methods=["*"],
        allow_headers=["*"],
    )
    
    @app.middleware("http")
    async def log_requests(request:Request, call_next):
        logger.info(f"incoming request: {request.method}{request.url}")
        response = await call_next(request)
        logger.info(f"Response status: {response.status_code} ")
        
        return response