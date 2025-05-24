import os
import dashscope

from fastapi import APIRouter, HTTPException, UploadFile, File
import base64
from uuid import uuid4
from pydantic import BaseModel
from api.tools import prompt_templates
from loguru import logger
from dotenv import load_dotenv
load_dotenv()


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

class ReviewRequest(BaseModel):
    image_id: str

@router.post("/image_reviewer")
def image_reviewer(request: ReviewRequest):
    try:
        image_data_url = image_store.get(request.image_id)
        print(f"Image data URL: {image_data_url}")
        if not image_data_url:
            raise HTTPException(status_code=404, detail="Image not found.")
        
        dashscope.base_http_api_url = 'https://dashscope-intl.aliyuncs.com/api/v1'
        messages = [
            {
                "role": "user",
                "content": [
                    {"image": image_data_url},
                    {"text": prompt_templates.prompts.image_reviewer}
                ]
            }
        ]
        response = dashscope.MultiModalConversation.call(
            api_key=os.getenv('MODEL_API_KEY'),
            model='qwen-vl-plus-latest',
            messages=messages
        )
        logger.info(f"Response: {response}")
        
        # Convert the response to a serializable format
        try:
            # The dashscope response object acts like a dict, so we can convert it directly
            serializable_response = dict(response)
        except Exception as e:
            logger.warning(f"Could not convert response to dict: {e}")
            # Fallback: manually extract the data structure
            serializable_response = {
                "status_code": getattr(response, 'status_code', None),
                "request_id": getattr(response, 'request_id', None),
                "code": getattr(response, 'code', None),
                "message": getattr(response, 'message', None),
                "output": dict(response.output) if hasattr(response, 'output') else None,
                "usage": dict(response.usage) if hasattr(response, 'usage') else None
            }
        
        # Extract just the text result if you only need that
        result_text = None
        try:
            if hasattr(response, 'output') and hasattr(response.output, 'choices'):
                result_text = response.output.choices[0].message.content[0]['text']
        except (AttributeError, IndexError, KeyError, TypeError) as e:
            logger.warning(f"Could not extract text from response: {e}")

        return {
            "result": serializable_response,
            "text": result_text  # Include extracted text for convenience
        }
        
    except Exception as e:
        logger.error(f"Unexpected error: {e}")
        raise HTTPException(status_code=500, detail="Internal Server Error")
