# Streamlit

Installation and getting started: https://docs.streamlit.io/get-started/installation/command-line


### Usage

the `app-tts.py` are using [TTS](https://docs.coqui.ai/en/latest/installation.html), a model for synthesizing speech.

reference of the script is taken from [here](https://github.com/Vidyut/vidyut-tts/tree/main).


### How to deploy the model in KIND cluster

1. Build docker image with provided `Dockerfile`

```
docker build -t streamlit-tts .
```

2. Create image archive from the docker image

```
docker save streamlit-tts:latest -o streamlit-tts.tar
```

3. Load the image archive to the KIND cluster

```
# if you use the default name for the cluster, omit --name option
kind load image-archive streamlit-tts.tar --name=<kind-cluster-name>
```

4. Deploy the model to the k8s cluste

```
kubectl apply -f streamlit-tts-deployment.yaml 
```

5. Check and confirm if the apps is successfully deployed

```
kubectl port-forward service/streamlit-tts 8501:8501
```

6. Access the app via `http://localhost:8501`

### Cheatsheet

##### Get list of images in KIND cluster

```
docker exec -it my-cluster-control-plane crictl images
```
