import com.microsoft.azure.cognitiveservices.vision.computervision.*;
import com.microsoft.azure.cognitiveservices.vision.computervision.models.*;

import java.io.File;
import java.io.FileInputStream;
import java.nio.file.Files;

import java.util.ArrayList;
import java.util.List;

public class AnalyzeImage {
    public static String subKey = System.getenv("AZURE_COMPUTERVISION_API_KEY");
    public static String baseURL = System.getenv("AZURE_ENDPOINT");
    
    public static void main(String[] args) {
        try {
            AnalyzeImageSample.RunSample(baseURL, subKey);
            
        } catch (Exception e) {
            System.out.println(e.getMessage());
            e.printStackTrace();
        }
    }
    
    private static class AnalyzeImageSample {
        public static void RunSample(String url, String key) {
        
            ComputerVisionClient compVisClient = ComputerVisionManager.authenticate(key).withEndpoint(url);
        
            String imgPath = "src\\main\\resources\\analyze-image.jpg";
            String remotePath = "https://github.com/Azure-Samples/cognitive-services-sample-data-files/raw/master/ComputerVision/Images/landmark.jpg";
        
            List<VisualFeatureTypes> features = new ArrayList<>();
            features.add(VisualFeatureTypes.DESCRIPTION);
            features.add(VisualFeatureTypes.CATEGORIES);
            features.add(VisualFeatureTypes.TAGS);
            features.add(VisualFeatureTypes.FACES);
            features.add(VisualFeatureTypes.ADULT);
            features.add(VisualFeatureTypes.COLOR);
            features.add(VisualFeatureTypes.IMAGE_TYPE);
            
            System.out.println("Analyzing image ...");
            AnalyzeLocal(compVisClient, imgPath, features);
            AnalyzeFromUrl(compVisClient, remotePath, features);
        }
        
        private static void AnalyzeLocal(ComputerVisionClient client, String path, List<VisualFeatureTypes> feats) {
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
        
        private static void AnalyzeFromUrl(ComputerVisionClient client, String path, List<VisualFeatureTypes> feats) {
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
            DisplayImageDescription(analysis);
            DisplayImageCategoryResults(analysis);
            DisplayTagResults(analysis);
            DisplayFaceResults(analysis);
            DisplayAdultResults(analysis);
            DisplayColorSchemeResults(analysis);
            DisplayDomainSpecificResults(analysis);
            DisplayImageTypeResults(analysis);
        }
        
        private static void DisplayImageDescription(ImageAnalysis analysis) {
            System.out.println("\nCaptions: ");
            for (ImageCaption caption : analysis.description().captions()) {
                System.out.printf("\'%s\' with confidence %f\n", caption.text(), caption.confidence());
            }
        }
        
        private static void DisplayImageCategoryResults(ImageAnalysis analysis) {
            System.out.println("\nCategories: ");
            for (Category category : analysis.categories()) {
                System.out.printf("\'%s\' with confidence %f\n", category.name(), category.score());
            }
        }
        
        private static void DisplayTagResults(ImageAnalysis analysis) {
            System.out.println("\nTags: ");
            for (ImageTag tag : analysis.tags()) {
                System.out.printf("\'%s\' with confidence %f\n", tag.name(), tag.confidence());
            }
        }

        private static void DisplayFaceResults(ImageAnalysis analysis) {
            System.out.println("\nFaces: ");
            for (FaceDescription face : analysis.faces()) {
                System.out.printf("\'%s\' of age %d at location (%d, %d), (%d, %d)\n", face.gender(), face.age(),
                    face.faceRectangle().left(), face.faceRectangle().top(),
                    face.faceRectangle().left() + face.faceRectangle().width(),
                    face.faceRectangle().top() + face.faceRectangle().height());
            }
        }

       private static void DisplayAdultResults(ImageAnalysis analysis) {
            System.out.println("\nAdult: ");
            System.out.printf("Is adult content: %b with confidence %f\n", analysis.adult().isAdultContent(), analysis.adult().adultScore());
            System.out.printf("Has racy content: %b with confidence %f\n\n", analysis.adult().isRacyContent(), analysis.adult().racyScore());
        }
        
        private static void DisplayColorSchemeResults(ImageAnalysis analysis) {
            System.out.println("\nColor scheme: ");
            System.out.println("Is black and white: " + analysis.color().isBWImg());
            System.out.println("Accent color: " + analysis.color().accentColor());
            System.out.println("Dominant background color: " + analysis.color().dominantColorBackground());
            System.out.println("Dominant foreground color: " + analysis.color().dominantColorForeground());
            System.out.println("Dominant colors: " + String.join(", ", analysis.color().dominantColors()));
        }
        
        private static void DisplayDomainSpecificResults(ImageAnalysis analysis)
        {
            System.out.println("\nCelebrities: ");
            for (Category category : analysis.categories())
            {
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
        
        private static void DisplayImageTypeResults(ImageAnalysis analysis) {
            System.out.println("\nImage type:");
            System.out.println("Clip art type: " + analysis.imageType().clipArtType());
            System.out.println("Line drawing type: " + analysis.imageType().lineDrawingType());
        }
    }
}
