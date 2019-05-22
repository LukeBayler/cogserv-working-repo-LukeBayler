package main

import (
	"context"
	"fmt"
  "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.0/computervision"
	"github.com/Azure/go-autorest/autorest"
	"io"
	"log"
	"os"
)

/*  This Quickstart for the Azure Cognitive Services Computer Vision API shows
 *  you how to recognize both handwritten and printed text with OCR using both
 * a local and a remote image.
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

  //  BEGIN - Configure and authenticate the Computer Vision client.
	// Get the context, which is required by the SDK methods.
	computerVisionContext := context.Background()

	//	Concatenate the Azure region with the Azure base URL to create the endpoint URL.
	endpointURL := "https://" + azureRegion + ".api.cognitive.microsoft.com"

	//	Create an instance of the client with the endpoint URL.
	computerVisionClient := computervision.New(endpointURL)

	// Set up the authorization on the client with the subscription key.
	computerVisionClient.Authorizer = autorest.NewCognitiveServicesAuthorizer(azureComputerVisionAPIKey)
  //  END - Configure and authenticate the Computer Vision client.

	//	BEGIN - Recognize handwritten text from a local image with OCR.
	//	Set the relative path to a local image.
	pathToLocalImage := "resources\\handwritten_text.jpg"

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

	//	Notify that we're beginning to recognize handwritten text from the local image.
	fmt.Println("Recognizing handwritten text from a local image ...\n")

	//  Call the Computer Vision service and tell it to recognize handwritten text.
	localImageOcrResult, err := computerVisionClient.RecognizePrintedTextInStream(
			computerVisionContext,
      true,
      localImageFile,
			"en")
		if err != nil {
			log.Fatal(err)
		}

  //  Display the results.
  fmt.Printf("Text:\n")
  fmt.Printf("Language: %v\n", *localImageOcrResult.Language)
  fmt.Printf("Text angle: %v\n", *localImageOcrResult.TextAngle)
  fmt.Printf("Orientation: %v\n", *localImageOcrResult.Orientation)

  fmt.Printf("Text regions:\n")
  for _, localImageOcrRegion := range *localImageOcrResult.Regions {
    fmt.Printf("\tRegion bounding box: %v\n", *localImageOcrRegion.BoundingBox)
    for _, localImageOcrLine := range *localImageOcrRegion.Lines {
      fmt.Printf("\tLine bounding box %v\n", *localImageOcrLine.BoundingBox)
      for _, localImageOcrWord := range *localImageOcrLine.Words {
        fmt.Printf("\t\tWord bounding box: %v\n", *localImageOcrWord.BoundingBox)
        fmt.Printf("\t\tText: %v\n\n", *localImageOcrWord.Text)
      }
      fmt.Println()
    }
    fmt.Println()
  }
  //	END - Recognize handwritten text from a local image with OCR.


	 //	BEGIN - Recognize printed text from a remote image with OCR.
	 //	Set a string variable equal to the path of a remote image.
	pathToRemoteImage := "https://github.com/Azure-Samples/cognitive-services-sample-data-files/raw/master/ComputerVision/Images/printed_text.jpg"

	//	Need ImageURL type to pass to AnalyzeImage.
	var imageURL computervision.ImageURL
	imageURL.URL = &pathToRemoteImage

	//	Print the image URL.
	fmt.Printf("\n\nImage URL:\n%v\n", pathToRemoteImage)

	//	Notify that we're beginning to recognize printed text from the remote image.
	fmt.Println("\nRecognizing printed text from a remote image ...")

	//  Call the Computer Vision service and tell it to analyze the remote image.
	remoteImageOcrResult, err := computerVisionClient.RecognizePrintedText(
		computerVisionContext,
    true,
    imageURL,
		"en")
	if err != nil {
	 	log.Fatal(err)
	}

  //  Display the results.
  fmt.Printf("Text:\n")
  fmt.Printf("Language: %v\n", *remoteImageOcrResult.Language)
  fmt.Printf("Text angle: %v\n", *remoteImageOcrResult.TextAngle)
  fmt.Printf("Orientation: %v\n", *remoteImageOcrResult.Orientation)

  fmt.Printf("Text regions:\n")
  for _, remoteImageOcrRegion := range *remoteImageOcrResult.Regions {
    fmt.Printf("\tRegion bounding box: %v\n", *remoteImageOcrRegion.BoundingBox)
    for _, remoteImageOcrLine := range *remoteImageOcrRegion.Lines {
      fmt.Printf("\tLine bounding box %v\n", *remoteImageOcrLine.BoundingBox)
      for _, remoteImageOcrWord := range *remoteImageOcrLine.Words {
        fmt.Printf("\t\tWord bounding box: %v\n", *remoteImageOcrWord.BoundingBox)
        fmt.Printf("\t\tText: %v\n\n", *remoteImageOcrWord.Text)
      }
      fmt.Println()
    }
    fmt.Println()
  }
  //	END - Recognize printed text from a remote image with OCR.
}
