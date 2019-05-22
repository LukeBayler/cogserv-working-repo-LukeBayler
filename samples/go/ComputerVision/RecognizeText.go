package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.0/computervision"
	"github.com/Azure/go-autorest/autorest"
)

/*  This Quickstart for the Azure Cognitive Services Computer Vision API shows
 *  you how to recognize both handwritten and printed text using both
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
	if "" == azureComputerVisionAPIKey {
		log.Fatal("Please set the AZURE_COMPUTERVISION_API_KEY environment variable. Note that you might need to restart your shell or IDE.")
	}

	azureRegion := os.Getenv("AZURE_REGION")
	if "" == azureRegion {
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

	//	BEGIN - Recognize printed text from a local image.
	//	Set the relative path to a local image.
	pathToLocalImage := "resources\\printed_text.jpg"

	//	Print the path to the local image.
	workingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nLocal image path:\n%v\n", workingDirectory+"\\"+pathToLocalImage)

	//	Instantiate a ReadCloser required by RecognizePrintedTextInStream.
	var localImageFile io.ReadCloser

	//	Open the file for reading.
	localImageFile, err = os.Open(pathToLocalImage)
	if err != nil {
		log.Fatal(err)
	}

	//	Notify that we're beginning to recognize printed text from the remote image.
	fmt.Println("\nRecognizing printed text from a local image ...\n")

	//	Set the text recognition mode to printed.
	textRecognitionMode := computervision.Printed

	//	When you use the Read Document interface, the response contains a field
	//	called "Operation-Location", which contains the URL to use for your
	//	Get Read Result operation" to access OCR results.
	textHeaders, err := computerVisionClient.BatchReadFileInStream(
		computerVisionContext,
		localImageFile,
		textRecognitionMode)
	if err != nil {
		log.Fatal(err)
	}

	//	Use ExtractHeader from the autorest library to get the Operation-Location URL
	operationLocation := autorest.ExtractHeaderValue("Operation-Location", textHeaders.Response)

	numberOfCharsInOperationId := 36
	operationId := string(operationLocation[len(operationLocation)-numberOfCharsInOperationId : len(operationLocation)])

	readOperationResult, err := computerVisionClient.GetReadOperationResult(computerVisionContext, operationId)
	if err != nil {
		log.Fatal(err)
	}

	// Wait for the operation to complete.
	i := 0
	maxRetries := 10

	for readOperationResult.Status != computervision.Failed &&
			readOperationResult.Status != computervision.Succeeded {
		if i >= maxRetries {
			break
		}
		i++

		fmt.Printf("Server status: %v, waiting %v seconds...\n", readOperationResult.Status, i)
		time.Sleep(1 * time.Second)

		readOperationResult, err = computerVisionClient.GetReadOperationResult(computerVisionContext, operationId)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Display the results.
	for _, recResult := range *(readOperationResult.RecognitionResults) {
		for _, line := range *recResult.Lines {
			fmt.Println(*line.Text)
		}
	}
	//	END - Recognize printed text from a local image.

	//	BEGIN - Recognize handwritten text from a remote image.
		pathToRemoteImage := "https://raw.githubusercontent.com/Azure-Samples/cognitive-services-sample-data-files/master/ComputerVision/Images/handwritten_text.jpg"
		fmt.Printf("\n\nRemote image path:\n%v\n", pathToRemoteImage)

		//	Need ImageURL type to pass to AnalyzeImage.
		var imageURL computervision.ImageURL
		imageURL.URL = &pathToRemoteImage

	//	Notify that we're beginning to recognize handwritten text from the local image.
	fmt.Println("\nRecognizing handwritten text from a remote image ...\n")

	//	Set the text recognition mode to handwritten.
	textRecognitionModeRemoteImage := computervision.Handwritten

	//	When you use the Read Document interface, the response contains a field
	//	called "Operation-Location", which contains the URL to use for your
	//	Get Read Result operation" to access OCR results.
	textHeadersRemoteImage, err := computerVisionClient.BatchReadFile(
		computerVisionContext,
		imageURL,
		textRecognitionModeRemoteImage)
	if err != nil {
		log.Fatal(err)
	}

	//	Use ExtractHeader from the autorest library to get the Operation-Location URL
	operationLocationRemoteImage := autorest.ExtractHeaderValue("Operation-Location", textHeadersRemoteImage.Response)

	numberOfCharsInOperationIdRemoteImage := 36
	operationIdRemoteImage := string(operationLocationRemoteImage[len(operationLocationRemoteImage) - numberOfCharsInOperationIdRemoteImage : len(operationLocationRemoteImage)])

	readOperationResultRemoteImage, err := computerVisionClient.GetReadOperationResult(computerVisionContext, operationIdRemoteImage)
	if err != nil {
		log.Fatal(err)
	}

	// Wait for the operation to complete.
	iRemoteImage := 0
	maxRetriesRemoteImage := 10

	for readOperationResultRemoteImage.Status != computervision.Failed &&
			readOperationResultRemoteImage.Status != computervision.Succeeded {
				if iRemoteImage >= maxRetriesRemoteImage {
					fmt.Printf("iRemoteImage: %v\n", iRemoteImage)
					break
				}
				iRemoteImage++

				fmt.Printf("Server status: %v, waiting %v seconds...\n", readOperationResultRemoteImage.Status, iRemoteImage)
				time.Sleep(1 * time.Second)

				readOperationResultRemoteImage, err = computerVisionClient.GetReadOperationResult(computerVisionContext, operationIdRemoteImage)
				if err != nil {
					log.Fatal(err)
				}
			}

		// Display the results.
		fmt.Println()
		for _, recResultRemoteImage := range *readOperationResultRemoteImage.RecognitionResults {
				for _, lineRemoteImage := range *recResultRemoteImage.Lines {
					fmt.Println(*lineRemoteImage.Text)
				}
			}
	//	END - Recognize handwritten text from a remote image.
}
