import com.google.gson.*;

import com.microsoft.azure.cognitiveservices.vision.contentmoderator.*;
import com.microsoft.azure.cognitiveservices.vision.contentmoderator.models.*;

import java.io.*;
import java.lang.Object.*;
import java.util.*;

/** 
 * 1. Obtain Azure resource for service
 * 2. Ensure correct Java version, ex: Java 8 or later
 * 3. Install/add dependencies and/or 3rd party libraries
 * 4. Follow IDE or command line instruction
 */ 

public class ImageModeration {
   /*
    * Contains the image moderation results for an image, including
    * text and face detection results.
    */
    public static class EvaluationData {
        // The URL of the evaluated image.
        public String ImageUrl;

        // The image moderation results.
        public Evaluate ImageModeration;

        // The text detection results.
        public OCR TextDetection;

        // The face detection results;
        public FoundFaces FaceDetection;
    }

    /*
     * The name of the file that contains the image URLs to evaluate.
     * You will need to create an input file and update this path
     * accordingly. Relative paths are relative to the execution directory.
    */
    private static String ImageUrlFile = "C:\\Users\\v-lubayl\\Documents\\Github\\cognitive-services-samples\\java\\content-moderator\\moderate-images\\src\\main\\resources\\ImageFiles.txt";
 
    /*
     * The name of the file to contain the output from the evaluation.
     * Relative paths are relative the execution directory.
     */
    private static String OutputFile = "C:\\Users\\v-lubayl\\Documents\\GitHub\\cognitive-services-samples\\java\\content-moderator\\moderate-images\\src\\main\\resources\\ModerationOutput.json";

    public static void main(String[] args) {
    
        // The base URL fragment for Content Moderator calls.
        String AzureBaseURL = "https://westus.api.cognitive.microsoft.com";

        // Your Content Moderator subscription key.
        String CMSubscriptionKey = "bed9632798b9496bab97d18e31d0fde9";
        
        ContentModeratorClient client = ContentModeratorManager.authenticate(new AzureRegionBaseUrl().fromString(AzureBaseURL), CMSubscriptionKey);
        System.out.println("baseUrl(): " + client.baseUrl());
        
        // Create an object in which to store the image moderation results.
        List<EvaluationData> evaluationData = new ArrayList<EvaluationData>();

        // Read image URLs from the input file and evaluate each one.
        try (BufferedReader inputStream = new BufferedReader(new FileReader(new File(ImageUrlFile)))) {
            String line;
            while ((line = inputStream.readLine()) != null) {
                if (line.length() > 0) {
                    //System.out.println("checking line in file: " + line);
                    //System.out.println("checking baseUrl(): " + client.baseUrl());
                    EvaluationData imageData = EvaluateImage(client, line);
                    
                    //System.out.println("adding imageData to list: " + imageData);
                    evaluationData.add(imageData);
                }
            }
        }   catch (Exception e) {
                System.out.println(e.getMessage());
                e.printStackTrace();
        }

        // Save the moderation results to a file.
        try (BufferedWriter writer = new BufferedWriter(new FileWriter(new File(OutputFile)))) {
        
            // Maybe don't need this next line in this try block?
            Gson gson = new GsonBuilder().setPrettyPrinting().create();            

            System.out.println("adding imageData to file: " + gson.toJson(evaluationData).toString());
            writer.write(gson.toJson(evaluationData).toString());
        
        }   catch (Exception e) {
                System.out.println(e.getMessage());
                e.printStackTrace();
        }
    }

    /*
     * Evaluates an image using the Image Moderation APIs.
     * This method throttles calls to the API.
     * Your Content Moderator service key will have a requests per second (RPS)
     * rate limit, and the SDK will throw an exception with a 429 error code
     * if you exceed that limit. A free tier key has a 1 RPS rate limit.
     * @param client The Content Moderator API wrapper to use.
     * @param imageUrl The URL of the image to evaluate.
     * @return Aggregated image moderation results for the image.
    */
    private static EvaluationData EvaluateImage(ContentModeratorClient client, String imageUrl) throws InterruptedException {
        BodyModelModel url = new BodyModelModel();
        url.withDataRepresentation("URL");
        url.withValue(imageUrl);
        
        EvaluationData imageData = new EvaluationData();     
        imageData.ImageUrl = url.value();
        
        //System.out.println("url: " + url);
        //System.out.println("imageData.ImageUrl: " + imageData.ImageUrl);

        // Evaluate for adult and racy content.
        imageData.ImageModeration = client.imageModerations().evaluateUrlInput("application/json", url, new EvaluateUrlInputOptionalParameter().withCacheImage(true));
        Thread.sleep(1000);

        // Detect and extract text.
        imageData.TextDetection = client.imageModerations().oCRUrlInput("eng", "application/json", url, new OCRUrlInputOptionalParameter().withCacheImage(true));
        Thread.sleep(1000);

        // Detect faces.
        imageData.FaceDetection = client.imageModerations().findFacesUrlInput("application/json", url, new FindFacesUrlInputOptionalParameter().withCacheImage(true));
        Thread.sleep(1000);

        return imageData;
    } 
}
