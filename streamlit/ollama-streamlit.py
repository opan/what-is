import streamlit as st
import logging
import time
from ollama import Client
import os

logging.basicConfig(level=logging.INFO)


# Define custom client for Ollama
client = Client(host=os.getenv("OLLAMA_HOST", "http://localhost:11434"))

# Initialize chat history in session state if not already present
if 'messages' not in st.session_state:
    st.session_state.messages = []

# Function to stream chat response based on selected model
def stream_chat(model, messages):
    try:

        # # Initialize the language model with a timeout
        resp = client.chat(model=model, messages=messages, stream=True)

        response = ""
        response_placeholder = st.empty()
        # Append each piece of the response to the output
        for r in resp:
            response += r['message']['content']
            response_placeholder.write(response)

        # Log the interaction details
        logging.info(f"Model: {model}, Messages: {messages}, Response: {response}")
        return response
    except Exception as e:
        # Log and re-raise any errors that occur
        logging.error(f"Error during streaming: {str(e)}")
        raise e

def main():
    st.title("GoGeeks 2.0: Chat with LLMs Models")  # Set the title of the Streamlit app
    logging.info("App started")  # Log that the app has started

    # Sidebar for model selection
    # model = st.sidebar.selectbox("Choose a model", ["llama3", "phi3", "mistral"])
    # logging.info(f"Model selected: {model}")
    model = "llama3.1:8b"

    logging.info(f"Pulling model: {model}")
    client.pull(model)

    # Prompt for user input and save to chat history
    if prompt := st.chat_input("Your question"):
        st.session_state.messages.append({"role": "user", "content": prompt})
        logging.info(f"User input: {prompt}")

        # Display the user's query
        for message in st.session_state.messages:
            with st.chat_message(message["role"]):
                st.write(message["content"])

        # Generate a new response if the last message is not from the assistant
        if st.session_state.messages[-1]["role"] != "assistant":
            with st.chat_message("assistant"):
                start_time = time.time()  # Start timing the response generation
                logging.info("Generating response")

                with st.spinner("Writing..."):
                    try:
                        # Prepare messages for the LLM and stream the response
                        # messages = [ChatMessage(role=msg["role"], content=msg["content"]) for msg in st.session_state.messages]
                        messages = [{'role': msg["role"], 'content': msg["content"]} for msg in st.session_state.messages]
                        response_message = stream_chat(model, messages)
                        duration = time.time() - start_time  # Calculate the duration
                        response_message_with_duration = f"{response_message}\n\nDuration: {duration:.2f} seconds"
                        st.session_state.messages.append({"role": "assistant", "content": response_message_with_duration})
                        st.write(f"Duration: {duration:.2f} seconds")
                        logging.info(f"Response: {response_message}, Duration: {duration:.2f} s")

                    except Exception as e:
                        # Handle errors and display an error message
                        st.session_state.messages.append({"role": "assistant", "content": str(e)})
                        st.error("An error occurred while generating the response.")
                        logging.error(f"Error: {str(e)}")

if __name__ == "__main__":
    main()
