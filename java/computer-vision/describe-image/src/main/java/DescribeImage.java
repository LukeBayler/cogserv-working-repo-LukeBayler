import com.microsoft.azure.cognitiveservices.vision.computervision.*;
import com.microsoft.azure.cognitiveservices.vision.computervision.models.*;

import java.io.File;
import java.nio.file.Files;

import java.util.ArrayList;
import java.util.List;

public class DescribeImage {

    public static void main(String[] args) {
        String AzureBaseURL = "https://westus.api.cognitive.microsoft.com";
        String CMSubscriptionKey = "bed9632798b9496bab97d18e31d0fde9";
        ComputerVisionClient compVisClient = ComputerVisionManager.authenticate(CMSubscriptionKey).withEndpoint(AzureBaseURL);
        System.out.println("compVisClient.endpoint(): " + compVisClient.endpoint());
        
        System.out.println("Describing local image...");
        
        String imagePath = "C:\\Users\\v-lubayl\\Documents\\GitHub\\cognitive-services-samples\\java\\computer-vision\\describe-image\\src\\main\\resources\\upside-down-mushroom.jpg";
        File rawImage = new File(imagePath);
        
        try {
            byte[] imageBytes = Files.readAllBytes(rawImage.toPath());
            
            List<VisualFeatureTypes> visualFeatureTypes = new ArrayList<>();
            visualFeatureTypes.add(VisualFeatureTypes.DESCRIPTION);
            ImageAnalysis imgAnalysis = compVisClient.computerVision().analyzeImageInStream().withImage(imageBytes).withVisualFeatures(visualFeatureTypes).execute();
    
            System.out.print("Description: ");           
            System.out.println(imgAnalysis.description().captions().get(0).text());

        } catch (Exception e) {
            System.out.println(e.getMessage());
            e.printStackTrace();
        }
    }
}
