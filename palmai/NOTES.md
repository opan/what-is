You can customize how the PaLM API behaves in response to your prompt by using the following parameters for text-bison:

- temperature: higher means more "creative" responses
- max_output_tokens: sets the max number of tokens in the output
- top_p: higher means it will pull from more possible next tokens, based on cumulative probability
- top_k: higher means it will sample from more possible next tokens


The temperature parameter (range: 0.0 - 1.0, default 0)
What is temperature?
The temperature is used for sampling during the response generation, which occurs when top_p and top_k are applied. Temperature controls the degree of randomness in token selection.

How does temperature affect the response?
Lower temperatures are good for prompts that require a more deterministic and less open-ended response. In comparison, higher temperatures can lead to more "creative" or diverse results. A temperature of 0 is deterministic: the highest probability response is always selected. For most use cases, try starting with a temperature of 0.2.

A higher temperature value will result in a more explorative output, with a higher likelihood of generating rare or unusual words or phrases. Conversely, a lower temperature value will result in a more conservative output, with a higher likelihood of generating common or expected words or phrases.

Example:
For example,

`temperature = 0.0`:

- The cat sat on the couch, watching the birds outside.
- The cat sat on the windowsill, basking in the sun.

`temperature = 0.9`:

- The cat sat on the moon, meowing at the stars.
- The cat sat on the cheeseburger, purring with delight.
- 
Note: It's important to note that while the temperature parameter can help generate more diverse and interesting text, it can also increase the likelihood of generating nonsensical or inappropriate text (i.e. hallucinations). Therefore, it's important to use it carefully and with consideration for the desired outcome.

For more information on the temperature parameter for text models, please refer to the [documentation on model parameters.](https://cloud.google.com/vertex-ai/generative-ai/docs/learn/models#text_model_parameters)


If you run the following cell multiple times, it should always return the same response, as `temperature=0` is deterministic.


```
temp_val = 0.0
prompt_temperature = "Complete the sentence: As I prepared the picture frame, I reached into my toolkit to fetch my:"

response = generation_model.predict(
    prompt=prompt_temperature,
    temperature=temp_val,
)

print(f"[temperature = {temp_val}]")
print(response.text)
```

If you run the following cell multiple times, it may return different responses, as higher temperature values can lead to more diverse results, even though the prompt is the same as the above cell.


```
temp_val = 1.0

response = generation_model.predict(
    prompt=prompt_temperature,
    temperature=temp_val,
)

print(f"[temperature = {temp_val}]")
print(response.text)
```