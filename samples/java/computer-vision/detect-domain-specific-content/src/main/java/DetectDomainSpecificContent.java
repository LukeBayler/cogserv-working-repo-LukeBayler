import com.microsoft.azure.cognitiveservices.vision.computervision.*;
import com.microsoft.azure.cognitiveservices.vision.computervision.models.*;

import java.io.File;
import java.io.FileInputStream;
import java.nio.file.Files;

import java.util.ArrayList;
import java.util.List;

public class DetectDomainSpecificContent {
    public static String subKey = System.getenv("AZURE_COMPUTERVISION_API_KEY");
    public static String baseURL = System.getenv("AZURE_ENDPOINT");
    
    public static void main(String[] args) {
        try {
            DetectDomainSpecificContentSample.RunSample(baseURL, subKey);
            
        } catch (Exception e) {
            System.out.println(e.getMessage());
            e.printStackTrace();
        }
    }
    
    private static class DetectDomainSpecificContentSample {
        public static void RunSample(String url, String key) {
        
            ComputerVisionClient compVisClient = ComputerVisionManager.authenticate(key).withEndpoint(url);
        
            String imgPath = "src\\main\\resources\\analyze-image.jpg";
            String remotePath = "https://github.com/Azure-Samples/cognitive-services-sample-data-files/raw/master/ComputerVision/Images/landmark.jpg";
        
            List<VisualFeatureTypes> features = new ArrayList<>();
            features.add(VisualFeatureTypes.DESCRIPTION);
            features.add(VisualFeatureTypes.CATEGORIES);
            
            System.out.println("\nAnalyzing local image ...");
            DetectDomainSpecificContentLocal(compVisClient, imgPath, features);
            
            System.out.println("\nAnalyzing image from URL ...");
            DetectDomainSpecificContentFromUrl(compVisClient, remotePath, features);
        }
        
        private static void DetectDomainSpecificContentLocal(ComputerVisionClient client, String path, List<VisualFeatureTypes> feats) {
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
        
        private static void DetectDomainSpecificContentFromUrl(ComputerVisionClient client, String path, List<VisualFeatureTypes> feats) {
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
            DisplayDomainSpecificResults(analysis);
        }
 
        private static void DisplayDomainSpecificResults(ImageAnalysis analysis)
        {
            System.out.println("\nCelebrities: ");
            for (Category category : analysis.categories())
            {
                //System.out.println("category: " + category);
                //System.out.println("category.detail(): " + category.detail());

                if (category.detail() != null && category.detail().celebrities() != null)
                {
                    for (CelebritiesModel celeb : category.detail().celebrities())
                    {
                        System.out.printf("\'%s\' with confidence %f at location (%d, %d), (%d, %d)\n", celeb.name(), celeb.confidence(),
                            celeb.faceRectangle().left(), celeb.faceRectangle().top(),
                            celeb.faceRectangle().left() + celeb.faceRectangle().width(),
                            celeb.faceRectangle().top() + celeb.faceRectangle().height());
                    }
                }
            }
            
            System.out.println("\nLandmarks: ");
            for (Category category : analysis.categories())
            {
                if (category.detail() != null && category.detail().landmarks() != null)
                {
                    for (LandmarksModel landmark : category.detail().landmarks())
                    {
                        System.out.printf("\'%s\' with confidence %f\n", landmark.name(), landmark.confidence());
                    }
                }
            }
        }
    }
}
