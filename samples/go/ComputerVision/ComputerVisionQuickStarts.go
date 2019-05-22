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
 *  - Displaying image captions
 *  - Displaying image category names
 *  - Displaying image tags
 *  - Displaying whether any faces were found in the image
 *  - Detecting adult or racy content
 *  - Displaying the color scheme of the image
 *  - Detecting any domain-specific content (celebrities/landmarks)
 *  - Detecting is the image is clip art or a line drawing
 *  - Recognizing printed and handwritten text with the Read API
 *  - Recognizing printed and handwritten text with OCR
 *  - Detecting objects in the image
 *  - Detecting brands in the image
 */
