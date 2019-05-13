import com.microsoft.azure.cognitiveservices.vision.computervision.*;
import com.microsoft.azure.cognitiveservices.vision.computervision.models.*;

import java.io.File;
import java.io.FileInputStream;
import java.nio.file.Files;

import java.util.ArrayList;
import java.util.List;

public class ComputerVisionQuickstarts
{   
    public static void main(String[] args)
    {
        RunQuickstarts();
    }
    
    public static void RunQuickstarts()
    {
        /*  Configure the local environment:
         *
         *  Set the AZURE_COMPUTERVISION_API_KEY and AZURE_REGION environment variables on your
         *  local machine using the appropriate method for your preferred command shell.
         *
         *  For AZURE_REGION, use the same region you used to get your subscription keys.
         ***Can we link to docs for the regions?
         *
         *  Note that environment variables cannot contain quotation marks, so the quotation marks
         *  are included in the code below to stringify them.
         *  
         *  Note that after setting these environment variables in your preferred command shell,
         *  you will need to close and then re-open your command shell.  
         */

        String azureComputerVisionApiKey = System.getenv("AZURE_COMPUTERVISION_API_KEY");
        String azureRegion = System.getenv("AZURE_REGION");
        //  END - Configure the local environment.
        
        
        /*  Create an authenticated Computer Vision client:
         *  
         *  Concatenate the Azure region with the Azure base URL to create the endpoint URL, and
         *  then create an authenticated client with the API key and the endpoint URL.
         */
        
        String endpointUrl = ("https://").concat(azureRegion).concat(".api.cognitive.microsoft.com");
        ComputerVisionClient compVisClient = ComputerVisionManager.authenticate(azureComputerVisionApiKey).withEndpoint(endpointUrl);
        //  END - Create an authenticated Computer Vision client.
    
    
        /*  Analyze a local image:
         *  
         *  Set a string variable equal to the path of a local image. The image path below is a relative path.  
         */
        String pathToLocalImage = "src\\main\\resources\\landmark.jpg";
//        String remotePath = "https://github.com/Azure-Samples/cognitive-services-sample-data-files/raw/master/ComputerVision/Images/landmark.jpg";
    
        //  This list defines the features to be extracted from the image. 
        List<VisualFeatureTypes> featuresToExtractFromImage = new ArrayList<>();
        featuresToExtractFromImage.add(VisualFeatureTypes.DESCRIPTION);
        featuresToExtractFromImage.add(VisualFeatureTypes.CATEGORIES);
        featuresToExtractFromImage.add(VisualFeatureTypes.TAGS);
        featuresToExtractFromImage.add(VisualFeatureTypes.FACES);
        featuresToExtractFromImage.add(VisualFeatureTypes.ADULT);
        featuresToExtractFromImage.add(VisualFeatureTypes.COLOR);
        featuresToExtractFromImage.add(VisualFeatureTypes.IMAGE_TYPE);
        
        System.out.println("\nAnalyzing local image ...");       
        
        try
        {
            File rawImg = new File(pathToLocalImage);
            byte[] imgBytes = Files.readAllBytes(rawImg.toPath());
        
            ImageAnalysis analysis = compVisClient.computerVision().analyzeImageInStream()
                .withImage(imgBytes)
                .withVisualFeatures(featuresToExtractFromImage)
                .execute();
        
            
            System.out.println("\nCaptions: ");
            for (ImageCaption caption : analysis.description().captions()) {
                System.out.printf("\'%s\' with confidence %f\n", caption.text(), caption.confidence());
            }
            
            System.out.println("\nCategories: ");
            for (Category category : analysis.categories()) {
                System.out.printf("\'%s\' with confidence %f\n", category.name(), category.score());
            }

            System.out.println("\nTags: ");
            for (ImageTag tag : analysis.tags()) {
                System.out.printf("\'%s\' with confidence %f\n", tag.name(), tag.confidence());
            }           
            
            /*  Switch to an image with faces so we get a result.
             *
             */
            System.out.println("\nFaces: ");
            for (FaceDescription face : analysis.faces()) {
                System.out.printf("\'%s\' of age %d at location (%d, %d), (%d, %d)\n", face.gender(), face.age(),
                    face.faceRectangle().left(), face.faceRectangle().top(),
                    face.faceRectangle().left() + face.faceRectangle().width(),
                    face.faceRectangle().top() + face.faceRectangle().height());
            }            


            System.out.println("\nAdult: ");
            System.out.printf("Is adult content: %b with confidence %f\n", analysis.adult().isAdultContent(), analysis.adult().adultScore());
            System.out.printf("Has racy content: %b with confidence %f\n", analysis.adult().isRacyContent(), analysis.adult().racyScore());


            System.out.println("\nColor scheme: ");
            System.out.println("Is black and white: " + analysis.color().isBWImg());
            System.out.println("Accent color: " + analysis.color().accentColor());
            System.out.println("Dominant background color: " + analysis.color().dominantColorBackground());
            System.out.println("Dominant foreground color: " + analysis.color().dominantColorForeground());
            System.out.println("Dominant colors: " + String.join(", ", analysis.color().dominantColors()));


            System.out.println("\nCelebrities: ");
            for (Category category : analysis.categories())
            {
                if (category.detail() != null && category.detail().celebrities() != null)
                {
                    for (CelebritiesModel celeb : category.detail().celebrities())
                    {
                        System.out.printf("\'%s\' with confidence %f at location (%d, %d), (%d, %d)\n", celeb.name(), celeb.confidence(),
                            celeb.faceRectangle().left(), celeb.faceRectangle().top(),
                            celeb.faceRectangle().left() + celeb.faceRectangle().width(),
                            celeb.faceRectangle().top() + celeb.faceRectangle().height());
                    }
                }
            }
            
            System.out.println("\nLandmarks: ");
            for (Category category : analysis.categories())
            {
                if (category.detail() != null && category.detail().landmarks() != null)
                {
                    for (LandmarksModel landmark : category.detail().landmarks())
                    {
                        System.out.printf("\'%s\' with confidence %f\n", landmark.name(), landmark.confidence());
                    }
                }
            }


            System.out.println("\nImage type:");
            System.out.println("Clip art type: " + analysis.imageType().clipArtType());
            System.out.println("Line drawing type: " + analysis.imageType().lineDrawingType());
        }
        
        catch (Exception e)
        {
            System.out.println(e.getMessage());
            e.printStackTrace();
        }
}
    
    private static void AnalyzeFromUrl(ComputerVisionClient client, String path, List<VisualFeatureTypes> feats) {
        try {
            ImageAnalysis analysis = client.computerVision().analyzeImage()
                .withUrl(path)
                .withVisualFeatures(feats)
                .execute();
            
        } catch (Exception e) {
            System.out.println(e.getMessage());
            e.printStackTrace();
        }
    }
}
