import os
import dashscope

from fastapi import APIRouter, HTTPException

from pydantic import BaseModel
from loguru import logger

router = APIRouter()

class queryRequest(BaseModel):
    query:str
    

@router.post("/image_reviewer")
def image_reviewer(request: queryRequest):
    """
    Execute an SQL query via FastAPI endpoint.
    """
    try:
        dashscope.base_http_api_url = 'https://dashscope-intl.aliyuncs.com/api/v1'
        messages = [
            {
                "role": "user",
                "content": [
                    {"image": "https://dashscope.oss-cn-beijing.aliyuncs.com/images/rabbit.png"},
                    {"text": "What are these?"}
                ]
            }
        ]
        response = dashscope.MultiModalConversation.call(
            # If environment variable is not configured, replace the line below with: api_key="sk-xxx",
            api_key=os.getenv('MODEL_API_KEY'),
            # This example uses qwen-vl-max. You can change the model name as needed. Model list: https://www.alibabacloud.com/help/zh/model-studio/getting-started/models
            model='qwen-vl-plus-latest',
            messages=messages
            )
        logger.info(f"Response: {response}")
        return response
    except Exception as e:
        logger.error(f"Unexpected error: {e}")
        raise HTTPException(status_code=500, detail="Internal Server Error")
    

