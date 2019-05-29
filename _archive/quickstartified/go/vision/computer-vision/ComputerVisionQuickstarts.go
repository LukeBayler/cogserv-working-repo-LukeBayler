package main

import (
	"context"
	"fmt"
  "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.0/computervision"
	"github.com/Azure/go-autorest/autorest"
	"io"
	"log"
	"os"
	// "path"
	"strings"
)

/*  This Quickstart for the Azure Cognitive Services Computer Vision API shows how to analyze
 *	an image both locally and from a URL.
 *
 *  Analyzing an image includes:
 *  - Displaying image captions and confidence values
 *  - Displaying image category names and confidence values
 *  - Displaying image tags and confidence values
 *  - Displaying any faces found in the image and their bounding boxes
 *  - Displaying whether any adult or racy content was detected and the confidence values
 *  - Displaying the image color scheme
 *  - Displaying any celebrities detected in the image and their bounding boxes
 *  - Displaying any landmarks detected in the image and their bounding boxes
 *  - Displaying what type of clip art or line drawing the image is
 */

func main() {
	/*  Configure the local environment:
 	*
 	*  Set the AZURE_COMPUTERVISION_API_KEY and AZURE_REGION environment variables on your
 	*  local machine using the appropriate method for your preferred command shell.
 	*
 	*  For AZURE_REGION, use the same region you used to get your subscription keys.
 	*
 	*  Note that:
 	*		- Environment variables cannot contain quotation marks, so the quotation marks
 	*  		are included in the code below to stringify them.
 	*		- After setting these environment variables in your preferred command shell,
 	*  		you will need to close and then re-open your command shell.
 	*/
	azureComputerVisionAPIKey := os.Getenv("AZURE_COMPUTERVISION_API_KEY")
	if ("" == azureComputerVisionAPIKey) {
		log.Fatal("Please set the AZURE_COMPUTERVISION_API_KEY environment variable. Note that you might need to restart your shell or IDE.")
	}

	 azureRegion := os.Getenv("AZURE_REGION")
	 if ("" == azureRegion) {
		 log.Fatal("Please set the AZURE_REGION environment variable. Note that you might need to restart your shell or IDE.")
	 }
	 //  END - Configure the local environment.

	fmt.Println("\nAzure Cognitive Services Computer Vision - Go Quickstart Sample")

	// Get the context, which is required by the SDK methods.
	computerVisionContext := context.Background()

	//	Concatenate the Azure region with the Azure base URL to create the endpoint URL.
	endpointURL := "https://" + azureRegion + ".api.cognitive.microsoft.com"

	//	Create an instance of the client with the endpoint URL.
	computerVisionClient := computervision.New(endpointURL)

	// Set up the authorization on the client with the subscription key.
	computerVisionClient.Authorizer = autorest.NewCognitiveServicesAuthorizer(azureComputerVisionAPIKey)

	//	BEGIN - Analyze a local image.
	//	Set the relative path to a local image.
	pathToLocalImage := "resources\\tech-writer.jpg"

	//	Print the path to the local image.
	workingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nLocal image path:\n%v\n", workingDirectory + "\\" + pathToLocalImage)

	//	Instantiate a ReadCloser required by AnalyzeImageInStream.
	var localImageFile io.ReadCloser

	//	Open the file for reading.
	localImageFile, err = os.Open(pathToLocalImage)
	if err != nil {
		log.Fatal(err)
	}

	//	Define what to extract frm the local image by initializing an array of VisualFeatureTypes.
	localImageVisualFeatureTypes := []computervision.VisualFeatureTypes{
		computervision.VisualFeatureTypesDescription,
		computervision.VisualFeatureTypesCategories,
		computervision.VisualFeatureTypesTags,
		computervision.VisualFeatureTypesFaces,
		computervision.VisualFeatureTypesAdult,
		computervision.VisualFeatureTypesColor,
		computervision.VisualFeatureTypesImageType,
	}

	//	Notify that we're beginning to analyze the local image.
	fmt.Println("\nAnalyzing local image ...")

	//  Call the Computer Vision service and tell it to analyze the local image.
	localImageAnalysis, err := computerVisionClient.AnalyzeImageInStream(
			computerVisionContext,
			localImageFile,
			localImageVisualFeatureTypes,
			[]computervision.Details{},
			"en")
		if err != nil {
			log.Fatal(err)
		}

	//  Display image captions and confidence values.
	fmt.Println("\nCaptions: ")
	for _, caption := range *localImageAnalysis.Description.Captions {
		fmt.Printf("'%v' with confidence %v\n", *caption.Text, *caption.Confidence)
	}

	//  Display image category names and confidence values.
	fmt.Println("\nCategories: ")
	for _, category := range *localImageAnalysis.Categories {
		fmt.Printf("'%v' with confidence %v\n", *category.Name, *category.Score)
	}

	//  Display image tags and confidence values.
	fmt.Println("\nTags: ")
	for _, tag := range *localImageAnalysis.Tags {
		fmt.Printf("'%v' with confidence %v\n", *tag.Name, *tag.Confidence)
	}

	//  Display any faces found in the image and their bounding boxes.
	fmt.Println("\nFaces: ")
	for _, face := range *localImageAnalysis.Faces {
		fmt.Printf("'%v' of age %v at location (%v, %v), (%v, %v)\n",
			face.Gender, *face.Age,
			*face.FaceRectangle.Left, *face.FaceRectangle.Top,
			*face.FaceRectangle.Left + *face.FaceRectangle.Width,
			*face.FaceRectangle.Top + *face.FaceRectangle.Height)
	}

	//  Display whether any adult or racy content was detected and the confidence values.
	fmt.Println("\nAdult: ");
	fmt.Printf("Is adult content: %v with confidence %f\n", *localImageAnalysis.Adult.IsAdultContent, *localImageAnalysis.Adult.AdultScore)
	fmt.Printf("Has racy content: %v with confidence %f\n", *localImageAnalysis.Adult.IsRacyContent, *localImageAnalysis.Adult.RacyScore)

	//  Display the image color scheme.
	fmt.Println("\nColor scheme: ");
	fmt.Printf("Is black and white: %v\n", *localImageAnalysis.Color.IsBWImg)
	fmt.Printf("Accent color: %v\n", *localImageAnalysis.Color.AccentColor)
	fmt.Printf("Dominant background color: %v\n", *localImageAnalysis.Color.DominantColorBackground)
	fmt.Printf("Dominant foreground color: %v\n", *localImageAnalysis.Color.DominantColorForeground)
	fmt.Printf("Dominant colors: %v\n", strings.Join(*localImageAnalysis.Color.DominantColors, ", "))

	//  Display any celebrities detected in the image and their bounding boxes.
	fmt.Println("\nCelebrities: ")
	for _, category := range *localImageAnalysis.Categories {
		if (category.Detail != nil && category.Detail.Celebrities != nil) {
			for _, celeb := range *category.Detail.Celebrities {
				fmt.Printf("'%v' with confidence %v at location (%v, %v), (%v, %v)\n",
					*celeb.Name, *celeb.Confidence,
					*celeb.FaceRectangle.Left, *celeb.FaceRectangle.Top,
					*celeb.FaceRectangle.Left + *celeb.FaceRectangle.Width,
					*celeb.FaceRectangle.Top + *celeb.FaceRectangle.Height)
			}
		}
	}

	//  Display any landmarks detected in the image and their bounding boxes.
	fmt.Println("\nLandmarks: ")
	for _, category := range *localImageAnalysis.Categories {
		if (category.Detail != nil && category.Detail.Landmarks != nil) {
			for _, landmark := range *category.Detail.Landmarks {
				fmt.Printf("'%v' with confidence %v\n", *landmark.Name, *landmark.Confidence)
			}
		}
	}

	//  Display what type of clip art or line drawing the image is.
	//	See the documentation for information about the meaning of the return values.
	fmt.Println("\nImage type:")
	fmt.Printf("Clip art type: %v\n", *localImageAnalysis.ImageType.ClipArtType)
	fmt.Printf("Line drawing type: %v\n", *localImageAnalysis.ImageType.LineDrawingType)
	//  END - Analyze a local image.


	 //	BEGIN - Analyze an image from a URL.
	 //	Set a string variable equal to the path of a remote image.
	pathToRemoteImage := "https://github.com/Azure-Samples/cognitive-services-sample-data-files/raw/master/ComputerVision/Images/faces.jpg"

	//	Need ImageURL type to pass to AnalyzeImage.
	var imageURL computervision.ImageURL
	imageURL.URL = &pathToRemoteImage

	//	Print the image URL.
	fmt.Printf("\n\nImage URL:\n%v\n", pathToRemoteImage)

	 //	Set up an array of VisualFeatureTypes, which defines what to extract from the image.
 	remoteImageVisualFeatureTypes := []computervision.VisualFeatureTypes{
 		computervision.VisualFeatureTypesDescription,
 		computervision.VisualFeatureTypesCategories,
 		computervision.VisualFeatureTypesTags,
 		computervision.VisualFeatureTypesFaces,
 		computervision.VisualFeatureTypesAdult,
 		computervision.VisualFeatureTypesColor,
 		computervision.VisualFeatureTypesImageType,
 	}

	//	Notify that we're beginning to analyze the local image.
	fmt.Println("\nAnalyzing an image from a URL ...")

	//  Call the Computer Vision service and tell it to analyze the remote image. (Ignoring any errors.)
	remoteImageAnalysis, err := computerVisionClient.AnalyzeImage(
		computerVisionContext,
		imageURL,
		remoteImageVisualFeatureTypes,
		[]computervision.Details{},
		"en")
	if err != nil {
	 	log.Fatal(err)
	}

	//  Display image captions and confidence values.
	fmt.Println("\nCaptions: ")
	for _, caption := range *remoteImageAnalysis.Description.Captions {
		fmt.Printf("'%v' with confidence %v\n", *caption.Text, *caption.Confidence)
	}

	//  Display image category names and confidence values.
	fmt.Println("\nCategories: ")
	for _, category := range *remoteImageAnalysis.Categories {
		fmt.Printf("'%v' with confidence %v\n", *category.Name, *category.Score)
	}

	//  Display image tags and confidence values.
	fmt.Println("\nTags: ")
	for _, tag := range *remoteImageAnalysis.Tags {
		fmt.Printf("'%v' with confidence %v\n", *tag.Name, *tag.Confidence)
	}

	//  Display any faces found in the image and their bounding boxes.
	fmt.Println("\nFaces: ")
	for _, face := range *remoteImageAnalysis.Faces {
		fmt.Printf("'%v' of age %v at location (%v, %v), (%v, %v)\n",
			face.Gender, *face.Age,
			*face.FaceRectangle.Left, *face.FaceRectangle.Top,
			*face.FaceRectangle.Left + *face.FaceRectangle.Width,
			*face.FaceRectangle.Top + *face.FaceRectangle.Height)
	}

	//  Display whether any adult or racy content was detected and the confidence values.
	fmt.Println("\nAdult: ");
	fmt.Printf("Is adult content: %v with confidence %f\n", *remoteImageAnalysis.Adult.IsAdultContent, *remoteImageAnalysis.Adult.AdultScore)
	fmt.Printf("Has racy content: %v with confidence %f\n", *remoteImageAnalysis.Adult.IsRacyContent, *remoteImageAnalysis.Adult.RacyScore)

	//  Display the image color scheme.
	fmt.Println("\nColor scheme: ");
	fmt.Printf("Is black and white: %v\n", *remoteImageAnalysis.Color.IsBWImg)
	fmt.Printf("Accent color: %v\n", *remoteImageAnalysis.Color.AccentColor)
	fmt.Printf("Dominant background color: %v\n", *remoteImageAnalysis.Color.DominantColorBackground)
	fmt.Printf("Dominant foreground color: %v\n", *remoteImageAnalysis.Color.DominantColorForeground)
	fmt.Printf("Dominant colors: %v\n", strings.Join(*remoteImageAnalysis.Color.DominantColors, ", "))

	//  Display any celebrities detected in the image and their bounding boxes.
	fmt.Println("\nCelebrities: ")
	for _, category := range *remoteImageAnalysis.Categories {
		if (category.Detail != nil && category.Detail.Celebrities != nil) {
			for _, celeb := range *category.Detail.Celebrities {
				fmt.Printf("'%v' with confidence %v at location (%v, %v), (%v, %v)\n",
					*celeb.Name, *celeb.Confidence,
					*celeb.FaceRectangle.Left, *celeb.FaceRectangle.Top,
					*celeb.FaceRectangle.Left + *celeb.FaceRectangle.Width,
					*celeb.FaceRectangle.Top + *celeb.FaceRectangle.Height)
			}
		}
	}

	//  Display any landmarks detected in the image and their bounding boxes.
	fmt.Println("\nLandmarks: ")
	for _, category := range *remoteImageAnalysis.Categories {
		if (category.Detail != nil && category.Detail.Landmarks != nil) {
			for _, landmark := range *category.Detail.Landmarks {
				fmt.Printf("'%v' with confidence %v\n", *landmark.Name, *landmark.Confidence)
			}
		}
	}

	//  Display what type of clip art or line drawing the image is.
	//	See the documentation for information about the meaning of the return values.
	fmt.Println("\nImage type:")
	fmt.Printf("Clip art type: %v\n", *remoteImageAnalysis.ImageType.ClipArtType)
	fmt.Printf("Line drawing type: %v\n", *remoteImageAnalysis.ImageType.LineDrawingType)
	//	END - Analyze an image from a URL.
}
