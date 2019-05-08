import com.microsoft.azure.cognitiveservices.vision.computervision.*;
import com.microsoft.azure.cognitiveservices.vision.computervision.models.*;

import java.io.File;
import java.nio.file.Files;

import java.util.ArrayList;
import java.util.List;

public class TagVisualFeatures {

    public static void main(String[] args) {
    
        String subKey = System.getenv("AZURE_COMPUTERVISION_API_KEY");
        String baseURL = System.getenv("AZURE_ENDPOINT");
 
        ComputerVisionClient compVisClient = ComputerVisionManager.authenticate(subKey).withEndpoint(baseURL);
        System.out.println("compVisClient.endpoint(): " + compVisClient.endpoint());
                
        String imagePath = "src\\main\\resources\\tag-visual-features.jpg";
        File rawImage = new File(imagePath);
        
        try {
            byte[] imageBytes = Files.readAllBytes(rawImage.toPath());
            
            List<VisualFeatureTypes> visualFeatureTypes = new ArrayList<>();
            visualFeatureTypes.add(VisualFeatureTypes.TAGS);
            ImageAnalysis imgAnalysis = compVisClient.computerVision().analyzeImageInStream().withImage(imageBytes).withVisualFeatures(visualFeatureTypes).execute();
    
            System.out.println("Tags\t\tConfidence");      
            for (ImageTag tag : imgAnalysis.tags()) {
                System.out.println(String.format("%s\t\t%s", tag.name(), tag.confidence()));
            }
                 
        } catch (Exception e) {
            System.out.println(e.getMessage());
            e.printStackTrace();
        }
    }
}
