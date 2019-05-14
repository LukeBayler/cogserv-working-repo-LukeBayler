package main

import (
	"context"
	"fmt"
  "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.0/computervision"
	"github.com/Azure/go-autorest/autorest"
	"io"
	// "io/ioutil"
	"log"
	"os"
)

func main() {

	fmt.Println("\nAzure Cognitive Services Computer Vision - Go Quickstart Sample")

	azureComputerVisionAPIKey := os.Getenv("AZURE_COMPUTERVISION_API_KEY")
	azureRegion := os.Getenv("AZURE_REGION")
	//fmt.Println(azureComputerVisionAPIKey)
	//fmt.Println(azureRegion)

	endpointURL := "https://" + azureRegion + ".api.cognitive.microsoft.com"
	//fmt.Println(endpointURL)

	// Get the context, which is required by the SDK methods.
	ctx := context.Background()

	//	Creates an instance of the BaseClient client.
	cvClient := computervision.New(endpointURL)

	// Set the subscription key on the client.
	cvClient.Authorizer = autorest.NewCognitiveServicesAuthorizer(azureComputerVisionAPIKey)

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("pwd: " + pwd)

	//	Set the relative path to a local image.
	pathToLocalImage := "resources\\landmark.jpg"

	//	Set up an array of VisualFeatureTypes, which defines what to extract from the image.
	visualFeatureTypes := []computervision.VisualFeatureTypes{
		computervision.VisualFeatureTypesDescription,
		computervision.VisualFeatureTypesCategories,
		computervision.VisualFeatureTypesTags,
		computervision.VisualFeatureTypesFaces,
		computervision.VisualFeatureTypesAdult,
		computervision.VisualFeatureTypesColor,
		computervision.VisualFeatureTypesImageType,
	}

	//	Loop through the array to make sure we've initialized it.
	for i := 0; i < len(visualFeatureTypes); i++ {
		fmt.Println(visualFeatureTypes[i])
	}

	fmt.Println("\nAnalyzing local image ...")

	//	Get the image data.
	var imgFile io.ReadCloser
	imgFile, err = os.Open(pathToLocalImage)
	// imgFile, err := os.Open(pathToLocalImage)

	// imageData, err := ioutil.ReadAll(imgFile)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	details := []computervision.Details{}

	//  Call the Computer Vision service and tell it to analyze the loaded image.
	imageAnalysis, err := cvClient.AnalyzeImageInStream(ctx, imgFile, visualFeatureTypes, details, "en")

	fmt.Println("\nCaptions: ")
	for _, caption := range *imageAnalysis.Description.Captions {
		fmt.Println(*caption.Text)
	}
}
