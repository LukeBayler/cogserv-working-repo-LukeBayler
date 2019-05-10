import com.microsoft.azure.cognitiveservices.vision.computervision.*;
import com.microsoft.azure.cognitiveservices.vision.computervision.models.*;

import java.io.File;
import java.nio.file.Files;

import java.util.ArrayList;
import java.util.List;

public class RecognizeTextOCR {

    public static String subKey = System.getenv("AZURE_COMPUTERVISION_API_KEY");
    public static String baseURL = System.getenv("AZURE_ENDPOINT");

    public static void main(String[] args) {
        try {
            RecognizeTextOCRSample.RunSample(baseURL, subKey);
            
        } catch (Exception e) {
            System.out.println(e.getMessage());
            e.printStackTrace();
        }
    }
    
    private static class RecognizeTextOCRSample {
        public static void RunSample(String url, String key) {
            ComputerVisionClient compVisClient = ComputerVisionManager.authenticate(key).withEndpoint(url);
            //System.out.println("compVisClient.endpoint(): " + compVisClient.endpoint());
            
            String imgPath = "src\\main\\resources\\printed_text.jpg";
            String remotePath = "https://raw.githubusercontent.com/Azure-Samples/cognitive-services-sample-data-files/master/ComputerVision/Images/handwritten_text.jpg";
            
            System.out.println("\nRecognizing printed text with OCR on a local image ...");
            RecognizeTextOCRLocal(compVisClient, imgPath);
            
            System.out.println("\nRecognizing handwritten text with OCR on a remote image ...");
            RecognizeTextOCRFromUrl(compVisClient, remotePath);
        }
        
        private static void RecognizeTextOCRLocal(ComputerVisionClient client, String path) {
            try {
                File rawImage = new File(path);
                byte[] imageBytes = Files.readAllBytes(rawImage.toPath());
                
                OcrResult ocrResult = client.computerVision().recognizePrintedTextInStream()
                    .withDetectOrientation(true)
                    .withImage(imageBytes)
                    .withLanguage(OcrLanguages.EN)
                    .execute();
                    
                DisplayResults(ocrResult);
               
            } catch (Exception e) {
                System.out.println(e.getMessage());
                e.printStackTrace();
            }
        }

        private static void RecognizeTextOCRFromUrl(ComputerVisionClient client, String path) {
            try {
                OcrResult ocrResult = client.computerVision().recognizePrintedText()
                    .withDetectOrientation(true)
                    .withUrl(path)
                    .withLanguage(OcrLanguages.EN)
                    .execute();
                    
                DisplayResults(ocrResult);
               
            } catch (Exception e) {
                System.out.println(e.getMessage());
                e.printStackTrace();
            }
        }
        
        private static void DisplayResults(OcrResult ocrResult) {
            System.out.println("Text: ");
            System.out.println("Language: " + ocrResult.language());  
            System.out.println("Text angle: " + ocrResult.textAngle());      
            System.out.println("Orientation: " + ocrResult.orientation());      

            System.out.println("Text regions: ");
            for (OcrRegion reg : ocrResult.regions()) {
                System.out.println("Region bounding box: " + reg.boundingBox());
                
                for (OcrLine line : reg.lines()) {
                    System.out.println("Line bounding box: " + line.boundingBox());
                    
                    for (OcrWord word : line.words()) {
                        System.out.println("Word bounding box: " + word.boundingBox());
                        System.out.println("Text: " + word.text() + " ");
                    }

                    System.out.println();                    
                
                }

                System.out.println();
            }
        }
    }
}
