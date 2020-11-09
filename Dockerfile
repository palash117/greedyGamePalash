# BASE IMAGE
FROM golang


# Move to working directory /build
WORKDIR /build

# RESOLVE DEPENDENCIES


# COPY ALL LOCAL PACKAGES FILES TO CONTAINER'S WORKSPACE.
ADD . /go/src/greedyGamePalash

RUN go test /go/src/greedyGamePalash/...

# BUILD FILESERVER COMMAND INSIDE CONTIANER
RUN go install greedyGamePalash/


# ENVIRONMENT VARIABLES
ENV DC_PORT=8080


# DOCUMENT THAT THE SERVICE LISTENS ON PORT 8080
EXPOSE 8080

# SET DIR TO THAT OF FILESERVER
WORKDIR /go/src/greedyGamePalash

# Command to run when starting the container
CMD ["greedyGamePalash"]

#docker run -p 9010:9010 --name fsl -v /media/pi/sgt/fileserver:/go/src/fileServer/fileServer c5df71a9f255
