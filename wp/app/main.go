package main

import (
	"fmt"
	streamer "wp-streamer"
)

func main() {
	// define number of workers and jobs
	const numJobs = 4
	const numWorkers = 2

	//create channels for work and result
	notifyChan := make(chan streamer.ProcessingMessage, numJobs)
	defer close(notifyChan)

	videoQueue := make(chan streamer.VideoProcessingJob)
	defer close(videoQueue)

	// get a worker pool
	wp := streamer.New(videoQueue, numWorkers)
	fmt.Println("wp: ", wp)

	// start worker pool

	// create 4 videos to send to the worker pool

	// send the videos to the worker pool

	// print out results

}
