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
| `ImageFiles.txt`       | URLs for the images to moderate. |
| `ImageModeration.java` | Sample code output. |
| `ModerationOutput.json`| Sample source code. |
| `README.md`            | This README file. |

## Prerequisites

- Java development environment
- Jar files required by Content Moderator

## Setup

- Clone or download this sample repository.

## Running the sample

1. Update the `AzureBaseURL` string with your region.
2. Update the `CMSubscriptionKey` with your subscription key.
3. Update the `ImageUrlFile` string representing a file path to a location on your local machine. This file contains URLs for the images to moderate.
4. Update the `OutputFile` string representing a file path to a location on your local machine. The sample code writes its output to this file. The sample also writes its output to the standard output stream.

## Next steps

You can learn more about image moderation with Content Moderator at the [official documentation site](https://docs.microsoft.com/en-us/azure/cognitive-services/content-moderator/image-moderation-api).
