#!/bin/bash
# bash script to migrate containers from one backend storage to another.


main() {
if [ "$USER" != "root" ];then
	echo "Run 'containers-migrate' as root user"
	exit 
fi

if [ -z "$1" ];then
	echo "See 'containers-migrate --help'"
	exit
fi

if [ "$1" = "--help" ];then 
	echo "Usage: containers-migrate COMMAND [OPTIONS]"
	echo -e "       containers-migrate [--help]\n"
	echo -e "A self-sufficient tool for migrating docker containers from one backend storage to another\n"
	echo "Commands:"
	echo "    export  Export a container from an existing storage"
	echo "    import  Import a container into a new storage" 
	exit
fi

if [ "$1" = "export" ];then
   if [ -z "$2" ]; then
	echo "containers-migrate: "export" requires a minimum of 1 argument." 
	echo -e "See 'containers-migrate export --help'\n"
	echo -e "Usage: containers-migrate export [OPTIONS]\n"
	echo "Export a container from an existing storage"
	exit
   elif [ "$2" = "--help" ];then
	echo -e "\nUsage: containers-migrate export [OPTIONS]\n"
	echo -e "Export a container from an existing storage\n"
	echo "--container-id   ID of the container to be exported." 
	echo "--graph   	 Root of the Docker runtime."
	exit
   else
	container_export $2 $3
	echo "Container exported succesfully"
   fi
fi

if [ "$1" = "import" ];then
   if [ -z "$2" ]; then
        echo "containers-migrate: "import" requires a minimum of 1 argument." 
        echo -e "See 'containers-migrate import --help'\n"
        echo -e "Usage: containers-migrate import [OPTIONS]\n"
        echo "Import a container into a new storage"
        exit
   elif [ "$2" = "--help" ];then
        echo -e "\nUsage: containers-migrate import [OPTIONS]\n"
        echo -e "Import a container into a new storage\n"
        echo "--container-id   ID of the container to be imported." 
        echo "--graph          Root of the Docker runtime."
        exit
   else
        container_import $2 $3
        echo "Container imported succesfully"
   fi
fi

}

container_export(){
	containerID="$1"
        dockerPid=$(ps aux|grep docker|awk 'NR==1{print $2}')
        dockerCmdline=$(cat /proc/$dockerPid/cmdline)
        if [[ $dockerCmdline =~ "-g=" ]] || [[ $dockerCmdline =~ "-g/" ]] || [[ $dockerCmdline =~ "--graph" ]];then
                if [ -z "$2" ];then
                        echo "Docker is not located at the default (/var/lib/docker) root location."
                        echo "Please provide the new root location of the docker runtime in --graph option."
                else
                        dockerRootDir="$2"
                fi
        else
                dockerRootDir="/var/lib/docker"
        fi
        notruncContainerID=$(sudo docker ps -aq --no-trunc|grep $containerID)
        tmpDir=$dockerRootDir/tmp/docker-migrate-$containerID
        mkdir $tmpDir
        cd $tmpDir
	echo $dockerRootDir>dockerRootDir.txt
        tar -cf container-metadata.tar $dockerRootDir/containers/$notruncContainerID 2> /dev/null
        imageID=$(docker commit $containerID)
        mkdir $tmpDir/temp
        docker save $imageID > $tmpDir/temp/image.tar
        cd $tmpDir/temp
        tar -xf image.tar
        cd $tmpDir/temp/$imageID
        cp layer.tar $tmpDir/container-diff.tar
        cd $tmpDir
        rm -rf temp
        docker rmi -f $imageID 1>/dev/null
}

container_import(){
	echo "Container imported successfully"
}

main "$@"
