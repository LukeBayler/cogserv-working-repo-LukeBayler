import com.microsoft.azure.cognitiveservices.vision.computervision.*;
import com.microsoft.azure.cognitiveservices.vision.computervision.models.*;

import java.io.File;
import java.io.FileInputStream;
import java.nio.file.Files;

import java.util.ArrayList;
import java.util.List;

public class TagVisualFeatures {
    public static String subKey = System.getenv("AZURE_COMPUTERVISION_API_KEY");
    public static String baseURL = System.getenv("AZURE_ENDPOINT");
    
    public static void main(String[] args) {
        try {
            TagVisualFeaturesSample.RunSample(baseURL, subKey);
            
        } catch (Exception e) {
            System.out.println(e.getMessage());
            e.printStackTrace();
        }
    }
    
    private static class TagVisualFeaturesSample {
        public static void RunSample(String url, String key) {
        
            ComputerVisionClient compVisClient = ComputerVisionManager.authenticate(key).withEndpoint(url);
        
            String imgPath = "src\\main\\resources\\analyze-image.jpg";
            String remotePath = "https://github.com/Azure-Samples/cognitive-services-sample-data-files/raw/master/ComputerVision/Images/landmark.jpg";
        
            List<VisualFeatureTypes> features = new ArrayList<>();
            features.add(VisualFeatureTypes.TAGS);
  
            System.out.println("\nAnalyzing local image ...");
            TagVisualFeaturesLocal(compVisClient, imgPath, features);
            
            System.out.println("\nAnalyzing image from URL ...");
            TagVisualFeaturesFromUrl(compVisClient, remotePath, features);
        }
        
        private static void TagVisualFeaturesLocal(ComputerVisionClient client, String path, List<VisualFeatureTypes> feats) {
            try {
                File rawImg = new File(path);
                byte[] imgBytes = Files.readAllBytes(rawImg.toPath());
                
                ImageAnalysis analysis = client.computerVision().analyzeImageInStream()
                    .withImage(imgBytes)
                    .withVisualFeatures(feats)
                    .execute();
                
                DisplayResults(analysis, path);
                
            } catch (Exception e) {
                System.out.println(e.getMessage());
                e.printStackTrace();
            }
        }
        
        private static void TagVisualFeaturesFromUrl(ComputerVisionClient client, String path, List<VisualFeatureTypes> feats) {
            try {
                ImageAnalysis analysis = client.computerVision().analyzeImage()
                    .withUrl(path)
                    .withVisualFeatures(feats)
                    .execute();
                
                DisplayResults(analysis, path);

            } catch (Exception e) {
                System.out.println(e.getMessage());
                e.printStackTrace();
            }
        }
        
        private static void DisplayResults(ImageAnalysis analysis, String path) {
            DisplayTagResults(analysis);
        }
        
        private static void DisplayTagResults(ImageAnalysis analysis) {
            System.out.println("\nTags: ");
            for (ImageTag tag : analysis.tags()) {
                System.out.printf("\'%s\' with confidence %f\n", tag.name(), tag.confidence());
            }
        }
    }
}
