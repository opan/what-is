#!/bin/sh
while ! nc -z ollama 11434; do
  echo "Waiting for ollama..."
  sleep 2
done

# Start the main application
exec streamlit run app.py
