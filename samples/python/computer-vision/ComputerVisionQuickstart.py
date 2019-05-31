from azure.cognitiveservices.vision.computervision import ComputerVisionClient
# from azure.cognitiveservices.vision.computervision import VisualFeatureTypes

from azure.cognitiveservices.vision.computervision.models import TextOperationStatusCodes

from msrest.authentication import CognitiveServicesCredentials

import os
import sys
import time

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

# Describe a local image by:
#   1. Opening the binary file for reading.
#   2. Defining what to extract from the image by initializing an array of VisualFeatureTypes.
#   3. Calling the Computer Vision service's analyze_image_in_stream with the:
#      - image
#      - features to extract
#   4. Displaying the image captions and their confidence values.
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
remote_image_url = "https://raw.githubusercontent.com/Azure-Samples/cognitive-services-sample-data-files/master/ComputerVision/Images/landmark.jpg"
print("\n\nRemote image URL:\n" + remote_image_url)
#   END - Get a remote image for analysis

# Describe a remote image by:
#   1. Defining what to extract from the image by initializing an array of VisualFeatureTypes.
#   2. Calling the Computer Vision service's analyze_image with the:
#      - image URL
#      - features to extract
#   3. Displaying the image captions and their confidence values.
remote_image_features = ["description"]
remote_image_analysis = computervision_client.analyze_image(remote_image_url, remote_image_features)

print("\nCaptions from remote image: ")
if (len(remote_image_analysis.description.captions) == 0):
    print("No captions detected.")
else:
    for caption in remote_image_analysis.description.captions:
        print("'" + caption.text + "'" + " with confidence " + str(caption.confidence))
#   END - Describe a remote image


# Categorize a local image by:
#   1. Opening the binary file for reading.
#   2. Defining what to extract from the image by initializing an array of VisualFeatureTypes.
#   3. Calling the Computer Vision service's AnalyzeImageInStream with the:
#      - image
#      - features to extract
#   4. Displaying the image categories and their confidence values.
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

# Categorize a remote image by:
#   1. Calling the Computer Vision service's AnalyzeImage with the:
#      - image URL
#      - features to extract
#   2. Displaying the image categories and their confidence values.
remote_image_features = ["categories"]
remote_image_analysis = computervision_client.analyze_image(remote_image_url, remote_image_features)

print("\nCategories from remote image: ")
if (len(remote_image_analysis.categories) == 0):
    print("No categories detected.")
else:
    for category in remote_image_analysis.categories:
        print("'" + category.name + "'" + " with confidence " + str(category.score))
#   END - Categorize a remote image

# Tag a local image by:
#   1. Opening the binary file for reading.
#   2. Defining what to extract from the image by initializing an array of VisualFeatureTypes.
#   3. Calling the Computer Vision service's AnalyzeImageInStream with the:
#      - image
#      - features to extract
#   4. Displaying the image captions and their confidence values.
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

# Tag a remote image by:
#   1. Calling the Computer Vision service's AnalyzeImageInStream with the:
#      - image URL
#      - features to extract
#   2. Displaying the image captions and their confidence values.
remote_image_features = ["tags"]
remote_image_analysis = computervision_client.analyze_image(remote_image_url, remote_image_features)

print("\nTags in the remote image: ")
if (len(remote_image_analysis.tags) == 0):
    print("No tags detected.")
else:
    for tag in remote_image_analysis.tags:
        print("'" + tag.name + "'" + " with confidence " + str(tag.confidence))
#   END - Tag a remote image

# Detect faces in a local image by:
#   1. Opening the binary file for reading.
#   2. Defining what to extract from the image by initializing an array of VisualFeatureTypes.
#   3. Calling the Computer Vision service's AnalyzeImageInStream with the:
#      - image
#      - features to extract
#   4. Displaying the image captions and their confidence values.
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

# Detect faces in a remote image by:
#   1. Calling the Computer Vision service's AnalyzeImageInStream with the:
#      - image URL
#      - features to extract
#   2. Displaying the image captions and their confidence values.
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

# Detect adult or racy content in a local image by:
#   1. Opening the binary file for reading.
#   2. Defining what to extract from the image by initializing an array of VisualFeatureTypes.
#   3. Calling the Computer Vision service's AnalyzeImageInStream with the:
#      - image
#      - features to extract
#   4. Displaying the image captions and their confidence values.
local_image = open(local_image_path, "rb")
local_image_features = ["adult"]
local_image_analysis = computervision_client.analyze_image_in_stream(local_image, local_image_features)

print("\nAnalyzing local image for adult or racy content ... ")
print("Is adult content: " + str(local_image_analysis.adult.is_adult_content) + " with confidence " + str(local_image_analysis.adult.adult_score))
print("Has racy content: " + str(local_image_analysis.adult.is_racy_content) + " with confidence " + str(local_image_analysis.adult.racy_score))
#   END - Detect adult or racy content in a local image

# Detect adult or racy content in a remote image by:
#   1. Calling the Computer Vision service's AnalyzeImageInStream with the:
#      - image URL
#      - features to extract
#   2. Displaying the image captions and their confidence values.
remote_image_features = ["adult"]
remote_image_analysis = computervision_client.analyze_image(remote_image_url, remote_image_features)

print("\nAnalyzing remote image for adult or racy content ... ")
print("Is adult content: " + str(local_image_analysis.adult.is_adult_content) + " with confidence " + str(local_image_analysis.adult.adult_score))
print("Has racy content: " + str(local_image_analysis.adult.is_racy_content) + " with confidence " + str(local_image_analysis.adult.racy_score))
#   END - Detect adult or racy content in a remote image

#   Detect the color scheme in a local image
local_image = open(local_image_path, "rb")
local_image_features = ["color"]
local_image_analysis = computervision_client.analyze_image_in_stream(local_image, local_image_features)

print("\nColor scheme of the local image: ");
print("Is black and white: " + str(local_image_analysis.color.is_bw_img))
print("Accent color: 0x" + local_image_analysis.color.accent_color)
print("Dominant background color: " + local_image_analysis.color.dominant_color_background)
print("Dominant foreground color: " + local_image_analysis.color.dominant_color_foreground)
print("Dominant colors: " + str(local_image_analysis.color.dominant_colors))
#   END - Detect the color scheme in a local image

#   Detect the color scheme in a remote image
remote_image_features = ["color"]
remote_image_analysis = computervision_client.analyze_image(remote_image_url, remote_image_features)

print("\nColor scheme of the local image: ");
print("Is black and white: " + str(remote_image_analysis.color.is_bw_img))
print("Accent color: 0x" + remote_image_analysis.color.accent_color)
print("Dominant background color: " + remote_image_analysis.color.dominant_color_background)
print("Dominant foreground color: " + remote_image_analysis.color.dominant_color_foreground)
print("Dominant colors: " + str(remote_image_analysis.color.dominant_colors))
#   END - Detect the color scheme in a remote image

#   Detect domain-specific content (celebrities/landmarks) in a local image
local_image = open(local_image_path, "rb")
local_image_features = ["description", "categories"]
local_image_analysis = computervision_client.analyze_image_in_stream(local_image, local_image_features)

print("\nCelebrities in the local image:")
for category in local_image_analysis.categories:
    if (category.detail != None and category.detail.celebrities != None):
        for celeb in category.detail.celebrities:
            print(celeb.name + " with confidence " + str(celeb.confidence) + " at location " + \
            str(celeb.face_rectangle.left) + ", " + str(celeb.face_rectangle.top) + ", " + \
            str(celeb.face_rectangle.left + celeb.face_rectangle.width) + ", " + \
            str(celeb.face_rectangle.top + celeb.face_rectangle.height))

print("\nLandmarks in the local image:")
for category in local_image_analysis.categories:
    if (category.detail != None and category.detail.landmarks != None):
        for landmark in category.detail.landmarks:
            print("'" + landmark.name + "'" + " with confidence " + str(landmark.confidence))

#   END Detect domain-specific content (celebrities/landmarks) in a local image

#   Detect domain-specific content (celebrities/landmarks) in a remote image
remote_image_features = ["description", "categories"]
remote_image_analysis = computervision_client.analyze_image(remote_image_url, remote_image_features)

print("\nCelebrities in the remote image:")
for category in remote_image_analysis.categories:
    if (category.detail != None and category.detail.celebrities != None):
        for celeb in category.detail.celebrities:
            print(celeb.name + " with confidence " + str(celeb.confidence) + " at location " + \
            str(celeb.face_rectangle.left) + ", " + str(celeb.face_rectangle.top) + ", " + \
            str(celeb.face_rectangle.left + celeb.face_rectangle.width) + ", " + \
            str(celeb.face_rectangle.top + celeb.face_rectangle.height))

print("\nLandmarks in the remote image:")
for category in remote_image_analysis.categories:
    if (category.detail != None and category.detail.landmarks != None):
        for landmark in category.detail.landmarks:
            print("'" + landmark.name + "'" + " with confidence " + str(landmark.confidence))
#   END Detect domain-specific content (celebrities/landmarks) in a remote image

#   Detect image types (clip art/line drawing) of a local image
local_image = open(local_image_path, "rb")
local_image_features = ["imagetype"]
local_image_analysis = computervision_client.analyze_image_in_stream(local_image, local_image_features)

print("\nImage type of local image:")
print("Clip art type: " + str(local_image_analysis.image_type.clip_art_type))
print("Line drawing type: " + str(local_image_analysis.image_type.line_drawing_type))

#   END - Detect image types (clip art/line drawing) of a local image

#   Detect image types (clip art/line drawing) of a remote image
remote_image_features = ["imagetype"]
remote_image_analysis = computervision_client.analyze_image(remote_image_url, remote_image_features)

print("\nImage type of remote image:")
print("Clip art type: " + str(remote_image_analysis.image_type.clip_art_type))
print("Line drawing type: " + str(remote_image_analysis.image_type.line_drawing_type))
#   END - Detect image types (clip art/line drawing) of a remote image

#   Detect objects in a local image
local_image = open(local_image_path, "rb")
local_image_analysis = computervision_client.detect_objects_in_stream(local_image)

print("\nDetecting objects in local image:")
if local_image_analysis.objects == None:
    print("No objects detected.")
else:
    for object in local_image_analysis.objects:
        print("object at location " + \
        str(object.rectangle.x) + ", " + str(object.rectangle.x + object.rectangle.w) + ", " + \
        str(object.rectangle.y) + ", " + str(object.rectangle.y + object.rectangle.h))
#   END - Detect objects in a local image

#   Detect objects in a remote image
remote_image_analysis = computervision_client.analyze_image(remote_image_url, remote_image_features)

print("\nDetecting objects in remote image:")
if remote_image_analysis.objects == None:
    print("No objects detected.")
else:
    for object in remote_image_analysis.objects:
        print("object at location " + \
        str(object.rectangle.x) + ", " + str(object.rectangle.x + object.rectangle.w) + ", " + \
        str(object.rectangle.y) + ", " + str(object.rectangle.y + object.rectangle.h))
#   END - Detect objects in a remote image

#   Detect brands in a local image
local_image_path = "resources\\gray-shirt-logo.jpg"
local_image = open(local_image_path, "rb")
local_image_features = ["brands"]
local_image_analysis = computervision_client.analyze_image_in_stream(local_image, local_image_features)

print("\nDetecting brands in local image: ")
if len(local_image_analysis.brands) == 0:
    print("No brands detected.")
else:
    for brand in local_image_analysis.brands:
        print("'" + brand.name + " with confidence " + str(brand.confidence) + " at location " + \
        str(brand.rectangle.x) + ", " + str(brand.rectangle.x + brand.rectangle.w) + ", " + \
        str(brand.rectangle.y) + ", " + str(brand.rectangle.y + brand.rectangle.h))
#   END - Detect brands in a local image

#   Detect brands in a remote image
remote_image_url = "https://docs.microsoft.com/en-us/azure/cognitive-services/computer-vision/images/gray-shirt-logo.jpg"
remote_image_features = ["brands"]
remote_image_analysis = computervision_client.analyze_image(remote_image_url, remote_image_features)

print("\nDetecting brands in remote image: ")
if len(remote_image_analysis.brands) == 0:
    print("No brands detected.")
else:
    for brand in remote_image_analysis.brands:
        print("'" + brand.name + " with confidence " + str(brand.confidence) + " at location " + \
        str(brand.rectangle.x) + ", " + str(brand.rectangle.x + brand.rectangle.w)  + ", " + \
        str(brand.rectangle.y) + ", " + str(brand.rectangle.y + brand.rectangle.h))
#   END - Detect brands in a remote image

#   Recognizing printed and handwritten text with the batch read API in a local image
local_image_path = "resources\\handwritten_text.jpg"
local_image = open(local_image_path, "rb")
text_recognition_mode = "handwritten"
num_chars_in_operation_id = 36

client_response = computervision_client.batch_read_file_in_stream(local_image, text_recognition_mode, raw=True)
operation_location = client_response.headers["Operation-Location"]
id_location = len(operation_location) - num_chars_in_operation_id
operation_id = operation_location[id_location:]

print("\nRecognizing text in a local image with the batch Read API ... \n")

while True:
    result = computervision_client.get_read_operation_result(operation_id)
    if result.status not in ['NotStarted', 'Running']:
        break
    time.sleep(1)

if result.status == TextOperationStatusCodes.succeeded:
    for text_result in result.recognition_results:
        for line in text_result.lines:
            print(line.text)
            print(line.bounding_box)
            print()
#   END - Recognizing printed and handwritten text with the batch read API in a local image

#   Recognizing printed and handwritten text with the batch read API in a remote image
remote_image_url = "https://raw.githubusercontent.com/Azure-Samples/cognitive-services-sample-data-files/master/ComputerVision/Images/printed_text.jpg"
text_recognition_mode = "printed"
num_chars_in_operation_id = 36

client_response = computervision_client.batch_read_file(remote_image_url, text_recognition_mode, raw=True)
operation_location = client_response.headers["Operation-Location"]
id_location = len(operation_location) - num_chars_in_operation_id
operation_id = operation_location[id_location:]

print("\nRecognizing text in a remote image with the batch Read API ... \n")

while True:
    result = computervision_client.get_read_operation_result(operation_id)
    if result.status not in ['NotStarted', 'Running']:
        break
    time.sleep(1)

if result.status == TextOperationStatusCodes.succeeded:
    for text_result in result.recognition_results:
        for line in text_result.lines:
            print(line.text)
            print(line.bounding_box)
            print()
#   END - Recognizing printed and handwritten text with the batch read API in a remote image
