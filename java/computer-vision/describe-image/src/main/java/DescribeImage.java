import com.microsoft.azure.cognitiveservices.vision.computervision.*;
import com.microsoft.azure.cognitiveservices.vision.computervision.models.*;

import java.io.File;
import java.nio.file.Files;

import java.util.ArrayList;
import java.util.List;

public class DescribeImage {

    public static void main(String[] args) {  

        //  First command-line argument is the subscription key.
        String subKey = args[0];

        String baseURL = "https://westus.api.cognitive.microsoft.com";
        ComputerVisionClient compVisClient = ComputerVisionManager.authenticate(subKey).withEndpoint(baseURL);
        System.out.println("compVisClient.endpoint(): " + compVisClient.endpoint());
        
        //System.out.println("Describing local image...");
        
        //  Path relative to the current working directory (user.dir).
        //System.out.println(System.getProperty("user.dir"));
        String imgPath = "src\\main\\resources\\upside-down-mushroom.jpg";
        
        File rawImg = new File(imgPath);
        
        try {
            byte[] imgBytes = Files.readAllBytes(rawImg.toPath());    

            List<VisualFeatureTypes> visualFeatureTypes = new ArrayList<VisualFeatureTypes>();
            visualFeatureTypes.add(VisualFeatureTypes.DESCRIPTION);
            ImageAnalysis imgAnalysis = compVisClient.computerVision().analyzeImageInStream().withImage(imgBytes).withVisualFeatures(visualFeatureTypes).execute();
    
            System.out.println("\nDescription: ");           
            System.out.println("\t" + imgAnalysis.description().captions().get(0).text() + "\n");

        } catch (Exception e) {
            System.out.println(e.getMessage());
            e.printStackTrace();
        }
    }
}
