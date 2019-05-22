package main

import (
	"context"
	"fmt"
  "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.0/computervision"
	"github.com/Azure/go-autorest/autorest"
	"io"
	"log"
	"os"
	// "strings"
)

/*  This Quickstart for the Azure Cognitive Services Computer Vision API shows
 *  you how to detect ojbects in a local and a remote image.
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
	pathToLocalImage := "resources\\horse-people.jpg"

	//	Print the path to the local image.
	workingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nLocal image path:\n%v\n", workingDirectory + "\\" + pathToLocalImage)

	//	Instantiate a ReadCloser required by DetectOjbectsInStream.
	var localImageFile io.ReadCloser

	//	Open the file for reading.
	localImageFile, err = os.Open(pathToLocalImage)
	if err != nil {
		log.Fatal(err)
	}

	//	Notify that we're beginning to analyze the local image.
	fmt.Println("\nAnalyzing local image ...")

	//  Call the Computer Vision service and tell it to analyze the local image.
	localImageAnalysis, err := computerVisionClient.DetectObjectsInStream(
			computerVisionContext,
			localImageFile,
			)
		if err != nil {
			log.Fatal(err)
		}

	//  Display image captions and confidence values.
	fmt.Println("\nObjects: ")
	for _, object := range *localImageAnalysis.Objects {
		fmt.Printf("'%v' with confidence %v at location (%v, %v), (%v, %v)\n", *object.Object, *object.Confidence,
      *object.Rectangle.X, *object.Rectangle.X + *object.Rectangle.W,
      *object.Rectangle.Y, *object.Rectangle.Y + *object.Rectangle.H)
	}
	//  END - Analyze a local image.

	 //	BEGIN - Analyze an image from a URL.
	 //	Set a string variable equal to the path of a remote image.
	pathToRemoteImage := "https://github.com/Azure-Samples/cognitive-services-sample-data-files/raw/master/ComputerVision/Images/faces.jpg"

	//	Need ImageURL type to pass to AnalyzeImage.
	var imageURL computervision.ImageURL
	imageURL.URL = &pathToRemoteImage

	//	Print the image URL.
	fmt.Printf("\n\nImage URL:\n%v\n", pathToRemoteImage)

	//	Notify that we're beginning to analyze the local image.
	fmt.Println("\nAnalyzing an image from a URL ...")

	//  Call the Computer Vision service and tell it to analyze the remote image. (Ignoring any errors.)
	remoteImageAnalysis, err := computerVisionClient.DetectObjects(
		computerVisionContext,
		imageURL,
  )
	if err != nil {
	 	log.Fatal(err)
	}

  //  Display image captions and confidence values.
	fmt.Println("\nObjects: ")
	for _, object := range *remoteImageAnalysis.Objects {
		fmt.Printf("'%v' with confidence %v at location (%v, %v), (%v, %v)\n", *object.Object, *object.Confidence,
      *object.Rectangle.X, *object.Rectangle.X + *object.Rectangle.W,
      *object.Rectangle.Y, *object.Rectangle.Y + *object.Rectangle.H)
	}
	//	END - Analyze an image from a URL.
}
