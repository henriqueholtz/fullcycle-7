/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/henriqueholtz/fullcycle-7/application/grpc"
	"github.com/henriqueholtz/fullcycle-7/application/kafka"
	"github.com/henriqueholtz/fullcycle-7/infrastructure/db"
	"github.com/spf13/cobra"
	"os"
)

var (
	gRPCPortNumber int
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Run gRPC and a Kafka Consumer",
	Run: func(cmd *cobra.Command, args []string) {
		//Start gRPC
		database := db.ConnectDB(os.Getenv("env"))
		go grpc.StartGrpcServer(database, portNumber) //New thread

		//start Kafka-Consumer
		deliveryChan := make(chan ckafka.Event)
		producer := kafka.NewKafkaProducer()
		kafka.Publish("Hello Kafka!", "test", producer, deliveryChan)
		go kafka.DeliveryReport(deliveryChan) //new Thread

		kafkaProcessor := kafka.NewKafkaProcessor(database, producer, deliveryChan)
		kafkaProcessor.Consume()
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
	allCmd.Flags().IntVarP(&gRPCPortNumber, "grpc-port", "p", 500051, "gRPC Port")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
