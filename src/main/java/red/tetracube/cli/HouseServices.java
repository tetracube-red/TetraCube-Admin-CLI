package red.tetracube.cli;

import java.io.IOException;
import java.net.URI;
import java.net.URISyntaxException;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.net.http.HttpResponse.BodyHandlers;
import java.util.HashMap;
import java.util.concurrent.Callable;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;

import picocli.CommandLine.Command;
import picocli.CommandLine.Option;
import picocli.CommandLine.ParseResult;
import picocli.CommandLine.Spec;
import picocli.CommandLine.Model.ArgSpec;
import picocli.CommandLine.Model.CommandSpec;
import picocli.CommandLine.Model.OptionSpec;

@Command(
    name = "house", 
    description = "Manage TetraCube houses."
)
public class HouseServices implements Callable<Integer> {

    @Option(names = {"-h", "--help"}, usageHelp = true)
    private boolean help;

    @Spec
    private CommandSpec spec;

    private final ObjectMapper objectMapper;
    private final HttpClient httpClient;

    public HouseServices() {
        this.objectMapper = new ObjectMapper();
        this.httpClient = HttpClient.newHttpClient();
    }

    @Override
    public Integer call() {
        System.out.println("Subcommand needed: 'create'");
        return 0;
    }

    @Command(name = "create", helpCommand = true, description = "Creates house in the brand new TetraCube installation.")
    public int create(
        @Option(names = {"-h", "--help"}, usageHelp = true) boolean help,
        @Option(names = {"-n", "--name"}, description = "House name", required = true) String houseName
    ) {
        System.out.println("Creating house of name " + houseName);
        var hostName = this.spec.options().get(1).getValue();
        URI uri;
        try {
            uri = new URI(String.format("http://%s:8080/houses", hostName));
        } catch (URISyntaxException e) {
            System.out.println("Failed to create TetraCube's URL");
            return -1;
        }

        var requestBody = new HashMap<String, Object>();
        requestBody.put("name", houseName);
        String requestBodySerialized;
        try {
            requestBodySerialized = this.objectMapper.writeValueAsString(requestBody);
        } catch (JsonProcessingException e) {
            System.out.println("Cannot build house creation request");
            return -1;
        }

        var request = HttpRequest.newBuilder()
                .uri(uri)
                .POST(HttpRequest.BodyPublishers.ofString(requestBodySerialized))
                .build();
        HttpResponse<String> response;
        try {
            response = this.httpClient.send(
                request, 
                BodyHandlers.ofString()
            );
        } catch (IOException | InterruptedException e) {
            System.out.println("Cannot send request.");
            return -1;
        }

        if (response.statusCode() != 200) {
            System.out.println("Invalid response.");
            return -1;
        }
        
        System.out.println("House created");
        return 0;
    }   
}

/*
 * fullApi := fmt.Sprintf("http://%s/houses", host)
 * bodyMap := map[string]string{"name": houseName}
 * jsonValue, _ := json.Marshal(bodyMap)
 * 
 * resp, err := http.Post(fullApi, "application/json",
 * bytes.NewBuffer(jsonValue))
 * if err != nil {
 * fmt.Printf("Error in house creation %s\n", err.Error())
 * }
 * if resp.StatusCode == 409 {
 * fmt.Printf("House %s already exists \n", houseName)
 * return
 * }
 * 
 * body, _ := io.ReadAll(resp.Body)
 * var responseMap map[string]string
 * _ = json.Unmarshal(body, &responseMap)
 * fmt.Printf("House %s created with id %s", responseMap["name"],
 * responseMap["id"])
 */