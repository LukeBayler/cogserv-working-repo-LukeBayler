import com.microsoft.azure.cognitiveservices.vision.computervision.*;
import com.microsoft.azure.cognitiveservices.vision.computervision.models.*;

import java.io.File;
import java.nio.file.Files;

import java.util.ArrayList;
import java.util.List;

public class RecognizePrintedText {

    public static void main(String[] args) {
        public static String subKey = System.getenv("AZURE_COMPUTERVISION_API_KEY");
        public static String baseURL = System.getenv("AZURE_ENDPOINT");
        ComputerVisionClient compVisClient = ComputerVisionManager.authenticate(CVSubscriptionKey).withEndpoint(AzureBaseURL);
        //System.out.println("compVisClient.endpoint(): " + compVisClient.endpoint());
                
        String imagePath = "C:\\Users\\v-lubayl\\Documents\\GitHub\\cognitive-services-samples\\java\\computer-vision\\recognize-printed-text\\src\\main\\resources\\ocr-sample.png";
        File rawImage = new File(imagePath);
                
        try {
            byte[] imageBytes = Files.readAllBytes(rawImage.toPath());
            OcrResult ocrResult = compVisClient.computerVision().recognizePrintedTextInStream().withDetectOrientation(true).withImage(imageBytes).withLanguage(OcrLanguages.EN).execute();
             
            System.out.println("Language: " + ocrResult.language());      
            System.out.println("Orientation: " + ocrResult.orientation());      
            System.out.println("Text angle: " + ocrResult.textAngle());      
            System.out.println();

            for (OcrRegion reg : ocrResult.regions()) {
                //System.out.println("Region: " + reg);
                //System.out.println("Bounding box: " + reg.boundingBox());
                
                //System.out.println("Lines: ");
                for (OcrLine line : reg.lines()) {
                    //System.out.println(line);
                    
                    System.out.println("Words: ");
                    for (OcrWord word : line.words()) {
                        System.out.print(word.text() + " ");
                    }
                    System.out.println();                    
                    System.out.println();                    
                
                }

                System.out.println();
            }
                 
        } catch (Exception e) {
            System.out.println(e.getMessage());
            e.printStackTrace();
        }
    }
}
