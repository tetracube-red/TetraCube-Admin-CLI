package red.tetracube;

import picocli.CommandLine;
import red.tetracube.cli.*;

public class App 
{
    public static void main( String[] args )
    {
        int exitCode = new CommandLine(new ServiceCli()).execute(args);
        System.exit(exitCode);
    }
}
