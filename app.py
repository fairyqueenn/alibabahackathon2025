import os
import dashscope

dashscope.base_http_api_url = 'https://dashscope-intl.aliyuncs.com/api/v1'
messages = [
    {
        "role": "user",
        "content": [
            {"image": "https://dashscope.oss-cn-beijing.aliyuncs.com/images/dog_and_girl.jpeg"},
            {"image": "https://dashscope.oss-cn-beijing.aliyuncs.com/images/tiger.png"},
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
print(response)