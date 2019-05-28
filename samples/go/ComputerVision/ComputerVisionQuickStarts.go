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
	"time"
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
 *  - Detecting objects
 *  - Detecting brands
 *  - Recognizing printed and handwritten text with (Read API)
 *  - Recognizing printed and handwritten text with (OCR)
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


	//	Analyze a local image
	localImagePath := "resources\\faces.jpg"
	workingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n\nLocal image path:\n%v\n", workingDirectory + "\\" + localImagePath)

	DescribeLocalImage(computerVisionClient, localImagePath)
	CategorizeLocalImage(computerVisionClient, localImagePath)
	TagLocalImage(computerVisionClient, localImagePath)
	DetectFacesLocalImage(computerVisionClient, localImagePath)
	DetectAdultOrRacyContentLocalImage(computerVisionClient, localImagePath)
	DetectColorSchemeLocalImage(computerVisionClient, localImagePath)
	DetectDomainSpecificContentLocalImage(computerVisionClient, localImagePath)
	DetectImageTypesLocalImage(computerVisionClient, localImagePath)
	DetectObjectsLocalImage(computerVisionClient, localImagePath)
	//	END - Analyze a local iamge

	//	Brand detection on a local image
	fmt.Println("\nGetting new local image for brand detection ... ")
	localImagePath = "resources\\gray-shirt-logo.jpg"
	workingDirectory, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Local image path:\n%v\n", workingDirectory + "\\" + localImagePath)

	DetectBrandsLocalImage(computerVisionClient, localImagePath)
	//	END - Brand detection


	//	Text recognition on a local image
	fmt.Println("\nGetting new local image for text recognition ... ")
	localImagePath = "resources\\printed_text.jpg"
	workingDirectory, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Local image path:\n%v\n", workingDirectory + "\\" + localImagePath)

	RecognizeTextReadAPILocalImage(computerVisionClient, localImagePath)
	RecognizeTextOCRLocalImage(computerVisionClient, localImagePath)
	//	END - Text recognition on a local image


	//	Analyze a remote image
	remoteImageURL := "https://github.com/Azure-Samples/cognitive-services-sample-data-files/raw/master/ComputerVision/Images/faces.jpg"
	fmt.Printf("\n\nRemote image path: \n%v\n", remoteImageURL)

	DescribeRemoteImage(computerVisionClient, remoteImageURL)
	CategorizeRemoteImage(computerVisionClient, remoteImageURL)
	TagRemoteImage(computerVisionClient, remoteImageURL)
	DetectFacesRemoteImage(computerVisionClient, remoteImageURL)
	DetectAdultOrRacyContentRemoteImage(computerVisionClient, remoteImageURL)
	DetectColorSchemeRemoteImage(computerVisionClient, remoteImageURL)
	DetectDomainSpecificContentRemoteImage(computerVisionClient, remoteImageURL)
	DetectImageTypesRemoteImage(computerVisionClient, remoteImageURL)
	DetectObjectsRemoteImage(computerVisionClient, remoteImageURL)
	//	END - Analyze a remote image

	//	Brand detection on a local image
	fmt.Println("\nGetting new remote image for brand recognition ... ")
	remoteImageURL = "https://docs.microsoft.com/en-us/azure/cognitive-services/computer-vision/images/gray-shirt-logo.jpg"
	fmt.Printf("Remote image path: \n%v\n", remoteImageURL)

	DetectBrandsRemoteImage(computerVisionClient, remoteImageURL)
	//	END - Brand detection on a remote image


	//	Text recognition on a remote image
	fmt.Println("\nGetting new remote image for text recognition ... ")
	remoteImageURL = "https://raw.githubusercontent.com/Azure-Samples/cognitive-services-sample-data-files/master/ComputerVision/Images/handwritten_text.jpg"
	fmt.Printf("Remote image path: \n%v\n", remoteImageURL)

	RecognizeTextReadAPIRemoteImage(computerVisionClient, remoteImageURL)
	RecognizeTextOCRRemoteImage(computerVisionClient, remoteImageURL)
	//	END - Text recognition on a remote image
}

/*  Describe a local image by:
 *    1. Instantiating a ReadCloser, which is required by AnalyzeImageInStream.
 *    2. Opening the ReadCloser instance for reading.
 *    3. Defining what to extract from the image by initializing an array of VisualFeatureTypes.
 *    4. Calling the Computer Vision service's AnalyzeImageInStream with the:
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
*    3. Calling the Computer Vision service's AnalyzeImage with the:
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
 *    4. Calling the Computer Vision service's AnalyzeImageInStream with the:
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
*    3. Calling the Computer Vision service's AnalyzeImage with the:
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
 *    4. Calling the Computer Vision service's AnalyzeImageInStream with the:
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
	*    3. Calling the Computer Vision service's AnalyzeImage with the:
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
 *    4. Calling the Computer Vision service's AnalyzeImageInStream with the:
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
*    3. Calling the Computer Vision service's AnalyzeImage with the:
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
 *    4. Calling the Computer Vision service's AnalyzeImageInStream with the:
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
*    3. Calling the Computer Vision service's AnalyzeImage with the:
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
 *    4. Calling the Computer Vision service's AnalyzeImageInStream with the:
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
*    3. Calling the Computer Vision service's AnalyzeImage with the:
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
 *    4. Calling the Computer Vision service's AnalyzeImageInStream with the:
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
*    3. Calling the Computer Vision service's AnalyzeImage with the:
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
 *    4. Calling the Computer Vision service's AnalyzeImageInStream with the:
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
*    3. Calling the Computer Vision service's AnalyzeImage with the:
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


/*  Detect objects in a local image by:
 *    1. Instantiating a ReadCloser, which is required by DetectObjectsInStream.
 *    2. Opening the ReadCloser instance for reading.
 *    3. Calling the Computer Vision service's DetectObjectsInStream with the:
 *       - context
 *       - image
 *    4. Displaying the objects and their bounding boxes.
 */
func DetectObjectsLocalImage(client computervision.BaseClient, localImagePath string) {
	var localImage io.ReadCloser
	localImage, err := os.Open(localImagePath)
	if err != nil {
		log.Fatal(err)
	}

	imageAnalysis, err := client.DetectObjectsInStream(
			computerVisionContext,
			localImage,
			)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\nDetecting objects in local image: ")
		if len(*imageAnalysis.Objects) == 0 {
			fmt.Println("No objects detected.")
		}
		for _, object := range *imageAnalysis.Objects {
			fmt.Printf("'%v' with confidence %v at location (%v, %v), (%v, %v)\n",
				*object.Object, *object.Confidence,
				*object.Rectangle.X, *object.Rectangle.X + *object.Rectangle.W,
				*object.Rectangle.Y, *object.Rectangle.Y + *object.Rectangle.H)
		}
}
//	END - Detect objects in local image


/*  Detect objects in a remote image by:
*    1. Saving the URL as an ImageURL type for passing to DetectObjects.
*    2. Calling the Computer Vision service's DetectObjects with the:
*       - context
*       - image
*    3. Displaying the objects and their bounding boxes.
 */
func DetectObjectsRemoteImage(client computervision.BaseClient, remoteImageURL string) {
	var remoteImage computervision.ImageURL
	remoteImage.URL = &remoteImageURL

	imageAnalysis, err := client.DetectObjects(
			computerVisionContext,
			remoteImage,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nDetecting objects in remote image: ")
	if len(*imageAnalysis.Objects) == 0 {
		fmt.Println("No objects detected.")
	}
	for _, object := range *imageAnalysis.Objects {
		fmt.Printf("'%v' with confidence %v at location (%v, %v), (%v, %v)\n",
			*object.Object, *object.Confidence,
			*object.Rectangle.X, *object.Rectangle.X + *object.Rectangle.W,
			*object.Rectangle.Y, *object.Rectangle.Y + *object.Rectangle.H)
	}
}
//	END - Detect objects in remote image


/*  Detect brands in a local image by:
 *    1. Instantiating a ReadCloser, which is required by AnalyzeImageInStream.
 *    2. Opening the ReadCloser instance for reading.
 *    3. Defining what to extract from the image by initializing an array of VisualFeatureTypes.
 *    4. Calling the Computer Vision service's AnalyzeImageInStream with the:
 *       - context
 *       - image
 *       - features to extract
 *       - an empty slice for the Details enumeration
 *       - "" to specify the default language ("en") as the output language
 *    5. Displaying the brands, confidence values, and their bounding boxes.
 */
func DetectBrandsLocalImage(client computervision.BaseClient, localImagePath string) {
	var localImage io.ReadCloser
	localImage, err := os.Open(localImagePath)
	if err != nil {
		log.Fatal(err)
	}

	features := []computervision.VisualFeatureTypes{computervision.VisualFeatureTypesBrands}

	imageAnalysis, err := client.AnalyzeImageInStream(
			computerVisionContext,
			localImage,
			features,
			[]computervision.Details{},
			"en")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\nDetecting brands in local image: ")
		if len(*imageAnalysis.Brands) == 0 {
			fmt.Println("No brands detected.")
		}
		for _, brand := range *imageAnalysis.Brands {
			fmt.Printf("'%v' with confidence %v at location (%v, %v), (%v, %v)\n",
				*brand.Name, *brand.Confidence,
	      *brand.Rectangle.X, *brand.Rectangle.X + *brand.Rectangle.W,
	      *brand.Rectangle.Y, *brand.Rectangle.Y + *brand.Rectangle.H)
		}
}
//	END - Detect brands in local image


/*  Detect brands in a remote image by:
*    1. Saving the URL as an ImageURL type for passing to AnalyzeImage.
*    2. Defining what to extract from the image by initializing an array of VisualFeatureTypes.
*    3. Calling the Computer Vision service's AnalyzeImage with the:
*       - context
*       - image
*       - features to extract
*       - an enumeration specifying the domain-specific details to return
*       - "" to specify the default language ("en") as the output language
*    5. Displaying the brands, confidence values, and their bounding boxes.
 */
func DetectBrandsRemoteImage(client computervision.BaseClient, remoteImageURL string) {
	var remoteImage computervision.ImageURL
	remoteImage.URL = &remoteImageURL

	features := []computervision.VisualFeatureTypes{computervision.VisualFeatureTypesBrands}

	imageAnalysis, err := client.AnalyzeImage(
		computerVisionContext,
		remoteImage,
		features,
		[]computervision.Details{},
		"en")
	if err != nil {
	 	log.Fatal(err)
	}

	fmt.Println("\nDetecting brands in remote image: ")
	if len(*imageAnalysis.Brands) == 0 {
		fmt.Println("No brands detected.")
	}
	for _, brand := range *imageAnalysis.Brands {
		fmt.Printf("'%v' with confidence %v at location (%v, %v), (%v, %v)\n",
			*brand.Name, *brand.Confidence,
			*brand.Rectangle.X, *brand.Rectangle.X + *brand.Rectangle.W,
			*brand.Rectangle.Y, *brand.Rectangle.Y + *brand.Rectangle.H)
	}
}
//	END - Detect brands in remote image


/*  Recognize text with the Read API in a local image by:
 *    1. Instantiating a ReadCloser, which is required by BatchReadFileInStream.
 *    2. Opening the ReadCloser instance for reading.
 *    3. Specifying whether the text to recognize is handwritten or printed.
 *    4. Calling the Computer Vision service's BatchReadFileInStream with the:
 *       - context
 *       - image
 *       - text recognition mode
 *    5. Extracting the Operation-Location URL value from the BatchReadFileInStream
 *       response
 *    6. Waiting for the operation to complete.
 *    7. Displaying the results.
 */
func RecognizeTextReadAPILocalImage(client computervision.BaseClient, localImagePath string) {
	var localImage io.ReadCloser
	localImage, err := os.Open(localImagePath)
	if err != nil {
		log.Fatal(err)
	}

	textRecognitionMode := computervision.Printed

	//	When you use the Read Document interface, the response contains a field
	//	called "Operation-Location", which contains the URL to use for your
	//	GetReadOperationResult to access OCR results.
	textHeaders, err := client.BatchReadFileInStream(
		computerVisionContext,
		localImage,
		textRecognitionMode)
	if err != nil {
		log.Fatal(err)
	}

	//	Use ExtractHeader from the autorest library to get the Operation-Location URL
	operationLocation := autorest.ExtractHeaderValue("Operation-Location", textHeaders.Response)

	numberOfCharsInOperationId := 36
	operationId := string(operationLocation[len(operationLocation)-numberOfCharsInOperationId : len(operationLocation)])

	readOperationResult, err := client.GetReadOperationResult(computerVisionContext, operationId)
	if err != nil {
		log.Fatal(err)
	}

	// Wait for the operation to complete.
	i := 0
	maxRetries := 10

	fmt.Println("Recognizing text in a local image with the batch Read API ... ")
	for readOperationResult.Status != computervision.Failed &&
			readOperationResult.Status != computervision.Succeeded {
		if i >= maxRetries {
			break
		}
		i++

		fmt.Printf("Server status: %v, waiting %v seconds...\n", readOperationResult.Status, i)
		time.Sleep(1 * time.Second)

		readOperationResult, err = client.GetReadOperationResult(computerVisionContext, operationId)
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
}
//	END - Recognize text with the Read API in a local image


/*  Recognize text with the Read API in a remote image by:
 *    1. Saving the URL as an ImageURL type for passing to BatchReadFile.
 *    2. Specifying whether the text to recognize is handwritten or printed.
 *    3. Calling the Computer Vision service's BatchReadFile with the:
 *       - context
 *       - image
 *       - text recognition mode
 *    4. Extracting the Operation-Location URL value from the BatchReadFile
 *       response
 *    5. Waiting for the operation to complete.
 *    6. Displaying the results.
 */
func RecognizeTextReadAPIRemoteImage(client computervision.BaseClient, remoteImageURL string) {
	var remoteImage computervision.ImageURL
	remoteImage.URL = &remoteImageURL

	textRecognitionMode := computervision.Printed

	//	When you use the Read Document interface, the response contains a field
	//	called "Operation-Location", which contains the URL to use for your
	//	GetReadOperationResult to access OCR results.
	textHeaders, err := client.BatchReadFile(
		computerVisionContext,
		remoteImage,
		textRecognitionMode)
	if err != nil {
		log.Fatal(err)
	}

	//	Use ExtractHeader from the autorest library to get the Operation-Location URL
	operationLocation := autorest.ExtractHeaderValue("Operation-Location", textHeaders.Response)

	numberOfCharsInOperationId := 36
	operationId := string(operationLocation[len(operationLocation)-numberOfCharsInOperationId : len(operationLocation)])

	readOperationResult, err := client.GetReadOperationResult(computerVisionContext, operationId)
	if err != nil {
		log.Fatal(err)
	}

	// Wait for the operation to complete.
	i := 0
	maxRetries := 10

	fmt.Println("Recognizing text in a remote image with the batch Read API ... ")
	for readOperationResult.Status != computervision.Failed &&
			readOperationResult.Status != computervision.Succeeded {
		if i >= maxRetries {
			break
		}
		i++

		fmt.Printf("Server status: %v, waiting %v seconds...\n", readOperationResult.Status, i)
		time.Sleep(1 * time.Second)

		readOperationResult, err = client.GetReadOperationResult(computerVisionContext, operationId)
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
}


/*  Recognize text with OCR in a local image by:
 *    1. Instantiating a ReadCloser, which is required by RecognizePrintedTextInStream.
 *    2. Opening the ReadCloser instance for reading.
 *    3. Calling the Computer Vision service's RecognizePrintedTextInStream with the:
 *       - context
 *       - image
 *       - whether to detect the orientation of the text before processing
 *       - language code of the text to detect
 *    4. Displaying the results.
 */
func RecognizeTextOCRLocalImage(client computervision.BaseClient, localImagePath string) {
	var localImage io.ReadCloser
	localImage, err := os.Open(localImagePath)
	if err != nil {
		log.Fatal(err)
	}

	localImageOcrResult, err := client.RecognizePrintedTextInStream(
			computerVisionContext,
      true,
      localImage,
			computervision.En)
		if err != nil {
			log.Fatal(err)
		}

	fmt.Println("\nRecognizing text in a local image with OCR ... ")
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
}
//	END - Recognize text with OCR in a local image


/*  Recognize text with OCR in a remote image by:
 *    1. Saving the URL as an ImageURL type for passing to AnalyzeImage.
 *    2. Calling the Computer Vision service's RecognizePrintedTextInStream with the:
 *       - context
 *       - image
 *       - whether to detect the orientation of the text before processing
 *       - language code of the text to detect
 *    3. Displaying the results.
 */
func RecognizeTextOCRRemoteImage(client computervision.BaseClient, remoteImageURL string) {
	var remoteImage computervision.ImageURL
	remoteImage.URL = &remoteImageURL

	remoteImageOcrResult, err := client.RecognizePrintedText(
		computerVisionContext,
    true,
    remoteImage,
		"en")
	if err != nil {
	 	log.Fatal(err)
	}

	fmt.Println("\nRecognizing text in a remote image with OCR ... ")
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
}
//	END - Recognize text with OCR in a remote image
