#!/usr/bin/env python

import sys
import os
import subprocess


def main():
	if len(sys.argv) == 1:
		show_help()
	elif len(sys.argv) == 2 and sys.argv[1] == "help":
		show_help()

	if os.geteuid() != 0:
                exit("You need to have root privileges to run this script.\nPlease try again, this time using 'sudo'. Exiting.")

	if len(sys.argv) == 2 and sys.argv[1] == "export":
		exportDir = "/var/lib/docker-migrate"
		export_docker(exportDir)		
	elif len(sys.argv) == 3 and sys.argv[1] == "export":
		exportDir = sys.argv[2]
		export_docker(exportDir)
	elif len(sys.argv) == 2 and sys.argv[1] == "import":
                importDir = "/var/lib/docker-migrate"
                import_docker(importDir)
        elif len(sys.argv) == 3 and sys.argv[1] == "import":
                importDir = sys.argv[2]
                import_docker(importDir)	
		
def export_docker(exportDir):
	if not os.path.isdir(exportDir):
		os.mkdir(exportDir)
	#export docker images
	export_images(exportDir)
	#export docker containers
	export_containers(exportDir)
	#export docker volumes
	export_volumes(exportDir)
	print("docker export completed successfully")

def import_docker(importDir):
        if not os.path.isdir(importDir):
        	sys.exit("Specified directory {0} does not exist".format(importDir))
	#import docker images
        import_images(importDir)
        #import docker containers
        import_containers(importDir)
        #import docker volumes
        import_volumes(importDir)
	print("docker import completed successfully")

def export_images(exportDir):
	if not os.path.isdir(exportDir + "/images"):
        	os.mkdir(exportDir + "/images")
	images = subprocess.check_output("docker images", shell=True)
        splitImages = images.split()[7:]  # cut off the headers
        names = []
        tags = []
        for i in range(0, len(splitImages)):
            # only take the image and its tags and the image ID (to help in the <none>:<none> case)
            if (i % 8 == 0):
                names.append(splitImages[i])
                tags.append(splitImages[i+1])
        for i in range(0, len(names)):
            print("Saving image {0}:{1}".format(names[i], tags[i]))
            if names[i] == '<none>':
            	print("This is a dangling image and will not be exported")
	    else:
                subprocess.call(
                    "docker save {0}:{1} > {2}/images/{3}-{4}.tar".format(
                      	names[i], tags[i], exportDir, names[i].replace("/", "~"), tags[i].replace("/", "~")), shell=True)

def export_containers(exportDir):
	if not os.path.isdir(exportDir + "/containers"):
            os.mkdir(exportDir + "/containers")

def export_volumes(exportDir):
	if not os.path.isdir(exportDir + "/volumes"):
            os.mkdir(exportDir + "/volumes")
	subprocess.call(
            "tar -zcvf {0}/volumes/volumeData.tar.gz -C /var/lib/docker/volumes . > /dev/null".format(exportDir), shell=True)
        if os.path.isdir("/var/lib/docker/vfs"):
            subprocess.call("tar -zcvf {0}/volumes/vfsData.tar.gz -C /var/lib/docker/vfs . > /dev/null".format(exportDir), shell=True)

def import_images(importDir):
	tarballs = subprocess.check_output("ls {0}/images".format(importDir), shell=True)
        splitTarballs = tarballs.split()
        for i in splitTarballs:
            print("Loading image {0}".format(i))
            subprocess.call("docker load < {0}/images/{1}".format(importDir, i), shell=True)

def import_containers(importDir):
	print("import_containers")

def import_volumes(importDir):
	subprocess.call(
            "tar xzvf {0}/volumes/volumeData.tar.gz -C /var/lib/docker/volumes > /dev/null".format(importDir), shell=True)
        if os.path.isdir("/var/lib/docker/vfs"):
            subprocess.call(
                "tar -xzvf {0}/volumes/vfsData.tar.gz -C /var/lib/docker/vfs > /dev/null".format(importDir), shell=True)

def show_help():
	os.system("clear")
	sys.exit("""

	DOCKER MIGRATE

	This tool allows the user to easily migrate images, volumes, and
        containers from one version of Docker to another. With this tool, 
        users can quickly save all their data from the current docker
        instance, change the docker storage backend, and then import all 
        their old data to the new system.

        ## ./docker-migrate export [directory]

        Specify the directory in which to temporarily store the files (can be
        an existing directory, or the command will create one). If no directory
        is specified, `/var/lib/docker-migrate` would be used as default.

        The export command will export all the current images, volumes, and
        containers to the specified directory, in the /images, /volumes,
        /containers subdirectories.

        ## ./docker-migrate import [directory]

        Specify the directory from which to read the files (must be an
        existing directory).If no directory is specified, 
        `/var/lib/docker-migrate` would be used as default.

        The import command will import images, volumes, and containers from
        the specified directory into the new docker instance.

        Primary Author: Jenny Ramseyer, 2015
        Secondary Author: Shishir Mahajan, 2015
""")

main()



	
