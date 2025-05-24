from langchain_qwq import ChatQwQ

llm = ChatQwQ(
    model="qwq-plus",
    max_tokens=3_000,
    timeout=None,
    max_retries=2,
    # other params...
)

messages = [
    (
        "system",
        "You are a helpful assistant that translates English to French."
        "Translate the user sentence.",
    ),
    ("human", "I love programming."),
]
ai_msg = llm.invoke(messages)
ai_msg