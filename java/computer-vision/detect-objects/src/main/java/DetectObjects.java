import com.microsoft.azure.cognitiveservices.vision.computervision.*;
import com.microsoft.azure.cognitiveservices.vision.computervision.models.*;

import java.io.File;
import java.io.FileInputStream;
import java.nio.file.Files;

import java.util.ArrayList;
import java.util.List;

public class DetectObjects {
    public static String subKey = System.getenv("AZURE_COMPUTERVISION_API_KEY");
    public static String baseURL = System.getenv("AZURE_ENDPOINT");
    
    public static void main(String[] args) {
        try {
            DetectObjectsSample.RunSample(baseURL, subKey);
            
        } catch (Exception e) {
            System.out.println(e.getMessage());
            e.printStackTrace();
        }
    }
    
    private static class DetectObjectsSample {
        public static void RunSample(String url, String key) {
        
            ComputerVisionClient compVisClient = ComputerVisionManager.authenticate(key).withEndpoint(url);
        
            String imgPath = "src\\main\\resources\\analyze-image.jpg";
            String remotePath = "https://github.com/Azure-Samples/cognitive-services-sample-data-files/raw/master/ComputerVision/Images/landmark.jpg";
        
/*            List<VisualFeatureTypes> features = new ArrayList<>();
            features.add(VisualFeatureTypes.DESCRIPTION);
            features.add(VisualFeatureTypes.CATEGORIES);
            features.add(VisualFeatureTypes.TAGS);
            features.add(VisualFeatureTypes.FACES);
            features.add(VisualFeatureTypes.ADULT);
            features.add(VisualFeatureTypes.COLOR);
            features.add(VisualFeatureTypes.IMAGE_TYPE);
         
            System.out.println("\nAnalyzing local image ...");
            DetectObjectsLocal(compVisClient, imgPath, features);
            
            System.out.println("\nAnalyzing image from URL ...");
            DetectObjectsFromUrl(compVisClient, remotePath, features);
*/      
            System.out.println("\nAnalyzing local image ...");
            DetectObjectsLocal(compVisClient, imgPath);
            
            System.out.println("\nAnalyzing image from URL ...");
            DetectObjectsFromUrl(compVisClient, remotePath);
        }
        
        private static void DetectObjectsLocal(ComputerVisionClient client, String path) {
            try {
                File rawImg = new File(path);
                byte[] imgBytes = Files.readAllBytes(rawImg.toPath());
                
                ImageAnalysis analysis = client.computerVision().analyzeImageInStream()
                    .withImage(imgBytes)
                    .execute();
                
                DisplayResults(analysis, path);
                
            } catch (Exception e) {
                System.out.println(e.getMessage());
                e.printStackTrace();
            }
        }
        
        private static void DetectObjectsFromUrl(ComputerVisionClient client, String path) {
            try {
                ImageAnalysis analysis = client.computerVision().analyzeImage()
                    .withUrl(path)
                    .execute();
                
                DisplayResults(analysis, path);

            } catch (Exception e) {
                System.out.println(e.getMessage());
                e.printStackTrace();
            }
        }
        
        private static void DisplayResults(ImageAnalysis analysis, String path) {
            DisplayObjects(analysis);
        }
        
        private static void DisplayObjects(ImageAnalysis analysis) {
            
        }
    }        
}
