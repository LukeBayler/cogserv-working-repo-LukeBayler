---
topic:
  - sample
languages:
  - Java
products:
  - Azure
  - Cognitive Services
  - Content Moderator
---

# Sample Code for Moderating Images with Content Moderator

This sample code shows you how to moderate images with Content Moderator.

## Contents

| File/folder | Description |
|-------------|-------------|
| `src\main\java` | Java source code. |
| `README.md`            | This README file. |
| `src\main\Resources\ImageFiles.txt`       | URLs for the images to moderate. |
| `src\main\Resources\ModerationOutput.json`| Program output. The sample also writes to standard output. |

## Prerequisites

- Java development environment
- Maven

## Setup

- [Download this sample repository](https://github.com/LukeBayler/cognitive-services-samples/archive/master.zip).

## Modifying the Sample for your Configuration

1. Store your Computer Vision API key in the `AZURE_COMPUTERVISION_API_KEY` environment variable.
2. Store your Azure endpoint in the `AZURE_ENDPOINT` environment variable.

## Building and Running the Sample

1. From the command line, navigate to the samples root directory: `...\cognitive-services-samples\java\content-moderator\moderate-images`.
2. Enter `mvn compile exec:java -Dexec.cleanupDaemonThreads=false`.

## Next steps

You can learn more about image moderation with Content Moderator at the [official documentation site](https://docs.microsoft.com/en-us/azure/cognitive-services/content-moderator/image-moderation-api).
