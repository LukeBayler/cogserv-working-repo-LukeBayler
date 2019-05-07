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

# Sample Code for Recognizing Handwritten Text with Computer Vision

This sample code shows you how to recognize handwritten text with Computer Vision.

## Contents

| File/folder | Description |
|-------------|-------------|
| `RecognizeHandwrittenText.java` | Java source file. |
| `README.md`            | This README file. |

## Prerequisites

- Java development environment
- Maven

## Setup

- Clone or download this sample repository.

## Running the sample

1. Update the `AzureBaseURL` string with your region.
2. Update the `CMSubscriptionKey` with your subscription key.
3. Update the `imagePath` string representing a file path to a location on your local machine that contains the image with handwritten text.

## Building and Running the Sample

1. From the command line, navigate to the samples root directory: `...\cognitive-services-samples\java\computer-vision\recognize-handwritten-text`.
2. Enter `mvn compile exec:java -Dexec.cleanupDaemonThreads=false`.

## Next steps

You can learn more about recognizing handwritten text with Computer Vision at the [official documentation site](https://docs.microsoft.com/en-us/azure/cognitive-services/computer-vision/concept-recognizing-text).
