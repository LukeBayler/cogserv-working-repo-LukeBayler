---
topic:
  - sample
languages:
  - Java
products:
  - Azure
  - Cognitive Services
  - Computer Vision
---

# Sample Code for Describing Images with Computer Vision

This sample code shows you how to describe images with Computer Vision.

## Contents

| File/folder | Description |
|-------------|-------------|
| `README.md`            | This README file. |
| `src\main\java` | Java source file. |
| `src\main\resources` | Image file to describe. |

## Prerequisites

- Java development environment
- Maven

## Setup

- [Download this sample repository](https://github.com/LukeBayler/cognitive-services-samples/archive/master.zip).

## Running the sample

1. Store your Computer Vision API key in the `AZURE_COMPUTERVISION_API_KEY` environment variable.
2. Store your Azure endpoint in the `AZURE_ENDPOINT` environment variable.
3. Place an image file to be desrbied in the `src\main\resources` directory, and name it `describe-image.jpg`.

## Building and Running the Sample

1. From the command line, navigate to the samples root directory: `...\cognitive-services-samples\java\computer-vision\describe-image`.
2. Enter `mvn compile exec:java -Dexec.cleanupDaemonThreads=false`.

## Next steps

You can learn more about describing images with Computer Vision at the [official documentation site](https://docs.microsoft.com/en-us/azure/cognitive-services/computer-vision/concept-describing-images).
