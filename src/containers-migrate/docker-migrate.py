#!/usr/bin/env python

# porting docker-migrate.sh to python
import sys
import os
import subprocess

if len(sys.argv) == 3:
    if sys.argv[1] == "export":
        if not os.path.isdir(sys.argv[2]):
            os.mkdir(sys.argv[2])
        if not os.path.isdir(sys.argv[2] + "/images"):
            os.mkdir(sys.argv[2] + "/images")
        if not os.path.isdir(sys.argv[2] + "/containers"):
            os.mkdir(sys.argv[2] + "/containers")
        if not os.path.isdir(sys.argv[2] + "/volumes"):
            os.mkdir(sys.argv[2] + "/volumes")
        # The following is some rather hacky list operations so that we
        # can save images by name:tags instead of by ID
        images = subprocess.check_output("sudo docker images -a", shell=True)
        splitImages = images.split()[7:]  # cut off the headers
        names = []
        tags = []
        ids = []
        for i in range(0, len(splitImages)):
            # only take the image and its tags and the image ID (to help in the <none>:<none> case)
            if (i % 8 == 0):
                names.append(splitImages[i])
                tags.append(splitImages[i+1])
                ids.append(splitImages[i+2])
        for i in range(0, len(names)):
            print("Saving image {0}:{1}".format(names[i], tags[i]))
            if names[i] == '<none>':
                subprocess.call(
                    "sudo docker save {0} > {1}/images/{0}-{0}.tar".format(
                        ids[i], sys.argv[2]), shell=True)
            else:
                subprocess.call(
                    "sudo docker save {0}:{1} > {2}/images/{3}-{4}.tar".format(
                        names[i], tags[i], sys.argv[2], names[i].replace("/", "~"), tags[i].replace("/", "~")), shell=True)
        subprocess.call(
            "sudo tar -zcvf {0}/volumes/volumeData.tar.gz -C /var/lib/docker/volumes . > /dev/null".format(sys.argv[2]), shell=True)
        if os.path.isdir("/var/lib/docker/vfs"):
            subprocess.call("sudo tar -zcvf {0}/volumes/vfsData.tar.gz -C /var/lib/docker/vfs . > /dev/null".format(sys.argv[2]), shell=True)

    elif sys.argv[1] == "import":
        if not os.path.isdir(sys.argv[2]):
            sys.exit("Specified directory {0} does not exist".format(sys.argv[2]))
        tarballs = subprocess.check_output("ls {0}/images".format(sys.argv[2]), shell=True)
        splitTarballs = tarballs.split()
        for i in splitTarballs:
            print("Loading image {0}".format(i))
            subprocess.call("sudo docker load < {0}/images/{1}".format(sys.argv[2], i), shell=True)
        subprocess.call(
            "sudo tar xzvf {0}/volumes/volumeData.tar.gz -C /var/lib/docker/volumes > /dev/null".format(sys.argv[2]), shell=True)
        if os.path.isdir("/var/lib/docker/vfs"):
            subprocess.call(
                "sudo tar -xzvf {0}/volumes/vfsData.tar.gz -C /var/lib/docker/vfs > /dev/null".format(sys.argv[2]), shell=True)
        print("If you created directory {0} solely for the purpose of temporary storage for your files during the docker-migrate import/export process, you may now remove it if you so desire".format(sys.argv[2]))
    else:
        print("Please specify an option: 'import', 'export', or 'help' and try again")

elif len(sys.argv) == 2:
    if (sys.argv[1] == "help"):
        sys.exit("""
# Docker-Migrate

This script allows the user to easily migrate images, volumes, and
containers from one version of Docker to another.  With the script, users can quickly save all their data from the current docker
instance, change the docker storage backend, and then import all their
old data to the new system.

## ./docker-migrate export [directory]

Specify the directory in which to temporarily store the files (can be
an existing directory, or the command will create one)

The export command will export all the current images, volumes, and
containers to the specified directory, in the /images, /volumes,
/containers subdirectories.

## ./docker-migrate import [directory]

Specify the directory from which to read the files (must be an
existing directory)

The import command will import images, volumes, and containers from
the specified directory into the new docker instance.

Jenny Ramseyer, 2015
""")
    else:
        print("Please specify an option: 'import', 'export', or 'help' and try again")
else:
    print("Please specify an option: 'import', 'export', or 'help' and try again")
