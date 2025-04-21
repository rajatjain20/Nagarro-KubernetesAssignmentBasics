FROM golang:1.23.2

# Create app directory
RUN mkdir -p /app/

# Set current working dir
WORKDIR /app/

# copy ./src directory into /app
COPY ./src .

# set timezone
ENV TZ=Asia/Kolkata

# Run go commands
RUN go mod init helloworld
RUN go build -v -o ./ .

# Expose the port where application is running
EXPOSE 3000

CMD [ "./helloworld" ]