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
from langchain.schema import Document  # <-- Tambahkan ini

from langchain_community.document_loaders import TextLoader
from langchain_postgres import PGVector

from langchain.prompts import ChatPromptTemplate
from langchain_core.prompts import PromptTemplate


from dotenv import load_dotenv
load_dotenv()
api_key=os.getenv('MODEL_API_KEY')
dashscope.base_http_api_url = 'https://dashscope-intl.aliyuncs.com/api/v1'

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
    language: str

@router.post("/image_reviewer")
def image_reviewer(request: ReviewRequest):
    try:

        dashscope.base_http_api_url = 'https://dashscope-intl.aliyuncs.com/api/v1'
        messages = [
            {
                "role": "user",
                "content": [
                    {"image": request.image_url},
                    {"text": prompt_templates.prompts.image_reviewer.format(description=request.description, language=request.language)}
                ]
            }
        ]
        logger.info(f"Messages prepared for API call: {messages}")
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

class RecRquest(BaseModel):
    query: str
@router.post("/embedding_query")
def embedding_query(request: RecRquest):
    
    try:
        # Validate API key
        api_key = os.getenv("MODEL_API_KEY")
        if not api_key:
            raise ValueError("MODEL_API_KEY environment variable is not set")
        
        print(f"API Key loaded: {api_key[:10]}..." if len(api_key) > 10 else "API Key too short")
        
        # Initialize embeddings and model
        try:
            embeddings = DashScopeEmbeddings(
                model="text-embedding-v3", 
                dashscope_api_key=api_key
            )
            print("Embeddings initialized successfully")
        except Exception as e:
            print(f"Error initializing embeddings: {e}")
            raise
        
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
            loader = TextLoader("data/indonesian_food_health_categories_cleaned.csv")
            docs = loader.load()
            
            if not docs:
                raise ValueError("No documents loaded from file")
            
            text_splitter = CharacterTextSplitter(chunk_size=1000, chunk_overlap=0)
            all_splits = text_splitter.split_documents(docs)
            print(f"Successfully split documents into {len(all_splits)} chunks")
            
            if not all_splits:
                raise ValueError("No document chunks created")
                
        except Exception as e:
            print(f"Error processing documents: {e}")
            raise
        
        # Initialize PostgreSQL connection
        try:
            PG_HOST = os.getenv("PG_HOST")
            PG_PORT = os.getenv("PG_PORT") 
            PG_DATABASE = os.getenv("PG_DB")
            PG_USER = os.getenv("PG_USER")
            PG_PASSWORD = os.getenv("PG_PASSWORD")
            
            # Validate all required environment variables
            pg_vars = {
                "PG_HOST": PG_HOST,
                "PG_PORT": PG_PORT,
                "PG_DATABASE": PG_DATABASE,
                "PG_USER": PG_USER,
                "PG_PASSWORD": PG_PASSWORD
            }
            
            missing_vars = [k for k, v in pg_vars.items() if not v]
            if missing_vars:
                raise ValueError(f"Missing PostgreSQL environment variables: {missing_vars}")

            connection = f"postgresql+psycopg://{PG_USER}:{PG_PASSWORD}@{PG_HOST}:{PG_PORT}/{PG_DATABASE}"
            print("PostgreSQL connection string created")
            
        except Exception as e:
            print(f"Error setting up PostgreSQL connection: {e}")
            raise
        
        # Initialize vector store
        try:
            collection_name = "new_food_recommender"
            vector_store = PGVector(
                embeddings=embeddings,
                collection_name=collection_name,
                connection=connection,
                use_jsonb=True,
            )
            print("Vector store initialized successfully")
            
        except Exception as e:
            print(f"Error initializing vector store: {e}")
            raise
        
        # Add documents to vector store (with batch processing)
        try:
            print(f"Indexing {len(all_splits)} chunks...")
            
            # Process in smaller batches to avoid memory issues
            batch_size = 50
            for i in range(0, len(all_splits), batch_size):
                batch = all_splits[i:i+batch_size]
                print(batch)
                print(f"Processing batch {i//batch_size + 1}: documents {i+1}-{min(i+batch_size, len(all_splits))}")
                vector_store.add_documents(documents=batch)
            
            print("Successfully indexed all chunks into vector store")
            
        except Exception as e:
            print(f"Error indexing documents: {e}")
            raise
        
        # Process query
        try:
            query_text = request.query

            if not query_text or not query_text.strip():
                raise ValueError("Query text is empty")
                
            print(f"Processing query: {query_text}")
            
        except Exception as e:
            print(f"Error processing query: {e}")
            raise
        
        # Retrieve similar documents
        try:
            print("Finding relevant documents...")
            # Reduce k and fetch_k to avoid potential issues
            retrieved_docs = vector_store.similarity_search(query_text, k=5)
            print(f"Found {len(retrieved_docs)} relevant documents")
            
            if not retrieved_docs:
                print("Warning: No relevant documents found")
                
        except Exception as e:
            print(f"Error retrieving documents: {e}")
            raise
        
        # Generate response
        try:
            # Initialize prompt template
            prompt = ChatPromptTemplate.from_template(prompt_templates.prompts.food_recommender)
            
            # Combine document contents
            docs_content = "\n\n".join(doc.page_content for doc in retrieved_docs) if retrieved_docs else "No relevant documents found."
            print(f"Combined docs length: {len(docs_content)} characters")
            
            # Generate LLM response
            print("Generating LLM response...")
            messages = prompt.invoke({"question": query_text, "context": docs_content})
            
            response = model.invoke(messages)
            print("Response generated successfully")
            
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
        

class RecRquest(BaseModel):
    query: str
@router.post("/nutrition_scoring")
def nutrition_scoring(request: RecRquest):
    
    try:
        # Validate API key
        api_key = os.getenv("MODEL_API_KEY")
        if not api_key:
            raise ValueError("MODEL_API_KEY environment variable is not set")
        
        print(f"API Key loaded: {api_key[:10]}..." if len(api_key) > 10 else "API Key too short")
        
        # Initialize embeddings and model
        try:
            embeddings = DashScopeEmbeddings(
                model="text-embedding-v3", 
                dashscope_api_key=api_key
            )
            print("Embeddings initialized successfully")
        except Exception as e:
            print(f"Error initializing embeddings: {e}")
            raise
        
        try:
            model = ChatQwQ(
                model="qwen-max-latest",
                api_key=api_key,
                temperature=0.0,  
            )
            print("Model initialized successfully")
        except Exception as e:
            print(f"Error initializing model: {e}")
            raise
        
        # Load and process documents
        try:
            loader = TextLoader(r"data\Nutri Score.txt", encoding='utf-8')
            docs = loader.load()
            
            if not docs:
                raise ValueError("No documents loaded from file")
            
            raw_text = docs[0].page_content  # Ambil isi dokumen sebagai string
            
            # Split berdasarkan delimiter ===[TEXT]===
            chunks = raw_text.split("===[TEXT]===")
            
            # Bersihkan dan bungkus ulang jadi list of Document
            all_splits = [Document(page_content=chunk.strip()) for chunk in chunks if chunk.strip()]
            
            print(f"Successfully split documents into {len(all_splits)} chunks")
            
            if not all_splits:
                raise ValueError("No document chunks created")

        except Exception as e:
            print(f"Error processing documents: {e}")
            raise
        
        # Initialize PostgreSQL connection
        try:
            PG_HOST = os.getenv("PG_HOST")
            PG_PORT = os.getenv("PG_PORT") 
            PG_DATABASE = os.getenv("PG_DB")
            PG_USER = os.getenv("PG_USER")
            PG_PASSWORD = os.getenv("PG_PASSWORD")
            
            # Validate all required environment variables
            pg_vars = {
                "PG_HOST": PG_HOST,
                "PG_PORT": PG_PORT,
                "PG_DATABASE": PG_DATABASE,
                "PG_USER": PG_USER,
                "PG_PASSWORD": PG_PASSWORD
            }
            
            missing_vars = [k for k, v in pg_vars.items() if not v]
            if missing_vars:
                raise ValueError(f"Missing PostgreSQL environment variables: {missing_vars}")

            connection = f"postgresql+psycopg://{PG_USER}:{PG_PASSWORD}@{PG_HOST}:{PG_PORT}/{PG_DATABASE}"
            print("PostgreSQL connection string created")
            
        except Exception as e:
            print(f"Error setting up PostgreSQL connection: {e}")
            raise
        
        # Initialize vector store
        try:
            collection_name = "nutrition_scoring"
            vector_store = PGVector(
                embeddings=embeddings,
                collection_name=collection_name,
                connection=connection,
                use_jsonb=True,
            )
            print("Vector store initialized successfully")
            
        except Exception as e:
            print(f"Error initializing vector store: {e}")
            raise
        
        # Add documents to vector store (with batch processing)
        try:
            print(f"Indexing {len(all_splits)} chunks...")
            
            # Process in smaller batches to avoid memory issues
            batch_size = 50
            for i in range(0, len(all_splits), batch_size):
                batch = all_splits[i:i+batch_size]
                print(batch)
                print(f"Processing batch {i//batch_size + 1}: documents {i+1}-{min(i+batch_size, len(all_splits))}")
                vector_store.add_documents(documents=batch)
            
            print("Successfully indexed all chunks into vector store")
            
        except Exception as e:
            print(f"Error indexing documents: {e}")
            raise
        
        # Process query
        try:
            query_text = request.query

            if not query_text or not query_text.strip():
                raise ValueError("Query text is empty")
                
            print(f"Processing query: {query_text}")
            
        except Exception as e:
            print(f"Error processing query: {e}")
            raise
        
        # Retrieve similar documents
        try:
            print("Finding relevant documents...")
            # Reduce k and fetch_k to avoid potential issues
            retrieved_docs = vector_store.similarity_search(query_text, k=5)
            print(f"Found {len(retrieved_docs)} relevant documents")
            
            if not retrieved_docs:
                print("Warning: No relevant documents found")
                
        except Exception as e:
            print(f"Error retrieving documents: {e}")
            raise
        
        # Generate response
        try:
            # Initialize prompt template
            prompt = ChatPromptTemplate.from_template(prompt_templates.prompts.FOOD_SCORING)
            
            # Combine document contents
            docs_content = "\n\n".join(doc.page_content for doc in retrieved_docs) if retrieved_docs else "No relevant documents found."
            print(f"Combined docs length: {len(docs_content)} characters")
            
            # Generate LLM response
            print("Generating LLM response...")
            messages = prompt.invoke({"question": query_text, "context": docs_content})
            
            response = model.invoke(messages)
            print("Response generated successfully")
            
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