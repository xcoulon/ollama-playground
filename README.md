# Ollama Playground

## Installation on macOS

[Download Ollama](https://ollama.com/download) and copy the `ollama.app` file in the Applications folder, then run the application and follow the instructions to install the `ollama` CLI.

## Pulling IBM's granite-code 3b model

[Granite Code](https://ollama.com/library/granite-code) is a family of decoder-only code model designed for code generative tasks (e.g. code generation, code explanation, code fixing, etc.).

```shell
$ ollama pull granite-code:3b
pulling manifest 
pulling 9699e81546fa... 100% ▕██████████████████████████████████████████████████████████████▏ 2.0 GB                         
pulling 977871d28ce4... 100% ▕██████████████████████████████████████████████████████████████▏  679 B                         
pulling 43070e2d4e53... 100% ▕██████████████████████████████████████████████████████████████▏  11 KB                         
pulling f1406f0e20d0... 100% ▕██████████████████████████████████████████████████████████████▏   43 B                         
pulling 1e60cb6006a8... 100% ▕██████████████████████████████████████████████████████████████▏  559 B                         
verifying sha256 digest 
writing manifest 
success 
```

