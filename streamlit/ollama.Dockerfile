FROM python:3.12-slim

# Set the working directory
WORKDIR /app

# Install required system packages
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Install Python dependencies
RUN pip install --no-cache-dir streamlit llama-index ollama

COPY streamlit.py streamlit.py

CMD ["streamlit", "run", "streamlit.py"]

EXPOSE 8501
