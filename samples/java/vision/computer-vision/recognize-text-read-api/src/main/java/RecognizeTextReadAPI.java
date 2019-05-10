import com.microsoft.azure.cognitiveservices.vision.computervision.*;
import com.microsoft.azure.cognitiveservices.vision.computervision.models.*;

import java.io.File;
import java.nio.file.Files;

import java.util.ArrayList;
import java.util.List;

public class RecognizeTextReadAPI {

    public static String subKey = System.getenv("AZURE_COMPUTERVISION_API_KEY");
    public static String baseURL = System.getenv("AZURE_ENDPOINT");

    public static void main(String[] args) {
        try {
            RecognizeTextReadAPISample.RunSample(baseURL, subKey);
            
        } catch (Exception e) {
            System.out.println(e.getMessage());
            e.printStackTrace();
        }
    }
    
    private static class RecognizeTextReadAPISample {
        public static void RunSample(String url, String key) {
            ComputerVisionClient compVisClient = ComputerVisionManager.authenticate(key).withEndpoint(url);
            //System.out.println("compVisClient.endpoint(): " + compVisClient.endpoint());
            
            String imgPath = "src\\main\\resources\\printed_text.jpg";
            String remotePath = "https://raw.githubusercontent.com/Azure-Samples/cognitive-services-sample-data-files/master/ComputerVision/Images/handwritten_text.jpg";
            
            int numCharsInOperationId = 36;
            
            System.out.println("\nRecognizing printed text with the Read API on a local image ...");
            RecognizeTextReadAPILocal(compVisClient, imgPath, numCharsInOperationId);
            
            //System.out.println("\nRecognizing handwritten text with the Read API on a remote image ...");
            //RecognizeTextReadAPIFromUrl(compVisClient, remotePath, numCharsInOperationId);
        }
        
        private static void RecognizeTextReadAPILocal(ComputerVisionClient client, String path, int numCharsInOpId) {
            
            TextRecognitionMode textRecogMode = TextRecognitionMode.PRINTED;
                              
            try {
                File rawImage = new File(path);
                byte[] imageBytes = Files.readAllBytes(rawImage.toPath());

                RecognizeTextInStreamHeaders textHeaders = client.computerVision().recognizeTextInStreamHeaders()
                    .withImage(imageBytes)
                    .withTextRecognitionMode(textRecogMode)
                    .execute();
                
//                BatchReadFileInStreamHeaders textHeaders = await computerVision.BatchReadFileInStreamAsync(imageStream, textRecognitionMode);
  //              await GetTextAsync(computerVision, textHeaders.OperationLocation, numberOfCharsInOperationId);               


            } catch (Exception e) {
                System.out.println(e.getMessage());
                e.printStackTrace();
            }
        }
        
        private static void GetTextAsync(ComputerVisionClient client, String path, int numCharsInOpId) {
            
        }
    }
}
