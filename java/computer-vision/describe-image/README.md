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
| `DescribeImage.java` | Java source file. |
| `README.md`            | This README file. |

## Prerequisites

- Java development environment
- Maven

## Setup

- [Download this sample repository](https://github.com/LukeBayler/cognitive-services-samples/archive/master.zip).

## Running the sample

1. Update the `BaseURL` string with your region.
2. Update the `subKey` with your subscription key.
3. Update the `imgPath` string representing a file path to a location on your local machine that contains the image to be described.

## Building and Running the Sample

1. From the command line, navigate to the samples root directory: `...\cognitive-services-samples\java\computer-vision\describe-image`.
2. Enter `mvn compile exec:java -Dexec.cleanupDaemonThreads=false -Dexec.args="YOUR-SUBSCRIPTION-KEY"`.

## Next steps

You can learn more about describing images with Computer Vision at the [official documentation site](https://docs.microsoft.com/en-us/azure/cognitive-services/computer-vision/concept-describing-images).
