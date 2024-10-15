FROM golang:1.23 as build

WORKDIR /app


COPY . /app


RUN go mod download && go mod verify
RUN go build -o main main.go

COPY . .

FROM public.ecr.aws/lambda/provided:al2023

WORKDIR /app

COPY ./transaction.csv /app/transaction.csv
COPY ./email_template.html /app/email_template.html

COPY --from=build /app/main /app/main
CMD [ "/app/main" ]