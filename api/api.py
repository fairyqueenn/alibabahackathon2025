import os
import dashscope
import numpy as np
from langchain_community.chat_models import ChatTongyi
from langchain_qwq import ChatQwQ


from fastapi import APIRouter, HTTPException, UploadFile, File
import base64
from uuid import uuid4
from pydantic import BaseModel
from api.tools import prompt_templates
from langchain_community.embeddings.dashscope import DashScopeEmbeddings
from loguru import logger
from langchain_text_splitters import CharacterTextSplitter

from langchain_community.document_loaders import TextLoader
from langchain_postgres import PGVector

from langchain.prompts import ChatPromptTemplate
from langchain_core.prompts import PromptTemplate
import faiss
from langchain_community.docstore.in_memory import InMemoryDocstore
from langchain_community.vectorstores import FAISS
from langchain_core.documents import Document
from langchain_text_splitters import RecursiveCharacterTextSplitter


from dotenv import load_dotenv
load_dotenv()

import warnings
warnings.filterwarnings("ignore", category=RuntimeWarning, module="numpy")

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
    image_url: str
    description: str

@router.post("/image_reviewer")
def image_reviewer(request: ReviewRequest):
    try:

        dashscope.base_http_api_url = 'https://dashscope-intl.aliyuncs.com/api/v1'
        messages = [
            {
                "role": "user",
                "content": [
                    {"image": request.image_url},
                    {"text": prompt_templates.prompts.image_reviewer.format(description=request.description)}
                ]
            }
        ]
        response = dashscope.MultiModalConversation.call(
            api_key=os.getenv('MODEL_API_KEY'),
            model='qwen-vl-plus-latest',
            messages=messages,
            temperature=0.0,
            max_tokens=1500,
            
        
        )
        logger.info(f"Response: {response}")

        # Convert the response to a serializable format
        try:
            # The dashscope response object acts like a dict, so we can convert it directly
            serializable_response = {} # TODO: adjust for keperluan
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
    
class ReviewRequest(BaseModel):
    query: str
    
@router.post("/embedding_query")
def embedding_query(request: queryRequest):
    try:
        # Validate API key
        api_key = os.getenv("MODEL_API_KEY")
        if not api_key:
            raise ValueError("MODEL_API_KEY environment variable is not set")
        
        print(f"API Key loaded: {api_key[:10]}..." if len(api_key) > 10 else "API Key too short")
        embeddings = DashScopeEmbeddings(
            model="text-embedding-v3", 
            dashscope_api_key=api_key
        )
        # Initialize embeddings and vector store
        embeddings_dim = len(embeddings.embed_query('hello world'))
        index = faiss.IndexFlatL2(embeddings_dim)
        vector_store = FAISS(
            embedding_function=embeddings,
            index=index,
            docstore=InMemoryDocstore(),
            index_to_docstore_id={}
        )
        
        # Initialize model
        try:
            model = ChatQwQ(
                model="qwen-plus-latest",
                api_key=api_key,
                temperature=0.0,  
            )
            print("Model initialized successfully")
        except Exception as e:
            print(f"Error initializing model: {e}")
            raise
        
        # Load and process documents
        try:
            loader = TextLoader(r"data\indonesian_food_health_categories_cleaned.csv")
            docs = loader.load()
            
            print("Splitting document into chunks")
            text_splitter = RecursiveCharacterTextSplitter(chunk_size=1000, chunk_overlap=200)
            all_splits = text_splitter.split_documents(docs)
            
            if not all_splits:
                raise ValueError("No document chunks created")
            
            print("Indexing chunks")
            vector_store.add_documents(documents=all_splits)
                
        except Exception as e:
            print(f"Error processing documents: {e}")
            raise
        
        # Retrieve relevant documents
        try:
            retriever = vector_store.as_retriever()
            retrieved_docs = retriever.get_relevant_documents(request.query)
            
            if not retrieved_docs:
                return {
                    "response": "No relevant documents found for your query.",
                    "documents_found": 0,
                    "status": "success"
                }
            
            # Combine document content
            print("Combining all chunks")
            docs_content = "\n\n".join(doc.page_content for doc in retrieved_docs)
            print("Combined docs:", docs_content)
            
            # Generate response using LLM
            print("Ingesting to LLM")
            prompt = ChatPromptTemplate.from_template(prompt_templates.prompts.food_recommender)

            messages = prompt.invoke({"question": ReviewRequest.query, "context": docs_content})
            print("===Query===")
            print(messages)
            
            response = model.invoke(messages)
            print("Response:", response)
            
            return {
                "response": response.content if hasattr(response, 'content') else str(response),
                "documents_found": len(retrieved_docs),
                "status": "success"
            }
            
        except Exception as e:
            print(f"Error generating response: {e}")
            raise
            
    except Exception as e:
        print(f"Critical error in embedding_query: {e}")
        return {
            "error": str(e),
            "status": "error"
        }