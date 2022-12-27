package red.tetracube.cli;

import java.util.concurrent.Callable;

import picocli.CommandLine.Command;
import picocli.CommandLine.Option;
import picocli.CommandLine.ScopeType;

@Command(
    name = "TetraCube Maintenance CLI",
    version = "4.0", 
    description = "Helps to manage users and other aspects of your TetraCube house.dd",
    subcommands = HouseServices.class
    )
public class ServiceCli implements Callable<Integer> {

    @Option(names = {"--hostname"}, description = "TetraCube's hostname", required = true, scope = ScopeType.INHERIT) 
    public String hostname;

    @Override public Integer call() {
        return 0;
    }
}