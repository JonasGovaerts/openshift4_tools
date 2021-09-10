package main

import (
	//"fmt"
	"os/exec"
	log "github.com/sirupsen/logrus"
	"flag"
)



func pruneIndex(operatorIndex string, packages string, localOperatorIndex string){
	log.Info("Pruning index: ", operatorIndex)
	app := "opm" 
	arg0 := "index" 
	arg1 := "prune"
	arg2 := "-f"
	arg3 := operatorIndex
	arg4 := "-p"
	arg5 := packages
	arg6 := "-t"
	arg7 := localOperatorIndex

	cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)

        stdout,stderr := cmd.CombinedOutput()

	if stderr != nil {
		log.Error("Failed to prune index ",operatorIndex," error_mesage: ",string(stdout))
		log.Debug("command executed: ",app," ", arg0," ", arg1," ", arg2," ", arg3," ", arg4," ", arg5," ", arg6," ", arg7)
	} else {
		log.Info("Successfully pruned index ",operatorIndex, "to ",localOperatorIndex)
		log.Debug("Stdout output: ",string(stdout))
	}
}

func listOperatorPackages(operatorIndex string, credentialsFile string){
	log.Info("Listing available operators in: ",operatorIndex)
	app := 	"podman"
	arg0 :=	"run"
	arg1 := "--authfile"
	arg2 := credentialsFile
	arg3 := "--name"
	arg4 := "temp-operator-index"
	arg5 := "-p"
	arg6 := "50051:50051"
	arg7 := "-it"
	arg8 := "-d"
	arg9 := operatorIndex

	cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)

	stdout,stderr := cmd.CombinedOutput()

	if stderr != nil {
		log.Error("Failed to run container image ",operatorIndex," error_mesage: ",string(stdout))
		log.Debug("command executed: ",app," ", arg0," ", arg1," ", arg2," ", arg3," ", arg4," ", arg5," ", arg6," ", arg7)
	} else {
		log.Info("Successfully started ",operatorIndex)
		log.Debug("Stdout output: ",string(stdout))
	}

	app1 	:= "grpcurl" 
	arg1_0 	:= "-plaintext"
	arg1_1 	:= "localhost:50051"
	arg1_2 	:= "api.Registry/ListPackages"

	cmd1 := exec.Command(app1, arg1_0, arg1_1, arg1_2)
	stdout1, stderr1 := cmd1.CombinedOutput()

	if stderr1 != nil {
		log.Error("Failed to get packages from ",operatorIndex," error_mesage: ",string(stdout1))
		log.Debug("command executed: ",app1," ", arg1_0," ", arg1_1," ", arg1_2," ")
	} else {
		log.Info("Packages available from  ",operatorIndex, ": ",string(stdout1))
		log.Debug("Stdout output: ",string(stdout1))
	}

	app2 :=  "podman"
        arg2_0 := "rm"
        arg2_1 := "-f"
        arg2_2 := "temp-operator-index"

        cmd2 := exec.Command(app2, arg2_0, arg2_1, arg2_2)

        stdout2,stderr2 := cmd2.CombinedOutput()

        if stderr2 != nil {
                log.Error("Failed to delete local pod temp-operator-index error_mesage: ",string(stdout2))
                log.Debug("command executed: ",app2," ", arg2_0," ", arg2_1," ", arg2_2)
        } else {
                log.Info("Successfully deleted local pod temp-operator-index")
                log.Debug("Stdout output: ",string(stdout))
        }

}

func setLogLevel(logLevel string) {
	log.SetFormatter(&log.JSONFormatter{})
	switch logLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
		log.Info("Set loglevel to debug")
	}
}

func main() {
	listPackages            := flag.Bool("list", false, "list packages available in the operator index")
        operatorIndex           := flag.String("operator", "", "operator index that you want to mirror")
        packages                := flag.String("packages", "", "comma seperated list of packages that need to be mirrored")
        credentialsFile         := flag.String("creds", "", "location to authentication file")
        localOperatorIndex      := flag.String("local-operator", "", "target to push the modified operator index to")
	logLevel		:= flag.String("loglevel", "debug" , "set log level: debug, info, warn")
        flag.Parse()

	setLogLevel(*logLevel)

	//debug log dumping passed variables
	log.Debug("Upstream operator index image: ",*operatorIndex)
	log.Debug("Local operator index target: ",*localOperatorIndex)
	log.Debug("packages: ",*packages)
	log.Debug("location of authentication file: ",*credentialsFile)

	if *listPackages {
		listOperatorPackages(*operatorIndex,*credentialsFile)
	} else {
		pruneIndex(*operatorIndex,*packages,*localOperatorIndex)
	}
}
