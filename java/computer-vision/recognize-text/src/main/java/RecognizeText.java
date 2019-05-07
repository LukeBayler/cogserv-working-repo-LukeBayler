import com.microsoft.azure.cognitiveservices.vision.computervision.*;
import com.microsoft.azure.cognitiveservices.vision.computervision.models.*;

import java.io.File;
import java.nio.file.Files;

import java.util.ArrayList;
import java.util.List;

public class RecognizeText {

    public static void main(String[] args) {
        String AzureBaseURL = "https://westus.api.cognitive.microsoft.com";
        String CVSubscriptionKey = "bed9632798b9496bab97d18e31d0fde9";
        ComputerVisionClient compVisClient = ComputerVisionManager.authenticate(CVSubscriptionKey).withEndpoint(AzureBaseURL);
        System.out.println("compVisClient.endpoint(): " + compVisClient.endpoint());
                
        String imagePath = "C:\\samples\\files\\printed-and-written-text.png";
        File rawImage = new File(imagePath);
        
        /*
         *  Notes here for how to use the service.
         *  1.  need to set text recogition mode to either handwritten or printed, there is an enum that does this
         */
        
        
        
        try {
            byte[] imageBytes = Files.readAllBytes(rawImage.toPath());
            
            List<VisualFeatureTypes> visualFeatureTypes = new ArrayList<>();
            visualFeatureTypes.add(VisualFeatureTypes.TAGS);
            //ImageAnalysis imgAnalysis = compVisClient.computerVision().recognizePrintedTextInStream().withDetectOrientation(false).withImage(imageBytes).withLanguage(OcrLanguages.EN).execute();
    
            System.out.print("Description: ");      
//            for (ImageTag tag : imgAnalysis.tags()) {
  //              System.out.println(String.format("%s\t\t%s", tag.name(), tag.confidence()));
    //        }
                 
        } catch (Exception e) {
            System.out.println(e.getMessage());
            e.printStackTrace();
        }
    }
}