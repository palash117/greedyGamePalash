# BASE IMAGE
FROM golang


# Move to working directory /build
WORKDIR /build

# RESOLVE DEPENDENCIES


# COPY ALL LOCAL PACKAGES FILES TO CONTAINER'S WORKSPACE.
ADD . /go/src/DataCollection

# BUILD FILESERVER COMMAND INSIDE CONTIANER
RUN go install DataCollection/


# ENVIRONMENT VARIABLES
ENV FS_PORT=9010
ENV FS_IP=192.168.1.27
ENV FS_GO_BASE_PATH=/go/src/DataCollection/fileServer

# DOCUMENT THAT THE SERVICE LISTENS ON PORT 8080
EXPOSE 9010

# SET DIR TO THAT OF FILESERVER
WORKDIR /go/src/DataCollection

# Command to run when starting the container
CMD ["DataCollection"]

#docker run -p 9010:9010 --name fsl -v /media/pi/sgt/fileserver:/go/src/fileServer/fileServer c5df71a9f255
