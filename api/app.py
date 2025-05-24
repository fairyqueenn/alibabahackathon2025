import os
import dashscope

from fastapi import APIRouter, HTTPException, UploadFile, File
import base64

from pydantic import BaseModel
from loguru import logger

router = APIRouter()

class queryRequest(BaseModel):
    query:str

image_store = {}

@router.post("/upload_image")
async def upload_image(file: UploadFile = File(...)):
    try:
        file_bytes = await file.read()
        image_base64 = base64.b64encode(file_bytes).decode('utf-8')
        image_format = file.content_type.split('/')[-1]  # e.g., 'png', 'jpeg'

        image_data_url = f"data:image/{image_format};base64,{image_base64}"
        
        image_id = str(uuid4())
        image_store[image_id] = image_data_url

        logger.info(f"Image uploaded successfully with ID: {image_id}")
        return {"image_id": image_id}

    except Exception as e:
        logger.error(f"Image upload failed: {e}")
        raise HTTPException(status_code=500, detail="Failed to upload image.")

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
    

