import streamlit as st
import time
from TTS.api import TTS

st.title("GOTO TTS Streamlit Frontend")
st.header('Text to speech generation')

models = ["tts_models/en/ljspeech/tacotron2-DDC", "tts_models/en/ljspeech/tacotron2-DDC_ph", "tts_models/en/ljspeech/glow-tts", "tts_models/en/ljspeech/speedy-speech", "tts_models/en/ljspeech/tacotron2-DCA", "tts_models/en/ljspeech/vits", "tts_models/en/ljspeech/vits--neon", "tts_models/en/ljspeech/fast_pitch", "tts_models/en/ljspeech/overflow", "tts_models/en/ljspeech/neural_hmm", "tts_models/en/sam/tacotron-DDC"]
model = st.selectbox('Choose a model', models)
tts = TTS(model_name=model, progress_bar=True, gpu=False)
text = st.text_area('Enter text to convert to audio format!!')
speed = st.slider('Speed', 0.1, 1.99, 1.0, 0.01)
if st.button('Convert text to audio!'):
    # Run TTS
    tts.tts_to_file(text=text, speed=speed, file_path="out.wav")
    st.success('Converted to audio successfully')

    audio_file = open('out.wav', 'rb')
    audio_bytes = audio_file.read()
    st.audio(audio_bytes, format='audio/wav')
    st.success("You can now play the audio by clicking on the play button!!")
