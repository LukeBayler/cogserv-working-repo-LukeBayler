package main

/*  Import the required libraries. If this is your first time running a Go program,
 *  you will need to 'go get' the azure-sdk-for-go and go-autorest packages.
 */
import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.0/computervision"
	"github.com/Azure/go-autorest/autorest"
	"io"
	"log"
	"os"
	"strings"
)

/*  The Quickstarts in this file are for the Computer Vision API for Microsoft
 *  Cognitive Services. In this file are Quickstarts for the following tasks:
 *  - Describing images
 *  - Categorizing images
 *  - Tagging images
 *  - Detecting faces
 *  - Detecting adult or racy content
 *  - Detecting the color scheme
 *  - Detecting domain-specific content (celebrities/landmarks)
 *  - Detecting image types (clip art/line drawing)
 *  - Recognizing printed and handwritten text with (Read API)
 *  - Recognizing printed and handwritten text with (OCR)
 *  - Detecting objects
 *  - Detecting brands
 */

//	Declare global so don't have to pass it to all of the tasks.
var computerVisionContext context.Context

func main() {
	/*	Configure the Computer Vision client by:
	 *    1. Reading the Computer Vision API key and the Azure region from environment
	 *       variables (COMPUTERVISION_API_KEY and COMPUTERVISION_REGION), which must
	 *       be set prior to running this code. After setting the	environment variables,
	 *       restart your command shell or your IDE.
	 *	  2. Constructing the endpoint URL from the base URL and the Azure region.
	 *	  3. Setting up the authorization on the client with the subscription key.
	 *	  4. Getting the context.
	 */
  computerVisionAPIKey := os.Getenv("COMPUTERVISION_API_KEY")
	if ("" == computerVisionAPIKey) {
		log.Fatal("\n\nPlease set the COMPUTERVISION_API_KEY environment variable.\n" +
							  "**Note that you might need to restart your shell or IDE.**\n")
	}

	computerVisionRegion := os.Getenv("COMPUTERVISION_REGION")
	if ("" == computerVisionRegion) {
		log.Fatal("\n\nPlease set the COMPUTERVISION_REGION environment variable.\n" +
							  "**Note that you might need to restart your shell or IDE.**")
	}

	endpointURL := "https://" + computerVisionRegion + ".api.cognitive.microsoft.com"

	computerVisionClient := computervision.New(endpointURL);
	computerVisionClient.Authorizer = autorest.NewCognitiveServicesAuthorizer(computerVisionAPIKey)

	computerVisionContext = context.Background()
	//	END - Configure the Computer Vision client

	//	Specify the relative path for a local image
	localImagePath := "resources\\landmark.jpg"
	workingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nLocal image path:\n%v\n", workingDirectory + "\\" + localImagePath)

	//	Call each of the Computer Vision tasks on the local image
	// DescribeLocalImage(computerVisionClient, localImagePath)
	// CategorizeLocalImage(computerVisionClient, localImagePath)
	// TagLocalImage(computerVisionClient, localImagePath)
	// DetectFacesLocalImage(computerVisionClient, localImagePath)
	// DetectAdultOrRacyContentLocalImage(computerVisionClient, localImagePath)
	// DetectColorSchemeLocalImage(computerVisionClient, localImagePath)
	//DetectDomainSpecificContentLocalImage(computerVisionClient, localImagePath)

	DetectImageTypesLocalImage(computerVisionClient, localImagePath)
	// RecognizeTextReadAPILocalImage(computerVisionClient, localImagePath)
	// RecognizeTextOCRLocalImage(computerVisionClient, localImagePath)
	// DetectObjectsLocalImage(computerVisionClient, localImagePath)
	// DetectBrandsLocalImage(computerVisionClient, localImagePath)


	//	Specify a URL for a remote image
	remoteImageURL := "https://github.com/Azure-Samples/cognitive-services-sample-data-files/raw/master/ComputerVision/Images/faces.jpg"
	fmt.Printf("\n\nRemote image path: \n%v\n", remoteImageURL)

	//	Call each of the Computer Vision tasks on the remote image
	// DescribeRemoteImage(computerVisionClient, remoteImageURL)
	// CategorizeRemoteImage(computerVisionClient, remoteImageURL)
	// TagRemoteImage(computerVisionClient, remoteImageURL)
	// DetectFacesRemoteImage(computerVisionClient, remoteImageURL)
	// DetectAdultOrRacyContentRemoteImage(computerVisionClient, remoteImageURL)
	// DetectColorSchemeRemoteImage(computerVisionClient, remoteImageURL)
	// DetectDomainSpecificContentRemoteImage(computerVisionClient, remoteImageURL)

	DetectImageTypesRemoteImage(computerVisionClient, remoteImageURL)
	// RecognizeTextReadAPIRemoteImage(computerVisionClient, remoteImageURL)
	// RecognizeTextOCRRemoteImage(computerVisionClient, remoteImageURL)
	// DetectObjectsRemoteImage(computerVisionClient, remoteImageURL)
	// DetectBrandsRemoteImage(computerVisionClient, remoteImageURL)
}

/*  Describe a local image by:
 *    1. Instantiating a ReadCloser, which is required by AnalyzeImageInStream.
 *    2. Opening the ReadCloser instance for reading.
 *    3. Defining what to extract from the image by initializing an array of VisualFeatureTypes.
 *    4. Calling the Computer Vision service with the:
 *       - context
 *       - image
 *       - features to extract
 *       - an empty slice for the Details enumeration
 *       - "" to specify the default language ("en") as the output language
 *    5. Displaying the image captions and their confidence values.
 */
func DescribeLocalImage(client computervision.BaseClient, localImagePath string) {
	var localImage io.ReadCloser
	localImage, err := os.Open(localImagePath)
	if err != nil {
		log.Fatal(err)
	}

	features := []computervision.VisualFeatureTypes{computervision.VisualFeatureTypesDescription}

	imageAnalysis, err := client.AnalyzeImageInStream(
			computerVisionContext,
			localImage,
			features,
			[]computervision.Details{},
			"")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\nCaptions from local image: ")
		if len(*imageAnalysis.Description.Captions) == 0 {
			fmt.Println("No captions detected.")
		}
		for _, caption := range *imageAnalysis.Description.Captions {
			fmt.Printf("'%v' with confidence %v\n", *caption.Text, *caption.Confidence)
		}

		localImage.Close()
}
//	END - Describe a local image

/*  Describe a remote image file by:
*    1. Saving the URL as an ImageURL type for passing to AnalyzeImage.
*    2. Defining what to extract from the image by initializing an array of VisualFeatureTypes.
*    3. Calling the Computer Vision service with the:
*       - context
*       - image
*       - features to extract
*       - an empty slice for the Details enumeration
*       - "" to specify the default language ("en") as the output language
*    4. Displaying the image captions and their confidence values.
 */
func DescribeRemoteImage(client computervision.BaseClient, remoteImageURL string) {
	var remoteImage computervision.ImageURL
	remoteImage.URL = &remoteImageURL

	features := []computervision.VisualFeatureTypes{computervision.VisualFeatureTypesDescription}
	imageAnalysis, err := client.AnalyzeImage(
		computerVisionContext,
		remoteImage,
		features,
		[]computervision.Details{},
		"")
	if err != nil {
	 	log.Fatal(err)
	}

	fmt.Println("\nCaptions from remote image: ")
	if len(*imageAnalysis.Description.Captions) == 0 {
		fmt.Println("No captions detected.")
	}
	for _, caption := range *imageAnalysis.Description.Captions {
		fmt.Printf("'%v' with confidence %v\n", *caption.Text, *caption.Confidence)
	}
}
//	END - Describe a remote image

/*  Categorize a local image by:
 *    1. Instantiating a ReadCloser, which is required by AnalyzeImageInStream.
 *    2. Opening the ReadCloser instance for reading.
 *    3. Defining what to extract from the image by initializing an array of VisualFeatureTypes.
 *    4. Calling the Computer Vision service with the:
 *       - context
 *       - image
 *       - features to extract
 *       - an empty slice for the Details enumeration
 *       - "" to specify the default language ("en") as the output language
 *    5. Displaying the image categories and their confidence values.
 */
func CategorizeLocalImage(client computervision.BaseClient, localImagePath string) {
	var localImage io.ReadCloser
	localImage, err := os.Open(localImagePath)
	if err != nil {
		log.Fatal(err)
	}

	features := []computervision.VisualFeatureTypes{computervision.VisualFeatureTypesCategories}
	imageAnalysis, err := client.AnalyzeImageInStream(
			computerVisionContext,
			localImage,
			features,
			[]computervision.Details{},
			"")
		if err != nil {
			log.Fatal(err)
		}

	fmt.Println("\nCategories from local image: ")
	if len(*imageAnalysis.Categories) == 0 {
		fmt.Println("No categories detected.")
	}
	for _, category := range *imageAnalysis.Categories {
		fmt.Printf("'%v' with confidence %v\n", *category.Name, *category.Score)
	}

	localImage.Close()
}
//	END - Categorize a local image


/*  Categorize a remote image by:
*    1. Saving the URL as an ImageURL type for passing to AnalyzeImage.
*    2. Defining what to extract from the image by initializing an array of VisualFeatureTypes.
*    3. Calling the Computer Vision service with the:
*       - context
*       - image
*       - features to extract
*       - an empty slice for the Details enumeration
*       - "" to specify the default language ("en") as the output language
*    4. Displaying the image categories and their confidence values.
 */
func CategorizeRemoteImage(client computervision.BaseClient, remoteImageURL string) {
	var remoteImage computervision.ImageURL
	remoteImage.URL = &remoteImageURL

	features := []computervision.VisualFeatureTypes{computervision.VisualFeatureTypesCategories}
	imageAnalysis, err := client.AnalyzeImage(
			computerVisionContext,
			remoteImage,
			features,
			[]computervision.Details{},
			"")
		if err != nil {
			log.Fatal(err)
		}

	fmt.Println("\nCategories from local image: ")
	if len(*imageAnalysis.Categories) == 0 {
		fmt.Println("No categories detected.")
	}
	for _, category := range *imageAnalysis.Categories {
		fmt.Printf("'%v' with confidence %v\n", *category.Name, *category.Score)
	}
}
//	END - Categorize a remote image


/*  Tag a local image by:
 *    1. Instantiating a ReadCloser, which is required by AnalyzeImageInStream.
 *    2. Opening the ReadCloser instance for reading.
 *    3. Defining what to extract from the image by initializing an array of VisualFeatureTypes.
 *    4. Calling the Computer Vision service with the:
 *       - context
 *       - image
 *       - features to extract
 *       - an empty slice for the Details enumeration
 *       - "" to specify the default language ("en") as the output language
 *    5. Displaying the image categories and their confidence values.
 */
func TagLocalImage(client computervision.BaseClient, localImagePath string) {
	var localImage io.ReadCloser
	localImage, err := os.Open(localImagePath)
	if err != nil {
		log.Fatal(err)
	}

	features := []computervision.VisualFeatureTypes{computervision.VisualFeatureTypesTags}
	imageAnalysis, err := client.AnalyzeImageInStream(
			computerVisionContext,
			localImage,
			features,
			[]computervision.Details{},
			"")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\nTags in the local image: ")
		if len(*imageAnalysis.Tags) == 0 {
			fmt.Println("No tags detected.")
		}
		for _, tag := range *imageAnalysis.Tags {
			fmt.Printf("'%v' with confidence %v\n", *tag.Name, *tag.Confidence)
		}

		localImage.Close()
	}
	//	END - Tag a local image


	/*  Tag a remote image file by:
	*    1. Saving the URL as an ImageURL type for passing to AnalyzeImage.
	*    2. Defining what to extract from the image by initializing an array of VisualFeatureTypes.
	*    3. Calling the Computer Vision service with the:
	*       - context
	*       - image
	*       - features to extract
	*       - an empty slice for the Details enumeration
	*       - "" to specify the default language ("en") as the output language
	*    4. Displaying the image categories and their confidence values.
	 */
func TagRemoteImage(client computervision.BaseClient, remoteImageURL string) {
	var remoteImage computervision.ImageURL
	remoteImage.URL = &remoteImageURL

	features := []computervision.VisualFeatureTypes{computervision.VisualFeatureTypesTags}
	imageAnalysis, err := client.AnalyzeImage(
			computerVisionContext,
			remoteImage,
			features,
			[]computervision.Details{},
			"")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\nTags in the remote image: ")
		if len(*imageAnalysis.Tags) == 0 {
			fmt.Println("No tags detected.")
		}
		for _, tag := range *imageAnalysis.Tags {
			fmt.Printf("'%v' with confidence %v\n", *tag.Name, *tag.Confidence)
		}
}
//	END - Tag a remote image


/*  Detect faces in a local image by:
 *    1. Instantiating a ReadCloser, which is required by AnalyzeImageInStream.
 *    2. Opening the ReadCloser instance for reading.
 *    3. Defining what to extract from the image by initializing an array of VisualFeatureTypes.
 *    4. Calling the Computer Vision service with the:
 *       - context
 *       - image
 *       - features to extract
 *       - an empty slice for the Details enumeration
 *       - "" to specify the default language ("en") as the output language
 *    5. Displaying the faces and their bounding boxes.
 */
func DetectFacesLocalImage(client computervision.BaseClient, localImagePath string) {
	var localImage io.ReadCloser
	localImage, err := os.Open(localImagePath)
	if err != nil {
		log.Fatal(err)
	}

	features := []computervision.VisualFeatureTypes{computervision.VisualFeatureTypesFaces}

	imageAnalysis, err := client.AnalyzeImageInStream(
			computerVisionContext,
			localImage,
			features,
			[]computervision.Details{},
			"")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\nDetecting faces in a local image ...")
		if len(*imageAnalysis.Faces) == 0 {
			fmt.Println("No faces detected.")
		}
		for _, face := range *imageAnalysis.Faces {
			fmt.Printf("'%v' of age %v at location (%v, %v), (%v, %v)\n",
				face.Gender, *face.Age,
				*face.FaceRectangle.Left, *face.FaceRectangle.Top,
				*face.FaceRectangle.Left + *face.FaceRectangle.Width,
				*face.FaceRectangle.Top + *face.FaceRectangle.Height)
		}
}
//	END - Detect faces in a local image


/*  Detect faces in a remote image file by:
*    1. Saving the URL as an ImageURL type for passing to AnalyzeImage.
*    2. Defining what to extract from the image by initializing an array of VisualFeatureTypes.
*    3. Calling the Computer Vision service with the:
*       - context
*       - image
*       - features to extract
*       - an empty slice for the Details enumeration
*       - "" to specify the default language ("en") as the output language
*    4. Displaying the image categories and their confidence values.
 */
func DetectFacesRemoteImage(client computervision.BaseClient, remoteImageURL string) {
	var remoteImage computervision.ImageURL
	remoteImage.URL = &remoteImageURL

	features := []computervision.VisualFeatureTypes{computervision.VisualFeatureTypesFaces}
	imageAnalysis, err := client.AnalyzeImage(
			computerVisionContext,
			remoteImage,
			features,
			[]computervision.Details{},
			"")
		if err != nil {
			log.Fatal(err)
		}

	fmt.Println("\nDetecting faces in a remote image ...")
	if len(*imageAnalysis.Faces) == 0 {
		fmt.Println("No faces detected.")
	}
	for _, face := range *imageAnalysis.Faces {
		fmt.Printf("'%v' of age %v at location (%v, %v), (%v, %v)\n",
			face.Gender, *face.Age,
			*face.FaceRectangle.Left, *face.FaceRectangle.Top,
			*face.FaceRectangle.Left + *face.FaceRectangle.Width,
			*face.FaceRectangle.Top + *face.FaceRectangle.Height)
	}
}
//	END - Detect faces in a remote image


/*  Detect adult or racy content in a local image by:
 *    1. Instantiating a ReadCloser, which is required by AnalyzeImageInStream.
 *    2. Opening the ReadCloser instance for reading.
 *    3. Defining what to extract from the image by initializing an array of VisualFeatureTypes.
 *    4. Calling the Computer Vision service with the:
 *       - context
 *       - image
 *       - features to extract
 *       - an empty slice for the Details enumeration
 *       - "" to specify the default language ("en") as the output language
 *    5. Displaying the faces and their bounding boxes.
 */
func DetectAdultOrRacyContentLocalImage(client computervision.BaseClient, localImagePath string) {
	var localImage io.ReadCloser
	localImage, err := os.Open(localImagePath)
	if err != nil {
		log.Fatal(err)
	}

	features := []computervision.VisualFeatureTypes{computervision.VisualFeatureTypesAdult}
	imageAnalysis, err := client.AnalyzeImageInStream(
			computerVisionContext,
			localImage,
			features,
			[]computervision.Details{},
			"")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\nAnalyzing local image for adult or racy content: ");
		fmt.Printf("Is adult content: %v with confidence %f\n", *imageAnalysis.Adult.IsAdultContent, *imageAnalysis.Adult.AdultScore)
		fmt.Printf("Has racy content: %v with confidence %f\n", *imageAnalysis.Adult.IsRacyContent, *imageAnalysis.Adult.RacyScore)
}
//	END - Detect adult or racy content in a local image


/*  Detect adult or racy content in a remote image file by:
*    1. Saving the URL as an ImageURL type for passing to AnalyzeImage.
*    2. Defining what to extract from the image by initializing an array of VisualFeatureTypes.
*    3. Calling the Computer Vision service with the:
*       - context
*       - image
*       - features to extract
*       - an empty slice for the Details enumeration
*       - "" to specify the default language ("en") as the output language
*    4. Displaying the image categories and their confidence values.
 */
func DetectAdultOrRacyContentRemoteImage(client computervision.BaseClient, remoteImageURL string) {
	var remoteImage computervision.ImageURL
	remoteImage.URL = &remoteImageURL

	features := []computervision.VisualFeatureTypes{computervision.VisualFeatureTypesAdult}
	imageAnalysis, err := client.AnalyzeImage(
			computerVisionContext,
			remoteImage,
			features,
			[]computervision.Details{},
			"")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\nAnalyzing remote image for adult or racy content: ");
		fmt.Printf("Is adult content: %v with confidence %f\n", *imageAnalysis.Adult.IsAdultContent, *imageAnalysis.Adult.AdultScore)
		fmt.Printf("Has racy content: %v with confidence %f\n", *imageAnalysis.Adult.IsRacyContent, *imageAnalysis.Adult.RacyScore)
}
//	END - Detect adult or racy content in a remote image


/*  Detect the color scheme of a local image by:
 *    1. Instantiating a ReadCloser, which is required by AnalyzeImageInStream.
 *    2. Opening the ReadCloser instance for reading.
 *    3. Defining what to extract from the image by initializing an array of VisualFeatureTypes.
 *    4. Calling the Computer Vision service with the:
 *       - context
 *       - image
 *       - features to extract
 *       - an empty slice for the Details enumeration
 *       - "" to specify the default language ("en") as the output language
 *    5. Displaying the faces and their bounding boxes.
 */
func DetectColorSchemeLocalImage(client computervision.BaseClient, localImagePath string) {
	var localImage io.ReadCloser
	localImage, err := os.Open(localImagePath)
	if err != nil {
		log.Fatal(err)
	}

	features := []computervision.VisualFeatureTypes{computervision.VisualFeatureTypesColor}
	imageAnalysis, err := client.AnalyzeImageInStream(
			computerVisionContext,
			localImage,
			features,
			[]computervision.Details{},
			"")
		if err != nil {
			log.Fatal(err)
		}

	fmt.Println("\nColor scheme of the local image: ");
	fmt.Printf("Is black and white: %v\n", *imageAnalysis.Color.IsBWImg)
	fmt.Printf("Accent color: 0x%v\n", *imageAnalysis.Color.AccentColor)
	fmt.Printf("Dominant background color: %v\n", *imageAnalysis.Color.DominantColorBackground)
	fmt.Printf("Dominant foreground color: %v\n", *imageAnalysis.Color.DominantColorForeground)
	fmt.Printf("Dominant colors: %v\n", strings.Join(*imageAnalysis.Color.DominantColors, ", "))
}
//	END - Detect the color scheme of a local image


/*  Detect the color scheme of a remote image file by:
*    1. Saving the URL as an ImageURL type for passing to AnalyzeImage.
*    2. Defining what to extract from the image by initializing an array of VisualFeatureTypes.
*    3. Calling the Computer Vision service with the:
*       - context
*       - image
*       - features to extract
*       - an empty slice for the Details enumeration
*       - "" to specify the default language ("en") as the output language
*    4. Displaying the image categories and their confidence values.
 */
func DetectColorSchemeRemoteImage(client computervision.BaseClient, remoteImageURL string) {
	var remoteImage computervision.ImageURL
	remoteImage.URL = &remoteImageURL

	features := []computervision.VisualFeatureTypes{computervision.VisualFeatureTypesColor}
	imageAnalysis, err := client.AnalyzeImage(
			computerVisionContext,
			remoteImage,
			features,
			[]computervision.Details{},
			"")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\nColor scheme of the remote image: ");
		fmt.Printf("Is black and white: %v\n", *imageAnalysis.Color.IsBWImg)
		fmt.Printf("Accent color: 0x%v\n", *imageAnalysis.Color.AccentColor)
		fmt.Printf("Dominant background color: %v\n", *imageAnalysis.Color.DominantColorBackground)
		fmt.Printf("Dominant foreground color: %v\n", *imageAnalysis.Color.DominantColorForeground)
		fmt.Printf("Dominant colors: %v\n", strings.Join(*imageAnalysis.Color.DominantColors, ", "))
}
//	END - Detect the color scheme of a remote image


/*  Detect domain-specific content (celebrities, landmarks) in a local image by:
 *    1. Instantiating a ReadCloser, which is required by AnalyzeImageInStream.
 *    2. Opening the ReadCloser instance for reading.
 *    3. Defining what to extract from the image by initializing an array of VisualFeatureTypes.
 *    4. Calling the Computer Vision service with the:
 *       - context
 *       - image
 *       - features to extract
 *       - an enumeration specifying the domain-specific details to return
 *       - "" to specify the default language ("en") as the output language
 *    5. Displaying the faces and their bounding boxes.
 */
func DetectDomainSpecificContentLocalImage(client computervision.BaseClient, localImagePath string) {
	var localImage io.ReadCloser
	localImage, err := os.Open(localImagePath)
	if err != nil {
		log.Fatal(err)
	}

	//	TODO - Not sure you need both of these for domain-specific content detection
	features := []computervision.VisualFeatureTypes{
    computervision.VisualFeatureTypesDescription,
		computervision.VisualFeatureTypesCategories,
	}

	//	TODO - Not sure this line does anything. Seems to work if empty.
	details := []computervision.Details{computervision.Celebrities, computervision.Landmarks}

	imageAnalysis, err := client.AnalyzeImageInStream(
			computerVisionContext,
			localImage,
			features,
			details,
			"")
		if err != nil {
			log.Fatal(err)
		}

	//	TODO - Check for whether celebrities or landmarks are not returned
	//				 Right now just checking if both are there or not
	fmt.Println("\nDetecting domain-specific content in the local image ...")
	//	TODO - Check for whether celebrities or landmarks are not returned
	//				 Right now just checking if both are there or not
	if len(*imageAnalysis.Categories) == 0 {
		fmt.Println("No celebrities or landmarks detected.")
	} else {
		fmt.Println("\nCelebrities: ")
		for _, category := range *imageAnalysis.Categories {
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

		fmt.Println("\nLandmarks: ")
		for _, category := range *imageAnalysis.Categories {
			if (category.Detail != nil && category.Detail.Landmarks != nil) {
				for _, landmark := range *category.Detail.Landmarks {
					fmt.Printf("'%v' with confidence %v\n", *landmark.Name, *landmark.Confidence)
				}
			}
		}
	}
}
//	END - Detect domain-specific content in a local image


/*  Detect domain-specific content (celebrities, landmarks) in remote image file by:
*    1. Saving the URL as an ImageURL type for passing to AnalyzeImage.
*    2. Defining what to extract from the image by initializing an array of VisualFeatureTypes.
*    3. Calling the Computer Vision service with the:
*       - context
*       - image
*       - features to extract
*       - an enumeration specifying the domain-specific details to return
*       - "" to specify the default language ("en") as the output language
*    4. Displaying the image categories and their confidence values.
 */
func DetectDomainSpecificContentRemoteImage(client computervision.BaseClient, remoteImageURL string) {
	var remoteImage computervision.ImageURL
	remoteImage.URL = &remoteImageURL

	//	TODO - Not sure you need both of these for domain-specific content detection
	features := []computervision.VisualFeatureTypes{
			computervision.VisualFeatureTypesDescription,
			computervision.VisualFeatureTypesCategories,
	}

	//	TODO - Not sure this line does anything. Seems to work if empty.
	details := []computervision.Details{computervision.Celebrities, computervision.Landmarks}

	imageAnalysis, err := client.AnalyzeImage(
			computerVisionContext,
			remoteImage,
			features,
			details,
			"")
		if err != nil {
			log.Fatal(err)
		}

	//	TODO - Check for whether celebrities or landmarks are not returned
	//				 Right now just checking if both are there or not
	fmt.Println("\nDetecting domain-specific content in the remote image ...")
	if len(*imageAnalysis.Categories) == 0 {
		fmt.Println("No celebrities or landmarks detected.")
	} else {
	  fmt.Println("\nCelebrities: ")
	  for _, category := range *imageAnalysis.Categories {
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

	  fmt.Println("\nLandmarks: ")
	  for _, category := range *imageAnalysis.Categories {
		  if (category.Detail != nil && category.Detail.Landmarks != nil) {
			  for _, landmark := range *category.Detail.Landmarks {
				  fmt.Printf("'%v' with confidence %v\n", *landmark.Name, *landmark.Confidence)
		  	}
		  }
	  }
  }
}
//	END - Detect domain-specific content in a remote image


/*  Detect the image type (clip art, line drawing) of a local image by:
 *    1. Instantiating a ReadCloser, which is required by AnalyzeImageInStream.
 *    2. Opening the ReadCloser instance for reading.
 *    3. Defining what to extract from the image by initializing an array of VisualFeatureTypes.
 *    4. Calling the Computer Vision service with the:
 *       - context
 *       - image
 *       - features to extract
 *       - an empty slice for the Details enumeration
 *       - "" to specify the default language ("en") as the output language
 *    5. Displaying the faces and their bounding boxes.
 */
func DetectImageTypesLocalImage(client computervision.BaseClient, localImagePath string) {
	var localImage io.ReadCloser
	localImage, err := os.Open(localImagePath)
	if err != nil {
		log.Fatal(err)
	}

	features := []computervision.VisualFeatureTypes{computervision.VisualFeatureTypesImageType}

	imageAnalysis, err := client.AnalyzeImageInStream(
			computerVisionContext,
			localImage,
			features,
			[]computervision.Details{},
			"")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\nImage type:")
		fmt.Printf("Clip art type: %v\n", *imageAnalysis.ImageType.ClipArtType)
		fmt.Printf("Line drawing type: %v\n", *imageAnalysis.ImageType.LineDrawingType)
		fmt.Printf("\nFor information about the values returned by the image type detection, please see:\nhttps://docs.microsoft.com/en-us/azure/cognitive-services/computer-vision/concept-detecting-image-types.\n")
}
//	END - Detect image type of a local image


/*  Detect the image type (clip art, line drawing) of a remote image by:
*    1. Saving the URL as an ImageURL type for passing to AnalyzeImage.
*    2. Defining what to extract from the image by initializing an array of VisualFeatureTypes.
*    3. Calling the Computer Vision service with the:
*       - context
*       - image
*       - features to extract
*       - an enumeration specifying the domain-specific details to return
*       - "" to specify the default language ("en") as the output language
*    4. Displaying the image categories and their confidence values.
 */
func DetectImageTypesRemoteImage(client computervision.BaseClient, remoteImageURL string) {
	var remoteImage computervision.ImageURL
	remoteImage.URL = &remoteImageURL

	features := []computervision.VisualFeatureTypes{computervision.VisualFeatureTypesImageType}

	imageAnalysis, err := client.AnalyzeImage(
			computerVisionContext,
			remoteImage,
			features,
			[]computervision.Details{},
			"")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\nImage type:")
		fmt.Printf("Clip art type: %v\n", *imageAnalysis.ImageType.ClipArtType)
		fmt.Printf("Line drawing type: %v\n", *imageAnalysis.ImageType.LineDrawingType)
		fmt.Printf("\nFor information about the values returned by the image type detection, please see:\nhttps://docs.microsoft.com/en-us/azure/cognitive-services/computer-vision/concept-detecting-image-types.\n")
	}
//	END - Detect image type of a remote image
