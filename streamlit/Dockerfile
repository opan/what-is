# Use the official Python 3.10 slim image as a base image
FROM python:3.10-slim

# Set the working directory
WORKDIR /app

# Install required system packages
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Install Python dependencies
RUN pip install --no-cache-dir TTS streamlit

# Copy the application code (if you have any)
COPY app-tts.py app-tts.py

# Specify the command to run the application
# CMD ["python3", "your_script.py"]

# Example CMD if you are running a Streamlit app
CMD ["streamlit", "run", "app-tts.py"]

# Expose the necessary port (e.g., if using Streamlit)
EXPOSE 8501
