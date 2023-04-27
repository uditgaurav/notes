import multiprocessing
import time
import psutil

def consume_cpu(percentage, duration):
    # Get the number of CPU cores
    num_cores = multiprocessing.cpu_count()

    # Calculate the number of processes needed to consume the desired percentage of CPU
    num_processes = int(num_cores * (percentage/100))

    # Start the processes to consume CPU
    processes = []
    for i in range(num_processes):
        p = multiprocessing.Process(target=busy_process)
        processes.append(p)
        p.start()

    # Wait for the specified duration
    time.sleep(duration)

    # Terminate the processes
    for p in processes:
        p.terminate()

def busy_process():
    # Consume CPU by running an infinite loop
    while True:
        pass

if __name__ == '__main__':
    # Consume 50% of CPU for 10 seconds
    consume_cpu(50, 10)
