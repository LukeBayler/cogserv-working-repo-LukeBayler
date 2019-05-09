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

# Sample Code for Detecting Color Schemes with Computer Vision

This sample code shows you how to detect color schemes with the Computer Vision API.

## Contents

| File/folder | Description |
|-------------|-------------|
| `README.md`            | This README file. |
| `src\main\java` | Contains the Java source file. |
| `src\main\resources` | Contains an image file. |

## Prerequisites

- Java development environment
- Maven

## Setup

- [Download and extract the sample repository](https://github.com/LukeBayler/cognitive-services-samples/archive/master.zip).

## Running the sample

1. Store your API key in the `AZURE_COMPUTERVISION_API_KEY` environment variable.
2. Store your Azure endpoint in the `AZURE_ENDPOINT` environment variable.
3. Find an image file for the Computer Vision API to describe.
4. Rename the image file `analyze-image.jpg`.
5. Place the image file in the `src\main\resources` directory.

## Building and Running the Sample

1. Open a command prompt and navigate to the directory where you extracted the sample repository.
2. Navigate to the subdirectory containing this sample, which is: `cognitive-services-samples\java\computer-vision\analyze-image`.
2. Enter `mvn compile exec:java -Dexec.cleanupDaemonThreads=false`.

## Next steps

For more information about detecting color schemes with the Computer Vision API, visit the [official documentation site](https://docs.microsoft.com/en-us/azure/cognitive-services/computer-vision/concept-detecting-color-schemes.
