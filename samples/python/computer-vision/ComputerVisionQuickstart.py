from azure.cognitiveservices.vision.computervision import ComputerVisionClient
# from azure.cognitiveservices.vision.computervision import VisualFeatureTypes
from msrest.authentication import CognitiveServicesCredentials

import os
import sys

#   The Quickstarts in this file are for the Computer Vision API for Microsoft
#   Cognitive Services. In this file are Quickstarts for the following tasks:
#     - Describing images
#     - Categorizing images
#     - Tagging images
#     - Detecting faces
#     - Detecting adult or racy content
#     - Detecting the color scheme
#     - Detecting domain-specific content (celebrities/landmarks)
#     - Detecting image types (clip art/line drawing)
#     - Detecting objects
#     - Detecting brands
#     - Recognizing printed and handwritten text with the batch read API


#   Configure the Computer Vision client by:
#     1. Reading the Computer Vision API key and the Azure region from environment
#        variables (COMPUTERVISION_API_KEY and COMPUTERVISION_REGION), which must
#        be set prior to running this code. After setting the	environment variables,
#        restart your command shell or your IDE.
#     2. Constructing the endpoint URL from the base URL and the Azure region.
#     3. Setting up the authorization on the client with the subscription key.
#     4. Getting the context.
if 'COMPUTERVISION_API_KEY' in os.environ:
    computervision_api_key = os.environ['COMPUTERVISION_API_KEY']
else:
    print("\nPlease set the COMPUTERVISION_API_KEY environment variable.\n**Note that you might need to restart your shell or IDE.**")
    sys.exit()

if 'COMPUTERVISION_REGION' in os.environ:
    computervision_region = os.environ['COMPUTERVISION_REGION']
else:
    print("\nPlease set the COMPUTERVISION_REGION environment variable.\n**Note that you might need to restart your shell or IDE.**")
    sys.exit()

endpoint_url = "https://" + computervision_region + ".api.cognitive.microsoft.com"

computervision_client = ComputerVisionClient(endpoint_url, CognitiveServicesCredentials(computervision_api_key))
#	END - Configure the Computer Vision client

#   Get a local image for analysis
local_image_path = "resources\\faces.jpg"
print("\n\nLocal image path:\n" + os.getcwd() + local_image_path)
#   END - Get a local image for analysis

#   Describe a local image
local_image = open(local_image_path, "rb")
local_image_features = ["description"]
local_image_analysis = computervision_client.analyze_image_in_stream(local_image, local_image_features)

print("\nCaptions from local image: ")
if (len(local_image_analysis.description.captions) == 0):
    print("No captions detected.")
else:
    for caption in local_image_analysis.description.captions:
        print("'" + caption.text + "'" + " with confidence " + str(caption.confidence))
#  END - Describe a local image

#   Get a remote image for analysis
remote_image_url = "https://raw.githubusercontent.com/Azure-Samples/cognitive-services-python-sdk-samples/master/samples/vision/images/house.jpg"
print("\n\nRemote image URL:\n" + remote_image_url)
#   END - Get a remote image for analysis

#   Describe a remote image
remote_image_features = ["description"]
remote_image_analysis = computervision_client.analyze_image(remote_image_url, remote_image_features)

print("\nCaptions from remote image: ")
if (len(remote_image_analysis.description.captions) == 0):
    print("No captions detected.")
else:
    for caption in remote_image_analysis.description.captions:
        print("'" + caption.text + "'" + " with confidence " + str(caption.confidence))
#   END - Describe a remote image


#   Categorizing a local image
local_image = open(local_image_path, "rb")
local_image_features = ["categories"]
local_image_analysis = computervision_client.analyze_image_in_stream(local_image, local_image_features)

print("\nCategories from local image: ")
if (len(local_image_analysis.categories) == 0):
    print("No categories detected.")
else:
    for category in local_image_analysis.categories:
        print("'" + category.name + "'" + " with confidence " + str(category.score))
#   END - Categorize a local image

#   Categorizing a remote image
remote_image_features = ["categories"]
remote_image_analysis = computervision_client.analyze_image(remote_image_url, remote_image_features)

print("\nCategories from remote image: ")
if (len(remote_image_analysis.categories) == 0):
    print("No categories detected.")
else:
    for category in remote_image_analysis.categories:
        print("'" + category.name + "'" + " with confidence " + str(category.score))
#   END - Categorize a remote image

#   Tag a local image
local_image = open(local_image_path, "rb")
local_image_features = ["tags"]
local_image_analysis = computervision_client.analyze_image_in_stream(local_image, local_image_features)

print("\nTags in the local image: ")
if (len(local_image_analysis.tags) == 0):
    print("No tags detected.")
else:
    for tag in local_image_analysis.tags:
        print("'" + tag.name + "'" + " with confidence " + str(tag.confidence))
#   END - Tag a local image

#   Tag a remote image
remote_image_features = ["tags"]
remote_image_analysis = computervision_client.analyze_image(remote_image_url, remote_image_features)

print("\nTags in the remote image: ")
if (len(remote_image_analysis.tags) == 0):
    print("No tags detected.")
else:
    for tag in remote_image_analysis.tags:
        print("'" + tag.name + "'" + " with confidence " + str(tag.confidence))
#   END - Tag a remote image

#   Detect faces in a local image
local_image = open(local_image_path, "rb")
local_image_features = ["faces"]
local_image_analysis = computervision_client.analyze_image_in_stream(local_image, local_image_features)

print("\nFaces in the local image: ")
if (len(local_image_analysis.faces) == 0):
    print("No faces detected.")
else:
    for face in local_image_analysis.faces:
        print("'" + face.gender + "'" + " of age " + str(face.age) + " at location " \
        + str(face.face_rectangle.left) + ", " + str(face.face_rectangle.top) + ", " \
        + str(face.face_rectangle.left + face.face_rectangle.width) + ", " \
        + str(face.face_rectangle.top + face.face_rectangle.height))
#   END - Detect faces in a local image

#   Detect faces in a remote image
remote_image_features = ["faces"]
remote_image_analysis = computervision_client.analyze_image(remote_image_url, remote_image_features)

print("\nFaces in the remote image: ")
if (len(remote_image_analysis.faces) == 0):
    print("No faces detected.")
else:
    for face in remote_image_analysis.faces:
        print("'" + face.gender + "'" + " of age " + str(face.age) + " at location " \
        + str(face.face_rectangle.left) + ", " + str(face.face_rectangle.top) + ", " \
        + str(face.face_rectangle.left + face.face_rectangle.width) + ", " \
        + str(face.face_rectangle.top + face.face_rectangle.height))
#   END - Detect faces in a remote image

#   Detect adult or racy content in a local image
#   END - Detect adult or racy content in a local image

#   Detect adult or racy content in a remote image
#   END - Detect adult or racy content in a remote image

#   Detect the color scheme in a local image
#   END - Detect the color scheme in a local image

#   Detect the color scheme in a remote image
#   END - Detect the color scheme in a remote image

#   Detect domain-specific content (celebrities/landmarks) in a local image
#   END Detect domain-specific content (celebrities/landmarks) in a local image

#   Detect domain-specific content (celebrities/landmarks) in a remote image
#   END Detect domain-specific content (celebrities/landmarks) in a remote image

#   Detect image types (clip art/line drawing) of a local image
#   END - Detect image types (clip art/line drawing) of a local image

#   Detect image types (clip art/line drawing) of a remote image
#   END - Detect image types (clip art/line drawing) of a remote image

#   Detect objects in a local image
#   END - Detect objects in a local image

#   Detect objects in a remote image
#   END - Detect objects in a remote image

#   Detect brands in a local image
#   END - Detect brands in a local image

#   Detect brands in a remote image
#   END - Detect brands in a remote image

#   Recognizing printed and handwritten text with the batch read API in a local image
#   END - Recognizing printed and handwritten text with the batch read API in a local image

#   Recognizing printed and handwritten text with the batch read API in a remote image
#   END - Recognizing printed and handwritten text with the batch read API in a remote image
